import { Email, Password } from "~/domain/shared"

export interface MeUseCaseInput {
  login(data: MeUseCaseLoginInputData): Promise<Error | null>
  signOut(): Promise<Error | null>
  checkEmailVerified(): Promise<Error | null>
  verifyEmail(): Promise<Error | null>
  sendEmailVerification(): Promise<Error | null>
  find(): Promise<Error | null>
}

export interface MeUseCaseLoginInputData {
  email: Email
  password: Password
}
