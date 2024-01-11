.PHONY: run
run:
	@docker-compose up -d

.PHONY: rerun
rerun:
	@docker-compose stop api
	@docker-compose up -d api

.PHONY: update-proto-package
update-proto-package:
	@go clean --modcache
	@go mod tidy

.PHONY: update-private-package
update-private-package:
	@go get -v github.com/ryo034/react-go-template/packages/go
	@go mod tidy

.PHONY: generate
generate:
	@go generate ./...

.PHONY: format
format:
	@go fmt .
	@which goimports > /dev/null || go install golang.org/x/tools/cmd/goimports@latest
	@goimports -d $$(find . -type f -name '*.go' -not -path "./schema/*")

.PHONY: test
test:
	go test $(go list ./... | grep -v /driver/)

.PHONY: test-integration
test-integration:
	go test ./driver/...
