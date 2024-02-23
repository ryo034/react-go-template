import { Invitations, Members } from "~/domain"

export interface WorkspaceUseCaseOutput {
  setMembers: (vs: Members) => void
  setMembersIsLoading: (v: boolean) => void
  clearMembers: () => void
  setInvitations: (vs: Invitations) => void
  setInvitationsIsLoading: (v: boolean) => void
}
