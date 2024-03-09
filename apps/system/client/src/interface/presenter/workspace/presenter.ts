import { Invitations, Member, Members } from "~/domain"
import { InvitationsStoreType } from "~/store/invitations/store"
import { WorkspaceStoreType } from "~/store/workspace/store"
import { WorkspaceUseCaseOutput } from "~/usecase"

export class WorkspacePresenter implements WorkspaceUseCaseOutput {
  constructor(private readonly store: WorkspaceStoreType, private readonly invitationsStore: InvitationsStoreType) {}

  setMembers(vs: Members) {
    this.store.getState().setMembers(vs)
  }

  updateMember(m: Member) {
    const tmpMembers = [...this.store.getState().members.values]
    const targetIndex = tmpMembers.findIndex((v) => v.id.value.asString === m.id.value.asString)
    if (targetIndex === -1) {
      return
    }
    tmpMembers[targetIndex] = m
    this.store.getState().setMembers(Members.create(tmpMembers))
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

  addInvitations(vs: Invitations) {
    const tmpInvitations = [...this.invitationsStore.getState().invitations.values]
    for (const v of vs.values) {
      tmpInvitations.unshift(v)
    }
    this.invitationsStore.getState().setInvitations(Invitations.create(tmpInvitations))
  }

  setInvitationsIsLoading(v: boolean) {
    this.invitationsStore.getState().setInvitationsIsLoading(v)
  }
}
