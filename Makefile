run:
	make lint
	make bundle
	make ruby
	make go
	make typescript
	make python
	make java
	make php
	make cli
	@echo Client libraries generated, check clients directory.
lint:
	npx swagger-cli validate main.yaml
	make bundle
	bash ./lint.sh
bundle:
	npx swagger-cli bundle -t json -w 300 main.yaml > doc/compiled.json
	npx swagger-cli bundle -t yaml -w 300 main.yaml > tmp/compiled.yaml
watch_bundle:
	make lint
	npx swagger-cli bundle -t json -w 300 main.yaml > doc/compiled.json
# The generator renders missing examples as "null" and object params as "{ ... }",
# which mustache can't detect; rewrite them so the doc code samples are valid Ruby.
ruby:
	openapi-generator-cli generate -i tmp/compiled.yaml -g ruby -o clients/ruby -c ./openapi-generator/ruby_lang.yaml
	perl -pi -e "s/: (?:null|'null'|Time\.parse\('null'\))(?=[,)]\$$)/: nil/; s/\{ \.\.\. \}/{}/g" clients/ruby/docs/*.md
go:
	openapi-generator-cli generate -i tmp/compiled.yaml -g go -o clients/go -c ./openapi-generator/go_lang.yaml --global-property apiTests=false,modelTests=false
	go install golang.org/x/tools/cmd/goimports@v0.24.0
	goimports -w clients/go
	cd clients/go && go mod tidy
typescript:
	openapi-generator-cli generate -i tmp/compiled.yaml -g typescript-fetch -o clients/typescript -c ./openapi-generator/typescript_lang.yaml
python:
	openapi-generator-cli generate -i tmp/compiled.yaml -g python -o clients/python -c ./openapi-generator/python_lang.yaml
java:
	openapi-generator-cli generate -i tmp/compiled.yaml -g java -o clients/java -c ./openapi-generator/java_lang.yaml
php:
	openapi-generator-cli generate -i tmp/compiled.yaml -g php -o clients/php -c ./openapi-generator/php_lang.yaml
cli:
	openapi-generator-cli generate -i tmp/compiled.yaml -g go -o tmp/cli -c ./openapi-generator/cli_lang.yaml -e handlebars
	cp tmp/cli/api_* clients/cli/cmd/
	cp tmp/cli/README.md clients/cli/
	go install golang.org/x/tools/cmd/goimports@v0.24.0
	goimports -w clients/cli
	cd clients/cli && go mod tidy
