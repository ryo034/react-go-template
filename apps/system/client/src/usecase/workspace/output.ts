import type { Invitations, Member, Members } from "~/domain"

export interface WorkspaceUseCaseOutput {
  setMembers: (vs: Members) => void
  updateMember(m: Member): void
  setMembersIsLoading: (v: boolean) => void
  clearMembers: () => void
  setInvitations: (vs: Invitations) => void
  addInvitations: (vs: Invitations) => void
  setInvitationsIsLoading: (v: boolean) => void
}
