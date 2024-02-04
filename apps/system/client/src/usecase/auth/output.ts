import { Email } from "~/domain"

export interface AuthUseCaseOutput {
  setEmail: (v: Email) => void
  setIsLoading: (v: boolean) => void
  clear: () => void
}
