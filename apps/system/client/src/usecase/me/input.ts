import { InvitationId, User } from "~/domain"

export type UpdateProfileInput = {
  user: User
}

export type AcceptInvitationInput = {
  invitationId: InvitationId
}

export type UpdateMemberProfileInput = {
  displayName: string
  idNumber: string
  bio: string
}
