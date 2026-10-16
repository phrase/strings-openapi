# strings-openapi

## Spec changes workflow

1. Before adding a path file or endpoint docs, grep `paths.yaml`, `doc/compiled.json` and `paths/` for the endpoint to confirm it is not already documented under another file name or tag.
2. Edit the source YAML.
3. Run `make lint` (runs `swagger-cli validate`, `make bundle`, then `lint.sh`). Do not run a second manual `swagger-cli bundle` afterwards.
4. `git diff` to confirm scope.
5. Commit the source change together with the regenerated `doc/compiled.json` (generated but committed). `tmp/compiled.yaml` is gitignored; never stage it.

## Finding endpoints

`paths.yaml` references path files, not literal URLs. Find where an endpoint is wired up with e.g. `grep -n "search.yaml" paths.yaml`.

## Consistency

When adding or editing a request-body property for a resource, apply it to both `paths/<resource>/create.yaml` and `paths/<resource>/update.yaml`, and to the shared response schema under `schemas/`, so POST/PATCH/GET contracts do not drift.

## Breaking changes CI

`.github/workflows/breaking-changes.yml` diffs the compiled spec against the base branch. Adding or reordering enum values, or altering shared response schemas such as `schemas/project_details.yaml`, can fail it. Inspect its diff output before pushing; if it fails on a purely additive enum/field change, explain the additive nature in the PR body instead of re-pushing the same change.

## Commits and PR titles

PR titles are used to generate client library changelogs and version bumps. Follow the "PR titles" section of the README (https://github.com/phrase/strings-openapi#pr-titles): `change_type(affected_library): description #STRINGS-TICKET`, e.g. `docs(API): clarify field description #STRINGS-1234`.

- `change_type`: `feat` (new endpoints, properties, query parameters), `fix` (schema fixes), `docs` (description-only changes, no schema change), `build` (build system or CI changes).
- `affected_library`: most commonly `API`; schema changes always use `API`.
