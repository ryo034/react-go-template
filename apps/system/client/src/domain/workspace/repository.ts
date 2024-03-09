import {
  Invitation,
  Invitations,
  Invitees,
  Member,
  MemberId,
  Members,
  SelectableRole,
  Workspace,
  WorkspaceId,
  WorkspaceName,
  WorkspaceSubdomain
} from "~/domain"
import { PromiseResult } from "~/infrastructure/shared"

export interface WorkspaceCreateInput {
  subdomain: WorkspaceSubdomain
}

export interface WorkspaceRepository {
  create(i: WorkspaceCreateInput): PromiseResult<Workspace, Error>
  findAllMembers(): PromiseResult<Members, Error>
  inviteMembers(invitees: Invitees): PromiseResult<Invitations, Error>
  findAllInvitations(): PromiseResult<Invitations, Error>
  resendInvitation(invitation: Invitation): PromiseResult<null, Error>
  revokeInvitation(invitation: Invitation): PromiseResult<Invitations, Error>
  updateMemberRole(memberId: MemberId, role: SelectableRole): PromiseResult<Member, Error>
  updateWorkspace(
    workspaceId: WorkspaceId,
    name: WorkspaceName,
    subdomain: WorkspaceSubdomain
  ): PromiseResult<Workspace, Error>
}
