export class NetworkBaseError extends Error {
  statusCode: number
  constructor(statusCode: number, message: string) {
    super(message)
    Object.setPrototypeOf(this, new.target.prototype)
    this.name = Error.name
    this.statusCode = statusCode
    if (Error.captureStackTrace !== undefined) {
      Error.captureStackTrace(this)
    }
  }
}

export class CannotConnectNetworkError extends Error {
  constructor(message: string) {
    super(message)
    Object.setPrototypeOf(this, new.target.prototype)
    this.name = Error.name
    if (Error.captureStackTrace !== undefined) {
      Error.captureStackTrace(this)
    }
  }
}

export class BadRequestError extends NetworkBaseError {}
export class ForbiddenError extends NetworkBaseError {}
export class AuthenticationError extends NetworkBaseError {}
export class NotFoundError extends NetworkBaseError {}
export class AlreadyExistError extends NetworkBaseError {}
export class InternalServerError extends NetworkBaseError {}

// custom error
export class EmailAlreadyInUseError extends NetworkBaseError {}
export class InvalidEmailUseError extends NetworkBaseError {}
export class InvalidAddressError extends NetworkBaseError {}
export class EmailNotVerifiedError extends NetworkBaseError {}
