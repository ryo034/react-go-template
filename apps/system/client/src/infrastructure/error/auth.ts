export class AuthBaseError extends Error {
  constructor(message?: string) {
    super(message)
    Object.setPrototypeOf(this, new.target.prototype)
    this.name = Error.name
    if (Error.captureStackTrace !== undefined) {
      Error.captureStackTrace(this)
    }
  }
}

export class UnAuthenticatedError extends AuthBaseError {}
