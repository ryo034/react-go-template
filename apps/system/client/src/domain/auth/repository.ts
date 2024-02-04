import { CustomToken, Otp } from "~/domain/auth"
import { Email } from "~/domain/shared"
import { PromiseResult } from "~/infrastructure/shared"

export interface AuthRepository {
  startWithEmail(email: Email): PromiseResult<null, Error>
  verifyOtp(email: Email, otp: Otp): PromiseResult<CustomToken, Error>
  signInWithCustomToken(customToken: CustomToken): PromiseResult<null, Error>
}
