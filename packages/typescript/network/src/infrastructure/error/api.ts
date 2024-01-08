import { CannotConnectNetworkError } from "./network"

export class UnknownError extends Error {
  constructor(message: string) {
    super(message)
    Object.setPrototypeOf(this, new.target.prototype)
    this.name = "UnknownError"
    if (Error.captureStackTrace !== undefined && typeof Error.captureStackTrace === "function") {
      Error.captureStackTrace(this, UnknownError)
    }
  }
}

export interface ApiErrorHandlerInterface {
  adapt(e: unknown): Error
}

/**
 * ApiErrorHandler
 * @param customErrorChecker
 *
 * If this function returns null, UnknownError is returned.
 * Pass the implementation that returns a custom error according to the error code defined in each service to customErrorChecker
 *
 * example:
 * ```ts
 * const apiErrorHandler = new ApiErrorHandler(err => {
 *   return CustomCodeNetworkError.create(err)
 * })
 * ```
 */
export class ApiErrorHandler {
  constructor(readonly customErrorChecker: (err: unknown) => Error | null) {}

  adapt(e: unknown) {
    if (navigator !== undefined && navigator.onLine !== undefined && navigator.onLine === false) {
      return new CannotConnectNetworkError("")
    }

    const customError = this.customErrorChecker(e)
    if (customError) {
      return customError
    }

    if (e instanceof Error) {
      return e
    }
    if (typeof e === "string") {
      return new UnknownError(e)
    }
    return new UnknownError("unknown error")
  }
}
