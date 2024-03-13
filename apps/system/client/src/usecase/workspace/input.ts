import type { Invitees, MemberId, WorkspaceId, WorkspaceName, WorkspaceSubdomain } from "~/domain"

export type CreateWorkspaceInput = {
  subdomain: WorkspaceSubdomain
}

export type InviteMembersInput = {
  invitees: Invitees
}

export type UpdateWorkspaceInput = {
  workspaceId: WorkspaceId
  name: WorkspaceName
  subdomain: WorkspaceSubdomain
}

export type LeaveWorkspaceInput = {
  memberId: MemberId
}
