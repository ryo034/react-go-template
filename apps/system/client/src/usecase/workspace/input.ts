import { Invitees, WorkspaceId, WorkspaceSubdomain } from "~/domain"

export type CreateWorkspaceInput = {
  subdomain: WorkspaceSubdomain
}

export type InviteMembersInput = {
  invitees: Invitees
}
