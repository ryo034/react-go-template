import { Invitation, Invitations, Invitees, Members, Workspace, WorkspaceSubdomain } from "~/domain"
import { PromiseResult } from "~/infrastructure/shared"

export interface WorkspaceCreateInput {
  subdomain: WorkspaceSubdomain
}

export interface WorkspaceRepository {
  create(i: WorkspaceCreateInput): PromiseResult<Workspace, Error>
  findAllMembers(): PromiseResult<Members, Error>
  inviteMembers(invitees: Invitees): PromiseResult<null, Error>
  findAllInvitations(): PromiseResult<Invitations, Error>
  resendInvitation(invitation: Invitation): PromiseResult<null, Error>
  revokeInvitation(invitation: Invitation): PromiseResult<Invitations, Error>
}
