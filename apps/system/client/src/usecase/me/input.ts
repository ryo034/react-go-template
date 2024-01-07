import { Email, Password } from "~/domain/shared"

export interface MeUseCaseInput {
  login(): Promise<Error | null>
  signUp(d: MeUseCaseSignUpInputData): Promise<Error | null>
  signOut(): Promise<Error | null>
  find(): Promise<Error | null>
}

export interface MeUseCaseSignUpInputData {
  firstName: string
  lastName: string
}
