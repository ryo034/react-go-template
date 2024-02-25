import { AccountFullName, InvitationId, MemberDisplayName, MemberIdNumber } from "~/domain"
import { MemberBio } from "~/domain/workspace/member/bio"

export type UpdateProfileInput = {
  name: AccountFullName
}

export type AcceptInvitationInput = {
  invitationId: InvitationId
}

export type UpdateMemberProfileInput = {
  displayName?: MemberDisplayName
  idNumber?: MemberIdNumber
  bio: MemberBio
}
