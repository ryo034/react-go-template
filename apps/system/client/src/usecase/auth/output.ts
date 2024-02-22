import { Email, ReceivedInvitation } from "~/domain"

export interface AuthUseCaseOutput {
  setEmail: (v: Email) => void
  setIsLoading: (v: boolean) => void
  clear: () => void
  setReceivedInvitation(v: ReceivedInvitation): void
  setReceivedInvitationIsLoading: (v: boolean) => void
  setIsInvitationProcessing: (v: boolean) => void
}
