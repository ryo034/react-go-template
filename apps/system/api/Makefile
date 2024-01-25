.PHONY: run
run:
	@docker-compose up -d

.PHONY: rerun
rerun:
	@docker-compose stop api
	@docker-compose up -d api

.PHONY: update-package
update-package:
	@go clean --modcache
	@go get -u ./...
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
	go test -v -count=1 -race -shuffle=on ./...

.PHONY: test-integration
test-integration:
	go test -v -count=1 -race -shuffle=on ./... -tags='testcontainers'
