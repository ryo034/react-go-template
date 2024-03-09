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

export class BadRequestError extends NetworkBaseError {
  constructor(statusCode: number, message: string) {
    super(statusCode, message)
    this.name = "BadRequestError"
  }
}
export class ForbiddenError extends NetworkBaseError {
  constructor(statusCode: number, message: string) {
    super(statusCode, message)
    this.name = "ForbiddenError"
  }
}
export class AuthenticationError extends NetworkBaseError {
  constructor(statusCode: number, message: string) {
    super(statusCode, message)
    this.name = "AuthenticationError"
  }
}
export class NotFoundError extends NetworkBaseError {
  constructor(statusCode: number, message: string) {
    super(statusCode, message)
    this.name = "NotFoundError"
  }
}
export class AlreadyExistError extends NetworkBaseError {
  constructor(statusCode: number, message: string) {
    super(statusCode, message)
    this.name = "AlreadyExistError"
  }
}
export class RequestTimeoutError extends NetworkBaseError {
  constructor(statusCode: number, message: string) {
    super(statusCode, message)
    this.name = "RequestTimeoutError"
  }
}
export class InternalServerError extends NetworkBaseError {
  constructor(statusCode: number, message: string) {
    super(statusCode, message)
    this.name = "InternalServerError"
  }
}

export const isCannotConnectNetworkError = (e: unknown): e is CannotConnectNetworkError =>
  e instanceof CannotConnectNetworkError
export const isBadRequestError = (e: unknown): e is BadRequestError => e instanceof BadRequestError
export const isForbiddenError = (e: unknown): e is ForbiddenError => e instanceof ForbiddenError
export const isAuthenticationError = (e: unknown): e is AuthenticationError => e instanceof AuthenticationError
export const isNotFoundError = (e: unknown): e is NotFoundError => e instanceof NotFoundError
export const isAlreadyExistError = (e: unknown): e is AlreadyExistError => e instanceof AlreadyExistError
export const isRequestTimeoutError = (e: unknown): e is RequestTimeoutError => e instanceof RequestTimeoutError
export const isInternalServerError = (e: unknown): e is InternalServerError => e instanceof InternalServerError
