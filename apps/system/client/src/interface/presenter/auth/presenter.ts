import type { ReceivedInvitation } from "~/domain"
import type { Email } from "~/domain/shared"
import type { AuthStoreType, ReceivedInvitationStoreType } from "~/store"
import type { AuthUseCaseOutput } from "~/usecase/auth"

export class AuthPresenter implements AuthUseCaseOutput {
  constructor(
    private readonly store: AuthStoreType,
    private readonly receivedInvitationStore: ReceivedInvitationStoreType
  ) {}

  setEmail(v: Email) {
    this.store.getState().setEmail(v)
  }

  clear() {
    this.store.getState().set(null)
  }

  setIsLoading(v: boolean) {
    this.store.getState().setIsLoading(v)
  }

  setReceivedInvitation(v: ReceivedInvitation) {
    this.receivedInvitationStore.getState().setReceivedInvitation(v)
  }

  setReceivedInvitationIsLoading(v: boolean) {
    this.receivedInvitationStore.getState().setReceivedInvitationIsLoading(v)
  }

  setIsInvitationProcessing(v: boolean) {
    this.receivedInvitationStore.getState().setIsInvitationProcessing(v)
  }
}
