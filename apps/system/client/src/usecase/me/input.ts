import { InvitationId, User } from "~/domain"

export type UpdateProfileInput = {
  user: User
}

export type AcceptInvitationInput = {
  invitationId: InvitationId
}
