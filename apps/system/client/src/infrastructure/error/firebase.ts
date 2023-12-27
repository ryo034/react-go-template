import { FirebaseError } from "firebase/app"
import {
  AuthProviderEmailAlreadyInUseError,
  AuthProviderIdTokenExpiredError,
  AuthProviderIdTokenRevokedError,
  AuthProviderInvalidPasswordError,
  AuthProviderInvalidPhoneNumberError,
  AuthProviderMissingEmail,
  AuthProviderUnverifiedEmailError,
  AuthProviderUserNotFoundError,
  AuthProviderWrongPasswordError
} from "~/infrastructure/error"
import { InternalServerError } from "~/infrastructure/error/network"
import { HttpStatusCode } from "~/infrastructure/error/statusCode"

const firebaseErrorCode = {
  missingEmailError: "auth/missing-email",
  invalidPhoneNumber: "auth/invalid-phone-number",
  userNotFoundError: "auth/user-not-found",
  invalidPasswordError: "auth/invalid-password",
  wrongPasswordError: "auth/wrong-password",
  emailAlreadyInUseError: "auth/email-already-in-use",
  idTokenExpiredError: "auth/id-token-expired",
  idTokenRevokedError: "auth/id-token-revoked",
  multiFactorAuthRequired: "auth/multi-factor-auth-required",
  unverifiedEmail: "auth/unverified-email"
}

export class FirebaseErrorAdapter {
  static create(err: FirebaseError): Error {
    switch (err.code) {
      case firebaseErrorCode.missingEmailError:
        return new AuthProviderMissingEmail("missing email")
      case firebaseErrorCode.invalidPhoneNumber:
        return new AuthProviderInvalidPhoneNumberError("invalid phone number")
      case firebaseErrorCode.userNotFoundError:
        return new AuthProviderUserNotFoundError("user not found")
      case firebaseErrorCode.invalidPasswordError:
        return new AuthProviderInvalidPasswordError("invalid password")
      case firebaseErrorCode.wrongPasswordError:
        return new AuthProviderWrongPasswordError("invalid password")
      case firebaseErrorCode.emailAlreadyInUseError:
        return new AuthProviderEmailAlreadyInUseError("email already in use")
      case firebaseErrorCode.idTokenExpiredError:
        return new AuthProviderIdTokenExpiredError("id token expired")
      case firebaseErrorCode.idTokenRevokedError:
        return new AuthProviderIdTokenRevokedError("id token revoked")
      case firebaseErrorCode.unverifiedEmail:
        return new AuthProviderUnverifiedEmailError("unverified email")
      default:
        return new InternalServerError(HttpStatusCode.INTERNAL_SERVER_ERROR, "Internal Server Error")
    }
  }
}
