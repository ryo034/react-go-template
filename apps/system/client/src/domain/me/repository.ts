import { AccountFullName, InvitationId, Me, MemberProfile } from "~/domain"
import { PromiseResult } from "~/infrastructure/shared"

export interface MeRepository {
  signOut(): PromiseResult<null, Error>
  find(): PromiseResult<Me, Error>
  acceptInvitation(invitationId: InvitationId): PromiseResult<Me, Error>
  updateProfile(name: AccountFullName): PromiseResult<Me, Error>
  updatePhoto(file: File): PromiseResult<Me, Error>
  removePhoto(): PromiseResult<Me, Error>
  updateMemberProfile(profile: MemberProfile): PromiseResult<Me, Error>
}
