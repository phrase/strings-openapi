# `.phrase.yml` schema reference

Every key the CLI understands. Tables use plain language and example values.

## Top-level keys (under `phrase:`)

| Key | What it does | Example |
|---|---|---|
| `project_id` | **Required.** Identifies your Phrase project. Find it in Project Settings → API. | `project_id: "abcd1234ef56..."` |
| `host` | Tells the CLI to talk to the US datacenter. Leave it out for EU. | `host: https://api.us.app.phrase.com/v2` |
| `access_token` | Don't use this — set the `PHRASE_ACCESS_TOKEN` environment variable instead so your token never lands in git. | (omit) |
| `file_format` | Default format if you don't repeat it on every source/target. | `file_format: yml` |
| `per_page` | How many items to fetch per page. Defaults to 100, you usually don't need this. | `per_page: 50` |
| `locale_mapping` | Use different locale names on disk than in Phrase. Keys are the names in Phrase, values are the names you want in your files. | <pre>locale_mapping:<br>  English: eng<br>  German: ger</pre> |
| `push` | The upload section. Contains `sources:`. | (see below) |
| `pull` | The download section. Contains `targets:`. | (see below) |

## Push (`phrase.push`)

```yaml
push:
  sources:
    - file: <pattern>
      project_id: <override>      # optional
      branch: <name>              # optional, push to a Phrase branch
      file_format: <override>     # optional
      params:
        file_format: <id>
        locale_id: <id-or-name>
        # ...further upload params (see table below)...
```

### Per-source keys (sibling to `params:`)

| Key | What it does | Example |
|---|---|---|
| `file` | **Required.** Where your locale files live. Supports `*`, `**`, and placeholders. | `file: config/locales/<locale_code>.yml` |
| `project_id` | Use a different project for just this source. | `project_id: "other-project-id"` |
| `branch` | Upload into a Phrase branch instead of the main project. | `branch: my-feature` |
| `file_format` | The format for this source if it differs from the top-level. | `file_format: i18next` |
| `params` | Per-source upload settings (see below). | (see below) |

### `params:` keys for push

| Key | What it does | Example |
|---|---|---|
| `file_format` | The format for this upload. | `file_format: yml` |
| `locale_id` | Force this upload into a specific locale instead of guessing from the file path. | `locale_id: "English"` |
| `tags` | Tag every new key from this upload. Comma-separated for multiple tags. | `tags: "frontend,v2"` |
| `update_translations` | Overwrite existing translations with what's in the file. By default only brand-new keys are created. | `update_translations: true` |
| `update_translations_on_source_match` | For bilingual files, only update if the source text in the file matches the source on Phrase. | `update_translations_on_source_match: true` |
| `update_translation_keys` | Set to `false` to upload translations without creating any new keys. | `update_translation_keys: false` |
| `update_descriptions` | Replace existing key descriptions with whatever's in the file (an empty description wipes the existing one). | `update_descriptions: true` |
| `update_custom_metadata` | Update custom metadata fields from the file. An empty value deletes the field. | `update_custom_metadata: true` |
| `source_locale_id` | The source locale of a bilingual file. | `source_locale_id: "English"` |
| `skip_upload_tags` | Don't tag this upload with the auto-generated upload tag. | `skip_upload_tags: true` |
| `skip_unverification` | Don't mark touched translations as unverified after this upload. | `skip_unverification: true` |
| `file_encoding` | Force a file encoding. Allowed values: `UTF-8`, `UTF-16`, `ISO-8859-1`. | `file_encoding: UTF-16` |
| `autotranslate` | Auto-translate the uploaded locale after the upload finishes. | `autotranslate: true` |
| `verify_mentioned_translations` | Mark every translation in the file as verified. | `verify_mentioned_translations: true` |
| `mark_reviewed` | Mark every translation in the file as reviewed (only available if review workflow is on). | `mark_reviewed: true` |
| `tag_only_affected_keys` | Only apply `tags` to keys whose translations actually changed. | `tag_only_affected_keys: true` |
| `translation_key_prefix` | Add a prefix to every key from this upload. The placeholder `<locale_code>` works here too. | `translation_key_prefix: "web."` |
| `locale_mapping` | For CSV/XLSX uploads: which column holds which locale. | <pre>locale_mapping:<br>  en: 2<br>  de: 3</pre> |
| `format_options` | Format-specific settings — see [format-options.md](./format-options.md). | <pre>format_options:<br>  enable_pluralization: true</pre> |

### Push-only top-level setting (under `phrase.push`, not under `sources:`)

| Key | What it does | Example |
|---|---|---|
| `delete_unmentioned_keys` | Delete any key that doesn't show up in any of the uploaded files. Same as the `--cleanup` / `-c` flag. Use carefully. | `delete_unmentioned_keys: true` |

## Pull (`phrase.pull`)

```yaml
pull:
  targets:
    - file: <pattern>
      project_id: <override>      # optional
      file_format: <override>     # optional
      params:
        file_format: <id>
        locale_id: <id-or-name>
        # ...further download params (see table below)...
```

### Per-target keys (sibling to `params:`)

| Key | What it does | Example |
|---|---|---|
| `file` | **Required.** Where downloaded files should land. Placeholders only — no `*` or `**`. | `file: src/locales/<locale_code>/translation.json` |
| `project_id` | Pull from a different project for this target. | `project_id: "other-project-id"` |
| `file_format` | The format for this target if it differs from the top-level. | `file_format: arb` |
| `params` | Per-target download settings (see below). | (see below) |

### `params:` keys for pull

| Key | What it does | Example |
|---|---|---|
| `file_format` | The format to download. | `file_format: yml` |
| `locale_id` | Download just one specific locale. Use either this *or* a placeholder in `file:`, not both. | `locale_id: "English"` |
| `branch` | Download from a Phrase branch instead of the main project. | `branch: my-feature` |
| `tags` | Only download keys with these tags. Comma-separated for multiple. | `tags: "frontend,v2"` |
| `include_empty_translations` | Include keys that have no translation yet. | `include_empty_translations: true` |
| `exclude_empty_zero_forms` | For plurals, drop the "zero" form when it's empty. | `exclude_empty_zero_forms: true` |
| `include_translated_keys` | Combined with `include_empty_translations: true`, lets you flip to download only the *untranslated* keys. | `include_translated_keys: false` |
| `keep_notranslate_tags` | Keep `[NOTRANSLATE]` markers in the output. | `keep_notranslate_tags: true` |
| `format_options` | Format-specific download settings — see [format-options.md](./format-options.md). | <pre>format_options:<br>  enclose_in_cdata: true</pre> |
| `encoding` | Force a file encoding. Allowed values: `UTF-8`, `UTF-16`, `ISO-8859-1`. | `encoding: UTF-8` |
| `include_unverified_translations` | Set `false` to skip unverified translations. | `include_unverified_translations: false` |
| `use_last_reviewed_version` | Download the last reviewed version of each translation (review workflow only). | `use_last_reviewed_version: true` |
| `fallback_locale_id` | If a translation is missing, fall back to this locale. Requires `include_empty_translations: true`. Don't combine with `use_locale_fallback`. | `fallback_locale_id: "English"` |
| `use_locale_fallback` | Use the fallback chain configured in your Phrase project. Don't combine with `fallback_locale_id`. | `use_locale_fallback: true` |
| `source_locale_id` | When downloading job-scoped files, the source locale to use. | `source_locale_id: "English"` |
| `translation_key_prefix` | Remove this prefix from key names in the downloaded file. | `translation_key_prefix: "web."` |
| `filter_by_prefix` | Only download keys starting with `translation_key_prefix`, and strip the prefix on the way out. | `filter_by_prefix: true` |
| `custom_metadata_filters` | Only download keys whose custom metadata matches. | <pre>custom_metadata_filters:<br>  team: mobile</pre> |
| `locale_ids` | Limit the download to a list of specific locales. | `locale_ids: ["English", "German"]` |
| `updated_since` | Only download keys/translations changed since this date (ISO 8601). | `updated_since: "2026-01-01T00:00:00Z"` |

## Placeholder rules

| Token | Push | Pull | Replaced with |
|---|---|---|---|
| `<locale_code>` | yes | yes | ISO code (e.g. `en`, `de-AT`) |
| `<locale_name>` | yes | yes | Display name (e.g. `English`) |
| `<tag>` | yes | yes (requires `tags:` in `params:`) | Tag name |
| `*` | yes (at most once) | **NO** | Single-segment glob |
| `**` | yes (at most once) | **NO** | Recursive glob |

Each placeholder appears at most once per `file:`. Pull `file:` must contain exactly one locale placeholder OR the target's `params.locale_id` must be set — not both.

## Validation rules the generated file must satisfy

- `pull.targets[].file` MUST NOT contain `*` or `**`. The CLI rejects wildcards in pull targets — only placeholders are allowed. If a push pattern uses a wildcard, generate the equivalent placeholder pattern for pull.
- Each placeholder (`<locale_code>`, `<locale_name>`, `<tag>`) appears at most once per `file:` value.
- Pull targets must have **either** a locale placeholder in `file:` OR `locale_id:` in `params:`, not both.
- If `<tag>` is used in a pull `file:`, `tags:` must be set in `params:`.
- If `locale_mapping` is used, no two remote locales may map to the same on-disk name — the mapping must be reversible.
- Never write `access_token:` into the file. Use `PHRASE_ACCESS_TOKEN` env var.

## JSON Schema

Phrase has a public JSON Schema on SchemaStore: <https://json.schemastore.org/phrase.json>. Most YAML editors (VS Code with the YAML extension, JetBrains, anything that uses `yaml-language-server`) will auto-detect it from the filename `.phrase.yml` and provide autocomplete plus inline validation — no directive needed in the file itself.

The schema covers the common subset (push/pull, basic params, format_options, locale_mapping). It does NOT cover every CLI-supported parameter — fields like `update_translation_keys`, `update_custom_metadata`, `source_locale_id`, `verify_mentioned_translations`, `delete_unmentioned_keys`, and `per_page` are missing. The tables above are the authoritative reference; the schema is a convenience.
