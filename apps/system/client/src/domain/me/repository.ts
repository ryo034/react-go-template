import { InvitationId, Me, User } from "~/domain"
import { PromiseResult } from "~/infrastructure/shared"

export interface MeRepository {
  signOut(): PromiseResult<null, Error>
  find(): PromiseResult<Me, Error>
  updateProfile(user: User): PromiseResult<Me, Error>
  acceptInvitation(invitationId: InvitationId): PromiseResult<Me, Error>
}
