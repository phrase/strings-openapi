# Config examples

Each example shows what the skill should produce for a typical project.

For first-time gotchas and error symptoms (locale not found, wildcard rejected, plurals splitting), see [`troubleshooting.md`](./troubleshooting.md).

**Default shape:** push uploads only the source locale (one fixed path, no placeholder). Pull downloads every locale using a placeholder.

## Rails

Detected: `config/locales/en.yml`, `config/locales/de.yml` + `Gemfile` mentions `rails`. Source locale: `en` ("English" in Phrase).

```yaml
phrase:
  project_id: "abc123..."
  file_format: yml
  push:
    sources:
      - file: config/locales/en.yml
        params:
          file_format: yml
          locale_id: "English"
          update_translations: true
  pull:
    targets:
      - file: config/locales/<locale_code>.yml
        params:
          file_format: yml
```

## React + i18next

Detected: `package.json` deps include `react-i18next`; `src/locales/en/translation.json`. Source locale: `en`.

```yaml
phrase:
  project_id: "abc123..."
  file_format: i18next
  push:
    sources:
      - file: src/locales/en/translation.json
        params:
          file_format: i18next
          locale_id: "English"
          update_translations: true
  pull:
    targets:
      - file: src/locales/<locale_code>/translation.json
        params:
          file_format: i18next
```

## Flutter

Detected: `pubspec.yaml`; `lib/l10n/app_en.arb`, `lib/l10n/app_de.arb`. Source locale: `en`.

```yaml
phrase:
  project_id: "abc123..."
  file_format: arb
  push:
    sources:
      - file: lib/l10n/app_en.arb
        params:
          file_format: arb
          locale_id: "English"
          update_translations: true
  pull:
    targets:
      - file: lib/l10n/app_<locale_code>.arb
        params:
          file_format: arb
```

## iOS + Android monorepo (multi-platform)

Source locale: `en` for both platforms.

```yaml
phrase:
  project_id: "abc123..."
  push:
    sources:
      - file: ios/en.lproj/Localizable.strings
        params:
          file_format: strings
          locale_id: "English"
          tags: ios
          update_translations: true
      - file: android/app/src/main/res/values/strings.xml
        params:
          file_format: xml
          locale_id: "English"
          tags: android
          update_translations: true
  pull:
    targets:
      - file: ios/<locale_code>.lproj/Localizable.strings
        params:
          file_format: strings
          tags: ios
      - file: android/app/src/main/res/values-<locale_code>/strings.xml
        params:
          file_format: xml
          tags: android
```

Note Android's quirk: `values/` (no suffix) holds the source locale, while translated locales go into `values-<locale_code>/`. The push source matches `values/` exactly; the pull target uses the placeholdered path.

## Android `r`-prefixed regional locales (locale_mapping)

Android writes regional locales as `values-<lang>-r<REGION>/` (e.g. `values-zh-rCN/`, `values-pt-rBR/`), but Phrase stores those locales as `zh-CN`, `pt-BR`. The `<locale_code>` placeholder expands to the Phrase code, so use `locale_mapping` to translate Phrase codes to the Android-flavored folder name on disk.

```yaml
phrase:
  project_id: "abc123..."
  file_format: xml
  locale_mapping:
    zh-CN: zh-rCN
    zh-TW: zh-rTW
    pt-BR: pt-rBR
    en-GB: en-rGB
  push:
    sources:
      - file: app/src/main/res/values/strings.xml
        params:
          file_format: xml
          locale_id: "English"
          update_translations: true
  pull:
    targets:
      - file: app/src/main/res/values-<locale_code>/strings.xml
        params:
          file_format: xml
```

`<locale_code>` will now expand to `zh-rCN` for the `zh-CN` Phrase locale, so the file lands in `values-zh-rCN/strings.xml`. Pure-language locales like `de` or `fr` need no mapping — they go to `values-de/`, `values-fr/` as-is.

## Splitting files per tag

When keys are tagged in Phrase (e.g. `checkout`, `dashboard`, `marketing`), use the `<tag>` placeholder to push and pull each tag into its own file. Both `<locale_code>` and `<tag>` appear in the path; `tags:` in `params:` lists which tags to expand.

```yaml
phrase:
  project_id: "abc123..."
  file_format: i18next
  push:
    sources:
      - file: src/locales/<locale_code>/<tag>.json
        params:
          file_format: i18next
          tags: checkout,dashboard,marketing
          update_translations: true
  pull:
    targets:
      - file: src/locales/<locale_code>/<tag>.json
        params:
          file_format: i18next
          tags: checkout,dashboard,marketing
```

On pull, one file is written per (locale, tag) pair. On push, the CLI uploads each matching file and tags its keys accordingly — keys in `checkout.json` get the `checkout` tag, and so on.

## US datacenter

Same shape as any of the above, with `host:` directly under `phrase:`:

```yaml
phrase:
  host: https://api.us.app.phrase.com/v2
  project_id: "abc123..."
  ...
```

## Pushing all locales (override)

If you edit translations locally and want to upload every locale, mirror the pull pattern in the push source and drop `locale_id`:

```yaml
phrase:
  project_id: "abc123..."
  file_format: yml
  push:
    sources:
      - file: config/locales/<locale_code>.yml
        params:
          file_format: yml
          update_translations: true
  pull:
    targets:
      - file: config/locales/<locale_code>.yml
        params:
          file_format: yml
```
