# Troubleshooting `.phrase.yml`

Keyed by error message or symptom. Most failures come from a misconfigured `file:` pattern, a missing `tags:` parameter, or wildcards in the wrong place.

## Common gotchas

Things that catch people on first use.

- **Android `values/` vs `values-en/`.** `values/` (no suffix) is the default locale. The `<locale_code>` placeholder will *not* match it. Push the source with an explicit fixed path + `locale_id:`, and only use the placeholdered pattern for translated locales.
- **Strings Catalog migration.** A project with both `*.lproj/Localizable.strings` and `*.xcstrings` is mid-migration. Pick one; don't push both, or the same keys will collide.
- **Source locale doesn't exist in Phrase yet.** The CLI won't auto-create it on first push. Add the locale in the Phrase web UI before pushing, or push will fail with "locale not found".
- **CSV with a BOM.** Excel-saved CSVs often have a UTF-8 BOM. Set `file_encoding: UTF-8` and verify the first column header parses cleanly — a stray BOM prepended to the key column header won't match `key_name_column: "Key"`.
- **Wildcards in pull.** `*` and `**` only work in push sources. Pull rejects them; use placeholders.
- **Filename guesses on push.** If the source path has a fixed name (e.g. `en.yml`) and no placeholder, set `locale_id:` explicitly. The CLI's filename guess can pick the wrong locale (`en` vs `en-US` vs `English`).
- **`update_translations: true` overwrites.** By default push only creates new keys; with this flag, existing translations get replaced by what's in the file. Standard for source-locale uploads, dangerous for translated ones.
- **`delete_unmentioned_keys: true` is project-wide.** It deletes any Phrase key not in the uploaded files — including keys for locales you didn't push. Only enable when push covers everything.
- **`<locale_code>` ≠ Android folder name.** Phrase emits `zh-CN`, Android expects `zh-rCN`. Use `locale_mapping:` to bridge — see `examples.md` → "Android `r`-prefixed regional locales".
- **Plurals in JSON need `enable_pluralization`.** Without it, `count_one` and `count_other` come in as two separate keys instead of one plural key.

## "Locale not found" / "could not determine locale"

The CLI couldn't map your file (push) or your `locale_id` (pull) to a Phrase locale.

- Check that the locale exists in the Phrase project — Project Settings → Locales.
- `locale_id:` accepts both the locale code (`en`, `de-AT`) and the display name (`English`, `Austrian German`). Pick whichever matches Phrase exactly.
- On push, if the path has no `<locale_code>` / `<locale_name>` placeholder and `locale_id` isn't in `params:`, the CLI has nothing to go on — add `locale_id:`.
- On pull, the placeholder expands to the Phrase locale code by default. Android's `values-zh-rCN/` style needs a `locale_mapping:` entry (see `examples.md` → "Android `r`-prefixed regional locales").

## "File pattern matches no files" / "no source files found" (push)

- The pattern is relative to repo root, not to where the user invoked the CLI from. Don't put `./` at the start.
- `*` matches a single path segment, `**` matches any number. They're not shell globs — `*.json` in a `file:` matches one segment, not a depth-3 path.
- Each placeholder appears at most once per `file:` value.
- Push sources can use `*` or `**` (at most once each). Pull targets cannot.

## "Wildcard not allowed in pull target"

`pull.targets[].file` rejects `*` and `**` outright. Replace them with placeholders:

- `src/locales/*/translation.json` → `src/locales/<locale_code>/translation.json`
- `**/*.arb` → `lib/l10n/app_<locale_code>.arb` (or whatever your real layout is — pull needs an exact path template, not a glob).

## "Placeholder appears multiple times" / "duplicate placeholder"

Each of `<locale_code>`, `<locale_name>`, `<tag>` may appear at most once per `file:`. If you legitimately need both `<locale_code>` and `<tag>` in the same path, that's allowed — the rule is one of *each*, not one total. See `examples.md` → "Splitting files per tag".

## "Both placeholder and locale_id specified"

A pull target may have **either** a locale placeholder in `file:` *or* `locale_id:` in `params:`, not both. Pick one:

- Placeholder → CLI writes one file per locale.
- `locale_id` (with no placeholder) → CLI writes a single file for that locale.

## "tags: required when using `<tag>` placeholder"

If a pull target's `file:` contains `<tag>`, list the tags to expand under `params.tags:`. Push allows `<tag>` without an explicit `tags:` (the CLI infers tags from the matching files); pull does not.

## "Locale mapping is not reversible" / "duplicate local name"

In `locale_mapping:`, no two Phrase locales may map to the same on-disk name — the CLI needs to round-trip in both directions. Fix by choosing distinct names for each entry.

## Push uploaded the wrong locale

The CLI uses (in order): `params.locale_id`, then the placeholder in `file:`, then a guess from filename. If the source path has a fixed name like `en.yml`, set `locale_id:` explicitly — don't rely on filename guessing.

## Pull wrote files into the wrong directory

The path is relative to the directory where you ran `phrase pull`, not the repo root. Run from repo root, or set the working directory in your CI.

## Push deleted keys I didn't expect

`delete_unmentioned_keys: true` (or `--cleanup` / `-c`) removes any Phrase key not present in the uploaded files. If push only uploads the source locale, that's the *source* file's key set — anything missing from it gets deleted project-wide. Default is `false`; only enable it deliberately.

## Pull downloads empty files / very few keys

- `tags:` is set in `params:` and excludes most of the project.
- `locale_ids:` is set and limits to one locale.
- `updated_since:` is set and filtered out everything.
- `include_empty_translations: false` (the default) drops untranslated keys.
- Custom metadata filter (`custom_metadata_filters:`) is too narrow.

## "Unauthorized" / 401 / 403

`PHRASE_ACCESS_TOKEN` is missing, expired, or wrong. Generate a new token at `https://app.phrase.com/settings/oauth_access_tokens` (or `/us/settings/...` on US) and `export PHRASE_ACCESS_TOKEN=<token>` in your shell. Don't put it in `.phrase.yml`.

## "Project not found" / 404

- `project_id:` is wrong. Find the correct one in Project Settings → API.
- You're hitting the wrong datacenter. EU users omit `host:`; US users set `host: https://api.us.app.phrase.com/v2`.

## YAML parse errors / "could not load config"

- The file MUST be wrapped in a top-level `phrase:` key.
- Tabs aren't valid YAML indentation — use spaces.
- VS Code with the YAML extension auto-loads the public schema from the `.phrase.yml` filename and flags issues inline.

## CSV/XLSX upload skipped most rows

`first_content_row:` defaults to `1`; if your file has a header row, set it to `2`. Also confirm `key_index` / `key_name_column` and your `translation_columns` map.

## Android `values/` keeps getting wiped or duplicated

`values/` (no suffix) is the default-locale folder; `values-<lang>/` are translations. Don't pull into `values/<locale_code>/` — the path on disk is `values-<locale_code>/`. The source file lives at `values/strings.xml` and needs an explicit push source with `locale_id:` set; it does *not* match the `values-<locale_code>/` placeholder.

## Plurals turning into separate keys (or vice versa)

JSON-family formats need `format_options.enable_pluralization: true` to treat `thing_one` / `thing_other` as one plural key instead of two. iOS plural keys live in `.stringsdict` by default; set `include_pluralized_keys: true` on the `strings` source to keep them in `.strings`.

## XLIFF round-trips lose state / extradata

Set `include_translation_state: true` to round-trip the `state` attribute, and `export_key_id_as_resname: true` / `export_key_name_hash_as_extradata: true` if you depend on those XLIFF attributes downstream.
