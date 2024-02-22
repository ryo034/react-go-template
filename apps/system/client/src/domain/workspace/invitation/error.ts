export class AlreadyAcceptedInvitationError extends Error {
  constructor(message: string) {
    super(message)
    Object.setPrototypeOf(this, new.target.prototype)
    this.name = Error.name
    if (Error.captureStackTrace !== undefined) {
      Error.captureStackTrace(this)
    }
  }
}

export class AlreadyRevokeInvitationError extends Error {
  constructor(message: string) {
    super(message)
    Object.setPrototypeOf(this, new.target.prototype)
    this.name = Error.name
    if (Error.captureStackTrace !== undefined) {
      Error.captureStackTrace(this)
    }
  }
}

export class AlreadyExpiredInvitationError extends Error {
  constructor(message: string) {
    super(message)
    Object.setPrototypeOf(this, new.target.prototype)
    this.name = Error.name
    if (Error.captureStackTrace !== undefined) {
      Error.captureStackTrace(this)
    }
  }
}

export const isAlreadyAcceptedInvitationError = (err: unknown): err is AlreadyAcceptedInvitationError =>
  err instanceof AlreadyAcceptedInvitationError

export const isAlreadyRevokeInvitationError = (err: unknown): err is AlreadyRevokeInvitationError =>
  err instanceof AlreadyRevokeInvitationError

export const isAlreadyExpiredInvitationError = (err: unknown): err is AlreadyExpiredInvitationError =>
  err instanceof AlreadyExpiredInvitationError
