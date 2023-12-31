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

# gomockでモックを生成
.PHONY: mock
mock:
	go generate ./...

.PHONY: format
format:
	@go fmt .
	@which goimports > /dev/null || go install golang.org/x/tools/cmd/goimports@latest
	@goimports -d $$(find . -type f -name '*.go' -not -path "./schema/*")

.PHONY: test
test:
	@go test ./infrastructure/... ./interface/... ./domain/... ./usecase/... ./util/...
