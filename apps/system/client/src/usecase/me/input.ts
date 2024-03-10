import type { AccountFullName, InvitationId, MemberDisplayName, MemberIdNumber } from "~/domain"
import type { MemberBio } from "~/domain/workspace/member/bio"

export type UpdateProfileInput = {
  name: AccountFullName
}

export type UpdatePhotoInput = {
  file: File
}

export type AcceptInvitationInput = {
  invitationId: InvitationId
}

export type UpdateMemberProfileInput = {
  displayName?: MemberDisplayName
  idNumber?: MemberIdNumber
  bio: MemberBio
}
