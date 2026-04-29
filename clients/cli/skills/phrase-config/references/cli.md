# CLI flags & environment variables

These don't go in `.phrase.yml` — they're passed on the command line or set in your shell. CLI flags override config file values.

## Installing the CLI

If `phrase` isn't on PATH, point the user at one of these:

**Homebrew** (macOS / Linux):

```sh
brew install phrase-cli
```

**asdf** (version-managed, useful for repos that pin a CLI version):

```sh
asdf plugin add phrase https://github.com/phrase/asdf-phrase
asdf install phrase latest
asdf set phrase latest    # or: asdf set -u phrase <version>  to write a .tool-versions
```

**GitHub releases** (any OS — pick the matching binary):
<https://github.com/phrase/phrase-cli/releases/latest>

Verify with `phrase --version`. Full install guide: <https://support.phrase.com/hc/en-us/articles/5784093863964>.

## `phrase push`

| Flag | Short | What it does |
|---|---|---|
| `--wait` | `-w` | Block until uploads are processed. |
| `--cleanup` | `-c` | Same as `delete_unmentioned_keys: true`. Deletes keys not in any uploaded file. |
| `--branch <name>` | `-b` | Push to a specific Phrase branch. |
| `--use-local-branch-name` | | Use the current git/hg branch name as the Phrase branch. |
| `--tag <name>` | | Apply a tag to every key in this upload. |
| `--token <t>` | `-t` | Override `access_token`. |
| `--host <url>` | `-h` | Override `host`. |

## `phrase pull`

| Flag | Short | What it does |
|---|---|---|
| `--branch <name>` | `-b` | Pull from a specific Phrase branch. |
| `--use-local-branch-name` | | Use the current git/hg branch name. |
| `--async` | `-a` | Asynchronous downloads. Useful for projects with many locales. |
| `--cache` | | Conditional downloads via ETags (sync mode only). |
| `--parallel` | `-p` | Download up to 4 locales in parallel (sync mode only). |
| `--token <t>` | `-t` | Override `access_token`. |
| `--host <url>` | `-h` | Override `host`. |

`--cache` and `--parallel` cannot be combined with `--async`.

## Environment variables

| Variable | What it does |
|---|---|
| `PHRASE_ACCESS_TOKEN` | Provides the access token. The recommended way — never put the token in `.phrase.yml`. |
| `PHRASE_PROJECT_ID` | Override `project_id`. |
| `PHRASE_HOST` | Override `host`. |

Precedence (highest first): CLI flag → environment variable → config file → built-in default.
