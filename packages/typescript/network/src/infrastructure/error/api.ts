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

/**
 * ApiErrorHandler
 * @param customErrorChecker
 *
 * If this function returns null, UnknownError is returned.
 * Used for error handling in libraries such as Axios and ky
 * Pass the implementation that returns a custom error according to the error code defined in each service to customErrorChecker
 *
 * example:
 * ```ts
 * const apiErrorHandler = new ApiErrorHandler(err => {
 *   const res = NetworkErrorInterpreter.adapt(err)
 *   if (res) {
 *     return res
 *   }
 *   return AxiosNetworkError.create(err)
 * })
 * ```
 */
export class ApiErrorHandler {
  constructor(readonly customErrorChecker: (err: unknown) => Error | null) {}

  adapt(origError: unknown) {
    if (navigator !== undefined && navigator.onLine !== undefined && navigator.onLine === false) {
      return new CannotConnectNetworkError("")
    }

    const customError = this.customErrorChecker(origError)
    if (customError) {
      return customError
    }

    if (origError instanceof Error) {
      return origError
    }
    if (typeof origError === "string") {
      return new UnknownError(origError)
    }
    return new UnknownError("unknown error")
  }
}
