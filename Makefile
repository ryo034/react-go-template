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

.PHONY: update-all-typescript-package
update-all-typescript-package:
	@cd ./packages/typescript/ui && ncu -u
	@cd ./apps/system/client && ncu -u
	@cd ./apps/system/test/e2e && ncu -u
	@cd ./apps/system/test/api && ncu -u
	@pnpm install -r

.PHONY: update-all-go-package
update-all-go-package:
	@cd ./apps/system/api && make update-private-package
