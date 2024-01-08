# https://github.com/awslabs/git-secrets
.PHONY: setup-git-secrets
setup-git-secrets:
	@brew install git-secrets

.PHONY: setup
setup:
	@make setup-git-secrets

.PHONY: init-system
init-react:
	@cp ./apps/system/client/.env.local.sample ./apps/system/client/.env.local

.PHONY: init
init:
	@cp .env.sample .env
	@make init-react
	@echo "Success!🎉"

.PHONY: start
start:
	docker-compose up -d api

.PHONY: restart
restart:
	docker-compose down
	docker-compose up -d api

.PHONY: restart-statefull
restart-statefull:
	docker-compose down
	rm -rf container/database/postgresql/write/.data
	rm -rf container/database/postgresql/read/.data
	docker-compose up -d api

.PHONY: restart-db
restart-db:
	docker-compose rm -fsv main_db_replica
	docker-compose rm -fsv main_db_primary
	rm -rf container/database/postgresql/primary/.data
	rm -rf container/database/postgresql/replica/.data
	docker-compose up -d main_db_replica

# ====================
#  Update packages
# ====================

.PHONY: update-all-typescript-package
update-all-typescript-package:
	@cd ./packages/typescript/ui && ncu -u
	@cd ./packages/typescript/network && ncu -u
	@cd ./apps/system/client && ncu -u
	@corepack pnpm install -r

.PHONY: update-all-go-package
update-all-go-package:
	@cd ./apps/system/api && make update-private-package

# ============
#  Schema
# ============

# ============
#  OpenAPI
# ============

.PHONY: merge-system-openapi
merge-system-openapi:
	@rm -rf ./schema/api/system/openapi/openapi.yaml
	@docker build -f ./container/schema/openapi/swagger-merger/Dockerfile -t swagger-merger-image .
	@docker run --rm -v ./schema/api/system/openapi:/swagger swagger-merger-image swagger-merger -i /swagger/index.yaml -o /swagger/openapi.yaml

.PHONY: gen-system-api-openapi
gen-system-api-openapi:
	@docker run --rm \
		-v ".:/workspace" ghcr.io/ogen-go/ogen:latest \
		-package openapi \
		-target workspace/apps/system/api/schema/openapi \
		-clean workspace/schema/api/system/openapi/openapi.yaml

.PHONY: gen-system-client-openapi
gen-system-client-openapi:
	@rm -rf ./apps/system/client/src/generated/schema/openapi
	@docker build -f ./container/schema/openapi/openapi-typescript/Dockerfile -t openapi-typescript-codegen-tmp .
	@docker run --rm -v .:/app \
		openapi-typescript-codegen-tmp \
		/app/schema/api/system/openapi/openapi.yaml -o /app/apps/system/client/src/generated/schema/openapi/systemApi.ts

.PHONY: gen-system-openapi
gen-system-openapi:
	@make merge-system-openapi
	@make gen-system-client-openapi
	@make gen-system-api-openapi

.PHONY: gen-openapi
gen-openapi:
	@make gen-system-openapi

# ====================
#  Technical document
# ====================

# Function to convert string to kebab-case
define to_kebab_case
$(shell echo '$1' | sed -E 's/([a-z])([A-Z])/\1-\2/g' | tr A-Z a-z | sed -E 's/[^a-z0-9-]+/-/g' | sed -E 's/^-+|-+$$//g')
endef

# Command to generate a new ADR
# Usage: make gen-adr TARGET=path/to/directory TITLE=title-of-document
.PHONY: gen-adr
gen-adr:
	@if [ -z "$(TARGET)" ]; then \
		echo "Error: TARGET is not specified. Usage: make gen-adr TARGET=path/to/directory TITLE=title-of-document"; \
		exit 1; \
	fi
	@if [ -z "$(TITLE)" ]; then \
		echo "Error: TITLE is not specified. Usage: make gen-adr TARGET=path/to/directory TITLE=title-of-document"; \
		exit 1; \
	fi
	$(eval SUGGESTED_TITLE := $(call to_kebab_case,$(TITLE)))
	@if ! echo "$(TITLE)" | grep -qE '^[a-z]+(-[a-z]+)*$$'; then \
		echo "Error: TITLE must be in kebab-case (e.g., this-is-kebab-case)."; \
		echo "Suggested command: \n\n> make adr TARGET=$(TARGET) TITLE=$(SUGGESTED_TITLE)\n"; \
		exit 1; \
	fi
	@mkdir -p $(TARGET)
	@cp ./docs/tech/adr/template.md $(TARGET)/$(shell date +%Y%m%d)-$(TITLE).md
	@echo "ADR created at $(TARGET)/$(shell date +%Y%m%d)-$(TITLE).md"

# Command to generate a new Design Doc
# Usage: make gen-design-doc TARGET=path/to/directory TITLE=title-of-document
.PHONY: gen-design-doc
gen-design-doc:
	@if [ -z "$(TARGET)" ]; then \
		echo "Error: TARGET is not specified. Usage: make gen-design-doc TARGET=path/to/directory TITLE=title-of-document"; \
		exit 1; \
	fi
	@if [ -z "$(TITLE)" ]; then \
		echo "Error: TITLE is not specified. Usage: make gen-design-doc TARGET=path/to/directory TITLE=title-of-document"; \
		exit 1; \
	fi
	$(eval SUGGESTED_TITLE := $(call to_kebab_case,$(TITLE)))
	@if ! echo "$(TITLE)" | grep -qE '^[a-z]+(-[a-z]+)*$$'; then \
		echo "Error: TITLE must be in kebab-case (e.g., this-is-kebab-case)."; \
		echo "Suggested command: \n\n> make gen-design-doc TARGET=$(TARGET) TITLE=$(SUGGESTED_TITLE)\n"; \
		exit 1; \
	fi
	@mkdir -p $(TARGET)
	@cp ./docs/tech/design-doc/template.md $(TARGET)/$(shell date +%Y%m%d)-$(TITLE).md
	@echo "Design document created at $(TARGET)/$(shell date +%Y%m%d)-$(TITLE).md"
