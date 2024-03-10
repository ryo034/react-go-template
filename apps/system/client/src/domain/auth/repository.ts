import type { CustomToken, Me, Otp, ReceivedInvitation } from "~/domain"
import type { Email } from "~/domain/shared"
import type { PromiseResult } from "~/infrastructure/shared"

export interface AuthRepository {
  startWithEmail(email: Email): PromiseResult<null, Error>
  startWithGoogle(): PromiseResult<null, Error>
  authByOAuth(): PromiseResult<Me, Error>
  verifyOtp(email: Email, otp: Otp): PromiseResult<CustomToken, Error>
  signInWithCustomToken(customToken: CustomToken): PromiseResult<null, Error>
  findInvitationByToken(token: string): PromiseResult<ReceivedInvitation, Error>
  proceedInvitationByEmail(token: string, email: Email): PromiseResult<null, Error>
  proceedInvitationByOAuth(token: string): PromiseResult<Me, Error>
}
