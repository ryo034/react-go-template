# Shared UI Component

## How to add component

1. Add component

example.

```shell
> corepack pnpm dlx shadcn-ui@latest add separator
```

2. Add export to src/components/index.ts

```ts
...
export * from "./ui/separator"
...
```

3. build component

```shell
> corepack pnpm run build
```

4. install

run on root directory

```shell
> corepack pnpm install
```
