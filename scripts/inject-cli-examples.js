#!/usr/bin/env node
'use strict';

// Inject generated CLI v2 code samples into a compiled OpenAPI bundle.
//
// Path files no longer carry the CLI v2 x-code-sample; it is generated into
// examples/cli.yaml (keyed by operationId) by the cli_examples.handlebars
// template during `make cli`, and appended to each operation's x-code-samples
// here, after bundling. Supports both JSON and YAML bundles.
//
// Usage: node scripts/inject-cli-examples.js <compiled.json|compiled.yaml> [...]

const fs = require('fs');
const path = require('path');
const yaml = require('js-yaml');

const EXAMPLES = path.join(__dirname, '..', 'examples', 'cli.yaml');
const HTTP_METHODS = ['get', 'post', 'put', 'patch', 'delete', 'options', 'head'];

function main() {
  const targets = process.argv.slice(2);
  if (targets.length === 0) {
    console.error('usage: inject-cli-examples.js <bundle> [...]');
    process.exit(1);
  }

  const examples = yaml.load(fs.readFileSync(EXAMPLES, 'utf8')) || {};

  for (const file of targets) {
    const isJson = file.endsWith('.json');
    const raw = fs.readFileSync(file, 'utf8');
    const doc = isJson ? JSON.parse(raw) : yaml.load(raw);

    let injected = 0;
    let missing = 0;
    for (const item of Object.values(doc.paths || {})) {
      for (const method of HTTP_METHODS) {
        const op = item[method];
        if (!op || !op.operationId) continue;
        const ex = examples[op.operationId];
        if (!ex) { missing++; continue; }

        const samples = (op['x-code-samples'] = op['x-code-samples'] || []);
        // replace an existing CLI v2 entry if present, else append
        const i = samples.findIndex((s) => s && s.lang === ex.lang);
        if (i >= 0) samples[i] = ex; else samples.push(ex);
        injected++;
      }
    }

    const out = isJson
      ? JSON.stringify(doc, null, 2) + '\n'
      : yaml.dump(doc, { lineWidth: 300, noRefs: true });
    fs.writeFileSync(file, out);
    console.error(`${file}: injected ${injected}, no example for ${missing}`);
  }
}

main();
