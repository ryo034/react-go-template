import { InvitationId, Me, MemberProfile, User } from "~/domain"
import { PromiseResult } from "~/infrastructure/shared"

export interface MeRepository {
  signOut(): PromiseResult<null, Error>
  find(): PromiseResult<Me, Error>
  acceptInvitation(invitationId: InvitationId): PromiseResult<Me, Error>
  updateProfile(user: User): PromiseResult<Me, Error>
  updateMemberProfile(profile: MemberProfile): PromiseResult<Me, Error>
}
