.PHONY: test-all
test-all:
	@corepack pnpm run test:stateful
	@corepack pnpm run test:stateless

.PHONY: package-update-latest
package-update-latest:
	@ncu -u && rm -rf node_modules && corepack pnpm install
