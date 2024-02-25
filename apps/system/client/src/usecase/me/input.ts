import { InvitationId, MemberDisplayName, MemberIdNumber, User } from "~/domain"
import { MemberBio } from "~/domain/workspace/member/bio"

export type UpdateProfileInput = {
  user: User
}

export type AcceptInvitationInput = {
  invitationId: InvitationId
}

export type UpdateMemberProfileInput = {
  displayName?: MemberDisplayName
  idNumber?: MemberIdNumber
  bio: MemberBio
}
