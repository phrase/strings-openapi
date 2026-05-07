# Formats: detection and live lookup

No hardcoded format list. Identifiers, default file patterns, and `format_options` drift — always look them up live.

## Step 1 — gather project signals

Walk repo root. Ignore `node_modules/`, `vendor/`, `.git/`, `dist/`, `build/`, `target/`, `Pods/`, `.dart_tool/`. Collect:

- Locale file extensions present (`.arb`, `.xcstrings`, `.strings`, `.stringsdict`, `.xml`, `.po`, `.pot`, `.xlf`/`.xliff`, `.resx`, `.resw`, `.properties`, `.json`, `.yml`/`.yaml`, `.php`, `.csv`, `.xlsx`, `.html`, `.docx`, `.ts`, `.tmx`, `.plist`, etc.).
- Directory layout hints (`*.lproj/`, `res/values*/`, `lib/l10n/`, `config/locales/`, `app/Resources/translations/`, `resources/lang/<locale>/`, `_locales/<locale>/`, `conf/messages.<locale>`).
- Manifest deps:
  - `pubspec.yaml` → Flutter
  - `package.json` → check for `i18next`, `react-intl`, `vue-i18n`, `next-intl`, `next-translate`, `angular-translate`
  - `Gemfile` mentioning `rails` → Rails
  - Symfony / Laravel layout markers
  - Go with `go-i18n`
- For ambiguous JSON/XLIFF/Vue: read a sample file.
  - Flat `{"key":"value"}` vs nested `{"a":{"b":"value"}}`.
  - XLIFF root `version="1.2"` vs `version="2.0"`.
  - Vue i18n with `.json` locale files vs `.yml`.

If multiple locale roots (e.g. both `ios/` and `android/`, or both `*.arb` and `locales/*.json`), surface that — emit one source/target per platform or ask which to wire.

## Step 2 — resolve format identifier live

Run:

```sh
phrase formats list
```

Output is JSON. Each entry:

- `api_name` — value for `file_format:` / `params.file_format:`.
- `name` — human-readable display name (matches the support docs).
- `extension` — file extension.
- `default_file` — default pull pattern. Use as fallback when project layout doesn't dictate one.
- `importable` / `exportable` — whether push/pull supported.

Match the project signals from Step 1 against `name` + `extension` + `default_file` to pick `api_name`. Confirm the choice with the user before writing.

If signals don't disambiguate (e.g. generic `.json` could be `simple_json`, `nested_json`, `react_simple_json`, `i18next`, `go_i18n`, `json`/Chrome, …), narrow by manifest deps + file shape, then ask the user to confirm.

## Step 3 — per-format options via support docs

`format_options` are not in the CLI output. Look them up only when the user wants non-default behavior or asks about an option.

Help center: <https://support.phrase.com/hc/en-us/articles/9652464547740-List-of-Supported-File-Types-Strings>

1. Fetch the help center page. Find the link for the format's display `name` (e.g. "Android Strings", "Apple Strings Catalog").
2. Fetch that detail page. Read the **Format Options** section — each option lists the YAML key, accepted values, behavior.
3. Only set options the page declares for that format. If a requested option isn't there, push back — likely belongs to a different format or has been renamed.
