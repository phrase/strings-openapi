---
name: phrase-config
description: Generate a .phrase.yml config file for the phrase-cli and Strings Repo Sync. Detects the project's i18n format and locale file paths and writes a working push/pull config. TRIGGER when the user wants to create, extend, or troubleshoot a .phrase.yml config, or has questions about push/pull behavior, CLI flags, locale file setup, or Repo Sync configuration. DO NOT TRIGGER when the user wants to run push/pull and isn't asking about the config itself.
---

# phrase-config

Generates `.phrase.yml` for any project that uses the Phrase Strings CLI or Strings Repo Sync. The CLI reads this file from the project root to drive `phrase push` and `phrase pull`; Repo Sync uses the same file to sync locale files between a git repository and Phrase.

**Behavior model:** detect → propose → write. Scan the project, infer format and pattern, show one combined confirmation, write the file. Don't ask questions you can defer to a placeholder or a chat-side note.

## References (load on demand)

- [`references/schema.md`](./references/schema.md) — every `.phrase.yml` key, placeholder rules, and the validation rules the generated file must satisfy.
- [`references/formats.md`](./references/formats.md) — detection rules, all 51 format identifiers + default file patterns, and per-format `format_options`.
- [`references/examples.md`](./references/examples.md) — config examples for common project layouts.
- [`references/troubleshooting.md`](./references/troubleshooting.md) — keyed by error message: locale not found, wildcard rejected, plural splitting, values/ vs values-en/, etc.
- [`references/cli.md`](./references/cli.md) — install instructions, `phrase push` / `phrase pull` flags, and `PHRASE_*` environment variables.

Official Phrase docs:
- Config file overview: <https://support.phrase.com/hc/en-us/sections/5784132012828>
- Push/pull configuration: <https://support.phrase.com/hc/en-us/articles/5784118494492>
- JSON Schema: <https://json.schemastore.org/phrase.json>

If `phrase` isn't on PATH, see install options in [`references/cli.md`](./references/cli.md) (Homebrew, asdf, GitHub releases).

## Workflow

If `.phrase.yml` already exists at repo root, skip to the **[Augment path](#augment-path)** below.

### Greenfield path

#### Step 1 — Detect format and path pattern

Use the detection rules in [`references/formats.md`](./references/formats.md). First match wins; ignore `node_modules/`, `vendor/`, `.git/`, `dist/`, `build/`, `target/`, `Pods/`, `.dart_tool/`.

When the format is ambiguous, **always read a sample file** before deciding:

- Generic JSON: open one file. Flat `{"key": "value"}` → `simple_json`. Nested `{"a": {"b": "value"}}` → `nested_json`.
- XLIFF: check the root element's `version=` attribute. `1.2` → `xlf`; `2.0` → `xliff_2`.
- Vue i18n: check whether the locale files are JSON or YAML.

**Detect monorepos upfront.** If you find multiple format roots (e.g. both `ios/` and `android/`, or both `*.arb` and `locales/*.json`), emit multiple `sources` and `targets` — one per platform — instead of silently picking one. Surface the ambiguity to the user.

Then infer the path pattern from the discovered files:

- Filename matches a known locale code (`en`, `de`, `pt-BR`, `zh-CN`) → `<locale_code>`.
- Filename or directory matches a display name (`English`, `German`) → `<locale_name>`.
- Android `values-de/strings.xml` → `res/values-<locale_code>/strings.xml`. `values/` (no suffix) is the default locale — flag it.
- iOS `de.lproj/Localizable.strings` → `<locale_code>.lproj/Localizable.strings`.
- Flutter `lib/l10n/app_en.arb` → `lib/l10n/app_<locale_code>.arb`.

If detection fails, fall back to the format's default pattern from `formats.md`.

**Source locale:** infer from the existing locale files (`en` is the most common). Don't ask — write the inferred value into `locale_id:` and tell the user in chat to change it if their Phrase project uses a different name.

#### Step 2 — One combined confirmation

Show the user a single proposal that includes everything you inferred. Only ask follow-up questions if something is genuinely ambiguous (multiple format roots, no clear source locale, unrecognized layout).

Format:

> Detected: **i18next** at `src/locales/<locale_code>/translation.json` (source locale `en`, EU datacenter). Project ID will be left as a placeholder — replace it with your real ID from Project Settings → API. Write `.phrase.yml`?

If the user is on the US datacenter, they'll say so — only then add `host:`. Don't ask about it preemptively.

#### Step 3 — Generate `.phrase.yml`

Wrap config in the top-level `phrase:` key.

**Default shape:** push uploads only the source locale (one fixed file, no locale placeholder). Pull downloads every locale into placeholdered paths.

```yaml
phrase:
  project_id: "<PROJECT_ID — replace with your Phrase project id>"
  file_format: <detected>

  push:
    sources:
      - file: <source-locale push path, no placeholder>
        params:
          file_format: <detected>
          locale_id: "<inferred source locale code, e.g. en>"
          update_translations: true

  pull:
    targets:
      - file: <detected pull pattern with <locale_code>>
        params:
          file_format: <detected>
```

The push `file:` is a fixed path to the source-locale file (e.g. `config/locales/en.yml`, `lib/l10n/app_en.arb`). `params.locale_id` tells Phrase which locale that file represents.

The pull `file:` keeps the `<locale_code>` placeholder so every locale lands in the right place.

If the user explicitly asks to push *all* locales (they edit translations locally), switch the push source to use the same placeholder pattern as pull and drop `locale_id` from `params`.

US datacenter: add `host: https://api.us.app.phrase.com/v2` directly under `phrase:`.

Multi-platform: emit one source/target per platform, each with its own `file_format` and (optionally) `tags:` to keep keys segmented.

Before writing, verify:
- No wildcards (`*`, `**`) in pull `file:` paths — use placeholders instead.
- Each placeholder (`<locale_code>`, `<locale_name>`, `<tag>`) appears at most once per `file:`.
- Never include `access_token:` in the file.
- The file must be wrapped in a top-level `phrase:` key.

Don't add YAML comments unless they encode a non-obvious constraint.

#### Step 4 — Post-generation

Show the generated file in chat, then end with this copy-pasteable block:

```sh
export PHRASE_ACCESS_TOKEN=<your-token>   # from app.phrase.com → Profile → OAuth Access Tokens
phrase pull
```

Set the token via env var only — never in `.phrase.yml`. Suggest persisting it in a shell rc file or `.envrc` (with direnv).

For ongoing reference, point the user at the `references/` files relevant to their workflow (branches, tags, cleanup, format-specific options, error keys).

### Augment path

When `.phrase.yml` already exists at repo root:

1. Read it. Show the user what's there in 1–2 lines (formats, source/target counts).
2. Ask what to add (e.g. "a new platform target", "a tag-segmented pull", "iOS Strings Catalog source").
3. Run only Step 1 detection for the *new* piece, then propose just the added block — don't rewrite the rest of the file.
4. Use `Edit`, not `Write`, to splice the new entry in.
