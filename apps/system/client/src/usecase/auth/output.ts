import { ReceivedInvitation } from "~/domain"
import { Email } from "~/domain/shared"

export interface AuthUseCaseOutput {
  setEmail: (v: Email) => void
  setIsLoading: (v: boolean) => void
  clear: () => void
  setReceivedInvitation(v: ReceivedInvitation): void
  setReceivedInvitationIsLoading: (v: boolean) => void
  setIsInvitationProcessing: (v: boolean) => void
}
