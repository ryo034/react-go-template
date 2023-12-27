import { Email, Password } from "~/domain/shared"

export interface MeUseCaseInput {
  login(email: Email, password: Password): Promise<Error | null>
  signOut(): Promise<Error | null>
  checkEmailVerified(): Promise<Error | null>
  verifyEmail(): Promise<Error | null>
  sendEmailVerification(): Promise<Error | null>
  find(): Promise<Error | null>
}
