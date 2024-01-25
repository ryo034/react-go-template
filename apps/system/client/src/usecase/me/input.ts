import { Email, Password } from "~/domain/shared"

export interface MeUseCaseInput {
  login(): Promise<Error | null>
  signOut(): Promise<Error | null>
  find(): Promise<Error | null>
}
