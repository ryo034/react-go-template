import { Jwt, Otp } from "~/domain/auth"
import { Email } from "~/domain/shared"
import { PromiseResult } from "~/infrastructure/shared"

export interface AuthRepository {
  startWithEmail(email: Email): PromiseResult<Otp, Error>
  verifyOtp(email: Email, otp: Otp): PromiseResult<Jwt, Error>
  signInWithCustomToken(jwt: Jwt): PromiseResult<null, Error>
}
