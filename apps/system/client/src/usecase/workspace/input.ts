import type { Invitees, WorkspaceId, WorkspaceName, WorkspaceSubdomain } from "~/domain"

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
