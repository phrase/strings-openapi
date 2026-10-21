# strings-openapi

## Spec changes workflow

1. Before adding a path file or endpoint docs, grep `paths.yaml`, `doc/compiled.json` and `paths/` for the endpoint to confirm it is not already documented under another file name or tag.
2. Edit the source YAML.
3. Run `make lint` (runs `swagger-cli validate`, `make bundle`, then `lint.sh`). Do not run a second manual `swagger-cli bundle` afterwards.
4. `git diff` to confirm scope.
5. Commit the source change together with the regenerated `doc/compiled.json` (generated but committed). `tmp/compiled.yaml` is gitignored; never stage it.
6. `bash ./lint.sh` reads the gitignored `tmp/compiled.yaml`; `make lint` or `make bundle` generates it first. Do not treat a failure on a missing file as your own regression. For an additive change, check `git diff doc/compiled.json` is purely additive.

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
- The `Validate PR title` check (`.github/workflows/check-pr-title.yml`) enforces this prefix; compose the title correctly on the first push. Use `fix(API):` for corrections to an existing spec (missing params, wrong fields) rather than `feat(API):`.

## Documenting endpoints

- Never infer response attributes from sibling endpoints or plausibility. Derive them from the live API (e.g. `curl -s https://api.phrase.com/v2/<path>`) or the actual serializer. For any field or query parameter you could not verify that way, say so in the PR body as inferred, not confirmed.
- Name tags and summaries the way TMS already names the same concept (e.g. tag `Supported Languages`, summary `List supported languages`).
- Keep endpoint descriptions limited to authentication requirements; do not document rate limiting or other operational behaviour that may change.
- If an endpoint's response shape or behaviour changes depending on query params, or two paths (e.g. `/accounts/.../members` vs `/projects/.../members`) use different serializers, do not document a single merged shape. Keep the schemas separate and flag the ambiguity in the PR body for a human decision.

## Copy-pasted descriptions

Before editing a description, field doc, or parameter block, run one repo-wide Grep for the same phrasing/field across `paths/` and `schemas/`; these are routinely copy-pasted across sibling endpoints. Fix every occurrence in the same PR or list the ones deliberately left alone in the PR body.
