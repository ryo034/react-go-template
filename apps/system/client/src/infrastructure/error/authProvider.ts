export class AuthProviderCustomError extends Error {
  constructor(message: string) {
    super(message)
    Object.setPrototypeOf(this, new.target.prototype)
    this.name = Error.name
    if (Error.captureStackTrace !== undefined) {
      Error.captureStackTrace(this)
    }
  }
}

export class AuthProviderCurrentUserNotFoundError extends AuthProviderCustomError {}
export class AuthProviderNotFoundCurrentLibraryIdInCustomClaims extends AuthProviderCustomError {}

export class AuthProviderMissingEmail extends AuthProviderCustomError {}
export class AuthProviderUserNotFoundError extends AuthProviderCustomError {}
export class AuthProviderInvalidPhoneNumberError extends AuthProviderCustomError {}
export class AuthProviderWrongPasswordError extends AuthProviderCustomError {}
export class AuthProviderInvalidPasswordError extends AuthProviderCustomError {}
export class AuthProviderEmailAlreadyInUseError extends AuthProviderCustomError {}
export class AuthProviderIdTokenExpiredError extends AuthProviderCustomError {}
export class AuthProviderIdTokenRevokedError extends AuthProviderCustomError {}
export class AuthProviderUnverifiedEmailError extends AuthProviderCustomError {}
export class AuthProviderInvalidContentTypeError extends AuthProviderCustomError {}
export class AuthProviderInternalError extends AuthProviderCustomError {}
