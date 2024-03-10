import type { Result } from "true-myth"

export type PromiseResult<T, E extends Error> = Promise<Result<T, E>>
