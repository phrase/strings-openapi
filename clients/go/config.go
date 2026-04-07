package phrase

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Credentials contains all information to authenticate against phrase.com or a custom host.
type Credentials struct {
	Username string
	Token    string
	TFA      bool
	TFAToken string
	Host     string
}

// Config contains all information from a .phraseapp.yml config file
type Config struct {
	Credentials
	Debug bool

	Page    *int
	PerPage *int

	DefaultProjectID  string
	DefaultFileFormat string

	Defaults map[string]map[string]interface{}

	Pull []byte
	Push []byte

	LocaleMapping map[string]string

	UserAgent string
}

var configNames = []string{".phrase.yml", ".phraseapp.yml"}

// ReadConfig reads a .phrase.yml config file
func ReadConfig(cfgFilePath string) (*Config, error) {
	rawCfg := map[string]*Config{}
	content, err := configContent(cfgFilePath)
	switch {
	case err != nil:
		return nil, err
	case content == nil:
		return &Config{}, nil
	default:
		err := yaml.Unmarshal(content, rawCfg)
		if err != nil {
			return nil, err
		}

		if cfg, found := rawCfg["phrase"]; found {
			return cfg, nil
		}
		if cfg, found := rawCfg["phraseapp"]; found {
			return cfg, nil
		}

		return nil, errors.New("'phrase' key is missing in config")
	}
}

func configContent(cfgFilePath string) ([]byte, error) {
	var path string
	var err error

	if cfgFilePath != "" {
		path = cfgFilePath
		err = nil
	} else {
		path, err = configPath()
	}

	switch {
	case err != nil:
		return nil, err
	case path == "":
		return nil, nil
	default:
		return os.ReadFile(path)
	}
}

func configPath() (string, error) {
	if possiblePath := os.Getenv("PHRASEAPP_CONFIG"); possiblePath != "" {
		_, err := os.Stat(possiblePath)
		if err == nil {
			return possiblePath, nil
		}

		if os.IsNotExist(err) {
			err = fmt.Errorf("file %q (from PHRASEAPP_CONFIG environment variable) doesn't exist", possiblePath)
		}

		return "", err
	}

	workingDir, err := os.Getwd()
	if err != nil {
		return "", nil
	}

	for _, configName := range configNames {
		dir := workingDir
		// loop up the directory tree
		for {
			possiblePath := filepath.Join(dir, configName)
			if _, err := os.Stat(possiblePath); err == nil {
				return possiblePath, nil
			}
			if filepath.Dir(dir) == dir { // reached the root directory
				break
			}
			dir = filepath.Dir(dir)
		}
	}

	for _, configName := range configNames {
		possiblePath := filepath.Join(defaultConfigDir(), configName)
		if _, err := os.Stat(possiblePath); err == nil {
			return possiblePath, nil
		}
	}

	return "", nil
}

func (cfg *Config) UnmarshalYAML(unmarshal func(i interface{}) error) error {
	m := map[string]interface{}{}
	localeMapping := map[string]interface{}{}
	err := ParseYAMLToMap(unmarshal, map[string]interface{}{
		"access_token":   &cfg.Credentials.Token,
		"host":           &cfg.Credentials.Host,
		"debug":          &cfg.Debug,
		"page":           &cfg.Page,
		"per_page":       &cfg.PerPage,
		"project_id":     &cfg.DefaultProjectID,
		"file_format":    &cfg.DefaultFileFormat,
		"push":           &cfg.Push,
		"pull":           &cfg.Pull,
		"locale_mapping": &localeMapping,
		"defaults":       &m,
	})
	if err != nil {
		return err
	}

	cfg.Defaults = map[string]map[string]interface{}{}
	for path, rawConfig := range m {
		cfg.Defaults[path], err = ValidateIsRawMap("defaults."+path, rawConfig)
		if err != nil {
			return err
		}
	}

	if len(localeMapping) > 0 {
		cfg.LocaleMapping, err = ConvertToStringMap(localeMapping)
		if err != nil {
			return err
		}
	}

	return nil
}

const cfgValueErrStr = "configuration key %q has invalid value: %T\nsee https://support.phrase.com/hc/en-us/sections/5784132012828"
const cfgKeyErrStr = "configuration key %q has invalid type: %T\nsee https://support.phrase.com/hc/en-us/sections/5784132012828"
const cfgInvalidKeyErrStr = "configuration key %q unknown\nsee https://support.phrase.com/hc/en-us/sections/5784132012828"

func ValidateIsString(k string, v interface{}) (string, error) {
	s, ok := v.(string)
	if !ok {
		return "", fmt.Errorf(cfgValueErrStr, k, v)
	}
	return s, nil
}

func ValidateIsBool(k string, v interface{}) (bool, error) {
	b, ok := v.(bool)
	if !ok {
		return false, fmt.Errorf(cfgValueErrStr, k, v)
	}
	return b, nil
}

func ValidateIsInt(k string, v interface{}) (int, error) {
	i, ok := v.(int)
	if !ok {
		return 0, fmt.Errorf(cfgValueErrStr, k, v)
	}
	return i, nil
}

func ValidateIsRawMap(k string, v interface{}) (map[string]interface{}, error) {
	raw, ok := v.(map[interface{}]interface{})
	if !ok {
		return nil, fmt.Errorf(cfgValueErrStr, k, v)
	}

	ps := map[string]interface{}{}
	for mk, mv := range raw {
		s, ok := mk.(string)
		if !ok {
			return nil, fmt.Errorf(cfgKeyErrStr, fmt.Sprintf("%s.%v", k, mk), mk)
		}
		ps[s] = mv
	}
	return ps, nil
}

func ConvertToStringMap(raw map[string]interface{}) (map[string]string, error) {
	ps := map[string]string{}
	for mk, mv := range raw {
		switch v := mv.(type) {
		case string:
			ps[mk] = v
		case bool:
			ps[mk] = fmt.Sprintf("%t", v)
		case int:
			ps[mk] = fmt.Sprintf("%d", v)
		default:
			return nil, fmt.Errorf("invalid type of key %q: %T", mk, mv)
		}
	}
	return ps, nil
}

// Calls the YAML parser function (see yaml.v2/Unmarshaler interface) with a map
// of string to interface. This map is then iterated to match against the given
// map of keys to fields, validates the type and sets the fields accordingly.
func ParseYAMLToMap(unmarshal func(interface{}) error, keysToField map[string]interface{}) error {
	m := map[string]interface{}{}
	if err := unmarshal(m); err != nil {
		return err
	}

	var err error
	for k, v := range m {
		value, found := keysToField[k]
		if !found {
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}

		switch val := value.(type) {
		case *string:
			*val, err = ValidateIsString(k, v)
		case *int:
			*val, err = ValidateIsInt(k, v)
		case **int:
			*val = new(int)
			**val, err = ValidateIsInt(k, v)
		case *bool:
			*val, err = ValidateIsBool(k, v)
		case *map[string]interface{}:
			*val, err = ValidateIsRawMap(k, v)
		case *[]byte:
			*val, err = yaml.Marshal(v)
		default:
			err = fmt.Errorf(cfgValueErrStr, k, value)
		}
		if err != nil {
			return err
		}
	}

	return nil
}
