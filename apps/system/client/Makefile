.PHONY: setup
setup:
	@cp .env.local.sample .env

.PHONY: package-update-latest
package-update-latest:
	@ncu -u && rm -rf node_modules && corepack pnpm install
