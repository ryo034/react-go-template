import { Email } from "~/domain"
import { Otp } from "~/domain/auth"

export interface AuthUseCaseOutput {
  set: (v: Otp) => void
  setEmail: (v: Email) => void
  setIsLoading: (v: boolean) => void
  clear: () => void
}
