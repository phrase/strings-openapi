# Formats: detection, identifiers, and per-format options

Three things in one file: how to recognize a project's format, the full list of identifiers + default file patterns, and per-format `format_options`.

## Detection rules

Walk the project tree from repo root. Ignore `node_modules/`, `vendor/`, `.git/`, `dist/`, `build/`, `target/`, `Pods/`, `.dart_tool/`. Apply rules in order; first match wins. For monorepos, multiple matches are allowed.

| Signal | `file_format` | Typical project |
|---|---|---|
| `pubspec.yaml` + `*.arb` files (often under `lib/l10n/`) | `arb` | Flutter |
| `*.xcstrings` files | `strings_catalog` | Apple String Catalog (Xcode 15+) |
| `*.lproj/Localizable.strings` | `strings` | iOS/macOS (legacy) |
| `*.lproj/Localizable.stringsdict` (alongside `.strings`) | `stringsdict` | iOS plurals (legacy) |
| `*.plist` files in localized form | `plist` | Objective-C/Cocoa property lists |
| `res/values*/strings.xml` | `xml` | Android |
| `*.po` files | `gettext` | Python/PHP/Django/C |
| `*.pot` template files | `gettext_template` | gettext templates |
| `config/locales/*.yml` + `Gemfile` mentioning `rails` | `yml` | Rails |
| `app/Resources/translations/*.yml` (Symfony layout) | `yml_symfony2` | Symfony |
| `app/Resources/translations/messages.*.xlf` (Symfony) | `symfony_xliff` | Symfony XLIFF |
| `resources/lang/<locale>/messages.php` | `laravel` | Laravel |
| `*.php` array translation files (non-Laravel) | `php_array` | Generic PHP |
| `package.json` deps include `i18next` | `i18next` (or `i18next_4` if v4+) | i18next (React/Vue/Node) |
| `package.json` deps include `react-intl` — flat files | `react_simple_json` | React Intl flat |
| `package.json` deps include `react-intl` — nested files | `react_nested_json` | React Intl nested |
| `package.json` deps include `vue-i18n` (inspect file) | `simple_json` or `yml` | Vue i18n |
| `package.json` deps include `next-intl` or `next-translate` | `nested_json` | Next.js |
| `package.json` deps include `angular-translate` | `angular_translate` | AngularJS |
| `app/locales/*/translations.js` | `ember_js` | Ember.js |
| `*.go` with `go-i18n` and `*.all.json` | `go_i18n` | Go i18n |
| Chrome extension `_locales/<locale>/messages.json` | `json` | Chrome i18n |
| `conf/messages.<locale>` | `play_properties` | Play Framework |
| `*.properties` under `src/main/resources/` | `properties` | Java |
| Mozilla addon `*.properties` | `mozilla_properties` | Firefox/XUL |
| `*.resx` files | `resx` | .NET / WinForms |
| `*.resw` files | `windows8_resource` | Windows Store apps |
| `Resources/AppResources.<locale>.resx` | `resx_windowsphone` | Windows Phone |
| `*.xlf` / `*.xliff` (XLIFF 1.2) | `xlf` | Angular i18n, .NET MAUI, generic |
| `*.xlf` / `*.xliff` declared as XLIFF 2.0 (`version="2.0"`) | `xliff_2` | Modern XLIFF |
| `*.ts` (Qt translation source) | `ts` | Qt Linguist |
| `*.tmx` | `tmx` | Translation memory |
| `*.xlsx` | `xlsx` | Excel workflow |
| `*.csv` | `csv` | Spreadsheet workflow |
| `*.html` / `*.htm` | `html` or `html_5` | Marketing/CMS pages |
| `*.docx` | `docx` | Documentation |
| `locales/*.json` (flat key/value) | `simple_json` | Generic |
| `locales/*.json` (nested objects) | `nested_json` | Generic |

**Always read a sample file when the format is ambiguous:**

- Generic JSON: open one locale file. Flat `{"key": "value"}` → `simple_json`. Nested `{"a": {"b": "value"}}` → `nested_json`.
- XLIFF: check the root element's `version=` attribute. `1.2` → `xlf`; `2.0` → `xliff_2`.
- Vue i18n: deps + a `.json` locale file → `simple_json` (or `nested_json`); deps + a `.yml` locale file → `yml`.

If detection finds **multiple** locale-file roots (e.g. both `*.arb` and `locales/*.json`, or both `ios/` and `android/`), surface that to the user and ask which to wire up — or whether to emit one source/target per platform.

If nothing matches, ask the user to pick from the identifiers below.

## Identifiers and default file patterns

The CLI accepts any identifier returned by `phrase formats list`. The 51 public formats:

| Identifier | Display name | Extension | Default file pattern |
|---|---|---|---|
| `arb` | ARB | `.arb` | `./<locale_name>.arb` |
| `yml` | Ruby/Rails YAML | `.yml` `.yaml` | `./config/locales/<locale_name>.yml` |
| `gettext` | Gettext | `.po` | `./<locale_name>.po` |
| `gettext_template` | Gettext template | `.pot` | `./<locale_name>.pot` |
| `gettext_mo` | Gettext compiled | `.mo` | `./<locale_name>.mo` |
| `xml` | Android Strings | `.xml` | `./values-<locale_code>/strings.xml` |
| `strings` | iOS Localizable Strings | `.strings` | `./<locale_code>.lproj/Localizable.strings` |
| `stringsdict` | iOS Localizable Stringsdict | `.stringsdict` | `./<locale_code>.lproj/Localizable.stringsdict` |
| `strings_catalog` | Apple Strings Catalog | `.xcstrings` | `./<locale_name>.xcstrings` |
| `xlf` | XLIFF 1.2 | `.xlf` `.xliff` | `./<locale_name>.xlf` |
| `xliff_2` | XLIFF 2.0 | `.xlf` `.xliff` | `./<locale_name>.xlf` |
| `symfony_xliff` | Symfony XLIFF | `.xlf` `.xliff` | `./messages.<locale_name>.xlf` |
| `qph` | Qt Phrase Book | `.qph` | `./<locale_name>.qph` |
| `ts` | Qt Translation Source | `.ts` | `./<locale_name>.ts` |
| `json` | Chrome JSON i18n | `.json` | `./<locale_name>.json` |
| `simple_json` | Simple JSON (flat) | `.json` | `./<locale_name>.json` |
| `nested_json` | Nested JSON | `.json` | `./<locale_name>.json` |
| `react_simple_json` | React-Intl Simple JSON | `.json` | `./<locale_name>.json` |
| `react_nested_json` | React-Intl Nested JSON | `.json` | `./<locale_name>.json` |
| `i18next` | i18next | `.json` | `./locales/<locale_name>/translations.json` |
| `i18next_4` | i18next 4 | `.json` | `./locales/<locale_name>/translations.json` |
| `go_i18n` | Go i18n JSON | `.json` | `./<locale_name>.all.json` |
| `node_json` | i18n-node-2 JSON | `.js` | `./locales/<locale_name>.js` |
| `angular_translate` | Angular Translate | `.json` | `./i18n/<locale_code>.json` |
| `ember_js` | Ember.js | `.js` | `./app/locales/<locale_code>/translations.js` |
| `genesys_json` | Genesys JSON | `.json` | `./<locale_name>.json` |
| `strings_json` | Strings JSON | `.json` | `./<locale_name>.json` |
| `resx` | .NET ResX | `.resx` | `./<locale_name>.resx` |
| `resx_windowsphone` | Windows Phone ResX | `.resx` | `./Resources/AppResources.<locale_code>.resx` |
| `windows8_resource` | Windows 8 Resource | `.resw` | `./Resources/AppResources.<locale_code>.resw` |
| `properties` | Java Properties | `.properties` | `./MessagesBundle_<locale_name>.properties` |
| `mozilla_properties` | Mozilla Properties | `.properties` | `./<locale_code>.properties` |
| `properties_xml` | Java Properties XML | `.xml` | `./MessagesBundle_<locale_code>.xml` |
| `play_properties` | Play Framework Properties | `.locale` | `./conf/messages.<locale_code>` |
| `ini` | INI | `.ini` | `./<locale_name>.ini` |
| `plist` | Objective-C/Cocoa Property List | `.plist` | `./<locale_code>.plist` |
| `tmx` | TMX Translation Memory | `.tmx` | `./<locale_code>.tmx` |
| `xlsx` | Excel XLSX | `.xlsx` | `./<locale_code>.xlsx` |
| `csv` | CSV | `.csv` | `./<locale_code>.csv` |
| `txt` | Tab-separated TXT | `.txt` | `./<locale_code>.txt` |
| `zendesk_csv` | Zendesk CSV | `.csv` | `./<locale_code>.csv` |
| `php_array` | PHP Array | `.php` | `./locale_<locale_code>.php` |
| `laravel` | Laravel/F3/Kohana Array | `.php` | `./resources/lang/<locale_code>/messages.php` |
| `yml_symfony` | Symfony YAML | `.yml` `.yaml` | `./app/Resources/translations/<locale_code>.yml` |
| `yml_symfony2` | Symfony2 YAML | `.yml` `.yaml` | `./app/Resources/translations/<locale_code>.yml` |
| `episerver` | Episerver XML | `.xml` | `./<locale_name>.xml` |
| `linguist_xml` | Linguist XML | `.xml` | `./linguist.xml` |
| `linguist_xml_2` | Linguist XML 2 | `.xml` | `./linguist.xml` |
| `html` | HTML | `.html` `.htm` | `./<locale_name>.html` |
| `html_5` | HTML 5 | `.html` `.htm` | `./<locale_name>.html` |
| `docx` | Word DOCX | `.docx` | `./<locale_name>.docx` |

The default file pattern is what `phrase init` would suggest for that format — a good fallback when you can't infer the pattern from existing files.

Full format guide with import/export quirks: <https://support.phrase.com/hc/en-us/articles/9652464547740>.

---

## Per-format options (`params.format_options`)

Per-format options go under `params.format_options:` in `.phrase.yml`. They are not interchangeable across formats — only set options that the chosen format declares.

### `xml` (Android Strings)

| Option | What it does | Example |
|---|---|---|
| `convert_placeholder` | Convert `%s`-style placeholders to and from Android's `%1$s` style. | `convert_placeholder: true` |
| `escape_linebreaks` | Escape line breaks as `\n` when downloading. | `escape_linebreaks: true` |
| `unescape_linebreaks` | Turn `\n` back into real newlines on upload. | `unescape_linebreaks: true` |
| `enclose_in_cdata` | Wrap each value in `<![CDATA[…]]>`. | `enclose_in_cdata: true` |
| `preserve_cdata` | Keep existing CDATA wrappers when re-uploading. | `preserve_cdata: true` |
| `escape_tags` / `unescape_tags` | Escape or unescape HTML/XML tags inside values. | `escape_tags: true` |
| `escape_android_chars` / `unescape_android_chars` | Escape or unescape Android-special characters (`'`, `"`, `@`, `?`, `\`). | `escape_android_chars: true` |
| `indent_size` | How many characters to indent. | `indent_size: 4` |
| `indent_style` | Indent with spaces or tabs. | `indent_style: "space"` |
| `include_tools_ignore` | Add `tools:ignore` attributes to keys that have them in Phrase. | `include_tools_ignore: true` |
| `include_tools_locale_definition` | Add `tools:locale` to the `<resources>` element. | `include_tools_locale_definition: true` |

### `strings` (iOS)

| Option | What it does | Example |
|---|---|---|
| `convert_placeholder` | Normalize placeholder syntax (`%s` ↔ `%@`). | `convert_placeholder: true` |
| `include_pluralized_keys` | Keep keys with plural forms in `.strings` (instead of moving them to `.stringsdict`). | `include_pluralized_keys: true` |
| `multiline_descriptions` | Wrap long key descriptions across multiple comment lines. | `multiline_descriptions: true` |

### `stringsdict`

| Option | What it does | Example |
|---|---|---|
| `convert_placeholder` | Normalize placeholder syntax. | `convert_placeholder: true` |

### `strings_catalog` (Apple `.xcstrings`)

| Option | What it does | Example |
|---|---|---|
| `convert_placeholder` | Normalize placeholder syntax. | `convert_placeholder: true` |
| `ignore_translation_state` | Don't read or write the translation state field. | `ignore_translation_state: true` |
| `default_extraction_state` | The `extractionState` value for new entries. | `default_extraction_state: "manual"` |

### `xlf` (XLIFF 1.2)

| Option | What it does | Example |
|---|---|---|
| `enclose_in_cdata` | Wrap target text in CDATA. | `enclose_in_cdata: true` |
| `content_as_literal` | Treat content as literal text — don't decode XML entities. | `content_as_literal: true` |
| `include_translation_state` | Add the `state` attribute to translation units. | `include_translation_state: true` |
| `replace_target_translations_with_empty_string` | When a translation is missing, write an empty target instead of falling back to the source. | `replace_target_translations_with_empty_string: true` |
| `keep_plural_skeletons` | Preserve the plural-form skeleton on round-trip. | `keep_plural_skeletons: true` |
| `ignore_source_translations` | Skip the source-language text on import. | `ignore_source_translations: true` |
| `ignore_target_translations` | Skip the target-language text on import. | `ignore_target_translations: true` |
| `export_key_id_as_resname` | Put the Phrase key id in the `resname` attribute. | `export_key_id_as_resname: true` |
| `export_key_name_hash_as_extradata` | Put a hash of the key name in the `extradata` attribute. | `export_key_name_hash_as_extradata: true` |
| `delimit_placeholders` | Wrap placeholders with delimiters during export. | `delimit_placeholders: true` |
| `strip_placeholder_delimiters` | Remove placeholder delimiters during import. | `strip_placeholder_delimiters: true` |
| `override_file_language` | Force the file's language attribute to match the downloaded locale. | `override_file_language: true` |
| `key_name_attribute` | Which XML attribute holds the key name. Defaults to `id`. | `key_name_attribute: "resname"` |
| `custom_metadata_columns` | Map your custom metadata fields to XLIFF attributes. | <pre>custom_metadata_columns:<br>  team: extradata</pre> |

### `xliff_2` (XLIFF 2.0)

Most of the XLIFF 1.2 options work here too: `enclose_in_cdata`, `include_translation_state`, `replace_target_translations_with_empty_string`, `content_as_literal`, `keep_plural_skeletons`, `ignore_source_translations`, `ignore_target_translations`, `override_file_language`.

### `symfony_xliff`

| Option | What it does | Example |
|---|---|---|
| `enclose_in_cdata` | Wrap target text in CDATA. | `enclose_in_cdata: true` |

### `simple_json` / `nested_json` / `react_nested_json`

| Option | What it does | Example |
|---|---|---|
| `enable_pluralization` | Treat keys with plural suffixes (`_one`, `_other`, etc.) as plural variants of one key. | `enable_pluralization: true` |

### `i18next` / `i18next_4`

| Option | What it does | Example |
|---|---|---|
| `nesting` | Output nested objects (the default). Turn off for flat keys. | `nesting: false` |

### `gettext` / `gettext_template`

| Option | What it does | Example |
|---|---|---|
| `msgid_as_default` | When `msgstr` is empty, use `msgid` as the fallback. | `msgid_as_default: true` |
| `is_bilingual_file` | (`gettext` only) Treat the file as bilingual — both source and target. | `is_bilingual_file: true` |

### `properties` (Java)

| Option | What it does | Example |
|---|---|---|
| `escape_single_quotes` | Double single-quotes (`'` → `''`) for Java `MessageFormat`. | `escape_single_quotes: true` |
| `omit_separator_space` | Don't put spaces around `=`. | `omit_separator_space: true` |
| `crlf_line_terminators` | Use Windows line endings (`\r\n`). | `crlf_line_terminators: true` |
| `escape_meta_chars` | Escape backslashes, tabs, and other control characters. | `escape_meta_chars: true` |

### `play_properties`

| Option | What it does | Example |
|---|---|---|
| `escape_single_quotes` | Double single-quotes for Play's MessageFormat. | `escape_single_quotes: true` |

### `csv` / `txt`

| Option | What it does | Example |
|---|---|---|
| `column_separator` | Field separator. Defaults to `,` for CSV, tab for TXT. | `column_separator: ";"` |
| `quote_char` | Quote character. Defaults to `"`. | `quote_char: "'"` |
| `first_content_row` | Row number where actual translation rows start (1-based). | `first_content_row: 2` |
| `key_index` / `key_name_column` | Which column holds the key. Use either an index or a header name. | `key_name_column: "Key"` |
| `key_id_column` | Which column holds the Phrase key id. | `key_id_column: 1` |
| `translation_index` / `translation_indexes` / `translation_columns` | Which column(s) hold translations. Map locale → column. | <pre>translation_columns:<br>  en: 2<br>  de: 3</pre> |
| `comment_index` / `comment_column` | Which column holds key descriptions. | `comment_column: "Notes"` |
| `tag_column` | Which column holds tags. | `tag_column: 5` |
| `max_characters_allowed_column` | Which column holds the per-key character limit. | `max_characters_allowed_column: "Limit"` |
| `custom_metadata_columns` | Map custom metadata fields to columns. | <pre>custom_metadata_columns:<br>  team: 6</pre> |
| `enable_pluralization` | Treat plural-suffixed keys as plurals. | `enable_pluralization: true` |
| `include_headers` | Write a header row when downloading. | `include_headers: true` |
| `export_tags` / `export_system_tags` | Include tag columns when downloading. | `export_tags: true` |
| `export_key_id` | Include the Phrase key id column when downloading. | `export_key_id: true` |
| `export_max_characters_allowed` | Include the character-limit column when downloading. | `export_max_characters_allowed: true` |
| `group_by_key_name` | Group rows by key name in the output. | `group_by_key_name: true` |

### `xlsx` (Excel)

These CSV options also work for Excel files: `first_content_row`, `key_name_column`, `translation_column`, `translation_columns`, `custom_metadata_columns`, `comment_column`, `tag_column`, `max_characters_allowed_column`, `enable_pluralization`, `export_tags`, `export_system_tags`, `export_max_characters_allowed`.

### `strings_json`

| Option | What it does | Example |
|---|---|---|
| `custom_metadata_columns` | Map custom metadata fields to JSON properties. | <pre>custom_metadata_columns:<br>  team: team_field</pre> |
| `export_description` | Include key descriptions in the output. | `export_description: true` |
| `export_tags` / `export_system_tags` | Include tag fields in the output. | `export_tags: true` |
| `export_max_characters_allowed` | Include the per-key character limit. | `export_max_characters_allowed: true` |
| `include_translation_state` | Include the translation state field. | `include_translation_state: true` |
| `ignore_translation_state` | Skip importing translation state. | `ignore_translation_state: true` |
| `export_use_ordinal_rules` | Use ordinal plural rules (1st, 2nd, 3rd) instead of cardinal. | `export_use_ordinal_rules: true` |

Formats not listed above (`arb`, `yml`, `json`, `react_simple_json`, `qph`, `ts`, `go_i18n`, `node_json`, `resx`, `resx_windowsphone`, `windows8_resource`, `mozilla_properties`, `properties_xml`, `plist`, `ini`, `tmx`, `zendesk_csv`, `php_array`, `laravel`, `yml_symfony`, `yml_symfony2`, `episerver`, `linguist_xml`, `linguist_xml_2`, `angular_translate`, `ember_js`, `genesys_json`, `html`, `html_5`, `docx`, `gettext_mo`) currently expose **no per-format options**.
