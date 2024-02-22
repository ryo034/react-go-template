import { Invitations, Members } from "~/domain"
import { InvitationStoreType } from "~/store/invitations/store"
import { WorkspaceStoreType } from "~/store/workspace/store"
import { WorkspaceUseCaseOutput } from "~/usecase"

export class WorkspacePresenter implements WorkspaceUseCaseOutput {
  constructor(private readonly store: WorkspaceStoreType, private readonly invitationsStore: InvitationStoreType) {}

  setMembers(vs: Members) {
    this.store.getState().setMembers(vs)
  }

  setMembersIsLoading(v: boolean) {
    this.store.getState().setMembersIsLoading(v)
  }

  clearMembers() {
    this.store.getState().setMembers(Members.empty())
  }

  setInvitations(vs: Invitations) {
    this.invitationsStore.getState().setInvitations(vs)
  }

  setInvitationsIsLoading(v: boolean) {
    this.invitationsStore.getState().setInvitationsIsLoading(v)
  }
}
