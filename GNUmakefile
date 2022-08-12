SWAGGER_URL ?= https://rootly-heroku.s3.amazonaws.com/swagger/v1/swagger.json
TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=hashicorp.com
NAMESPACE=eduW
NAME=rootly
BINARY=terraform-provider-${NAME}
VERSION=0.1.1
OS_ARCH=darwin_amd64

default: testacc

# Run acceptance tests
.PHONY: testacc schema build release install test docs
build:
	go build -o ${BINARY}

docs:
	@go get -v github.com/hashicorp/terraform-plugin-docs/...
	tfplugindocs generate
	cp docs/index.md README.md
	find ./docs/resources/workflow_task_*.md -type f -print0 | xargs -0 sed -i '' 's/subcategory:$$/subcategory: Workflow Tasks/g'
	find ./docs/resources/workflow_*.md -type f -print0 | xargs -0 sed -i '' 's/subcategory:$$/subcategory: Workflows/g'

release:
	goreleaser release --rm-dist --snapshot --skip-publish  --skip-sign

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

tools:
	cd providerlint && go install .
	cd tools && go install github.com/bflad/tfproviderdocs
	cd tools && go install github.com/client9/misspell/cmd/misspell
	cd tools && go install github.com/golangci/golangci-lint/cmd/golangci-lint
	cd tools && go install github.com/katbyte/terrafmt
	cd tools && go install github.com/terraform-linters/tflint
	cd tools && go install github.com/pavius/impi/cmd/impi
	cd tools && go install github.com/hashicorp/go-changelog/cmd/changelog-build

test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=5m -parallel=4

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

schema:
	cd schema && curl $(SWAGGER_URL) -o swagger.json
	cd schema && oapi-codegen --config=oapi-config.yml swagger.json
	cd schema && rm swagger.json
