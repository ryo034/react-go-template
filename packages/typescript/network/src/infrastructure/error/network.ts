export class NetworkBaseError extends Error {
  statusCode: number
  constructor(statusCode: number, message: string) {
    super(message)
    Object.setPrototypeOf(this, new.target.prototype)
    this.name = "NetworkBaseError"
    this.statusCode = statusCode
    if (Error.captureStackTrace !== undefined && typeof Error.captureStackTrace === "function") {
      Error.captureStackTrace(this, NetworkBaseError)
    }
  }
}

export class CannotConnectNetworkError extends Error {
  constructor(message: string) {
    super(message)
    Object.setPrototypeOf(this, new.target.prototype)
    this.name = "CannotConnectNetworkError"
    if (Error.captureStackTrace !== undefined && typeof Error.captureStackTrace === "function") {
      Error.captureStackTrace(this, CannotConnectNetworkError)
    }
  }
}

export class BadRequestError extends NetworkBaseError {}
export class ForbiddenError extends NetworkBaseError {}
export class AuthenticationError extends NetworkBaseError {}
export class NotFoundError extends NetworkBaseError {}
export class AlreadyExistError extends NetworkBaseError {}
export class RequestTimeoutError extends NetworkBaseError {}
export class InternalServerError extends NetworkBaseError {}

export const isCannotConnectNetworkError = (e: unknown): e is CannotConnectNetworkError =>
  e instanceof CannotConnectNetworkError
export const isBadRequestError = (e: unknown): e is BadRequestError => e instanceof BadRequestError
export const isForbiddenError = (e: unknown): e is ForbiddenError => e instanceof ForbiddenError
export const isAuthenticationError = (e: unknown): e is AuthenticationError => e instanceof AuthenticationError
export const isNotFoundError = (e: unknown): e is NotFoundError => e instanceof NotFoundError
export const isAlreadyExistError = (e: unknown): e is AlreadyExistError => e instanceof AlreadyExistError
export const isRequestTimeoutError = (e: unknown): e is RequestTimeoutError => e instanceof RequestTimeoutError
export const isInternalServerError = (e: unknown): e is InternalServerError => e instanceof InternalServerError
