import { AppDateTime, Email, Entity } from "@/domain/shared"
import { MemberDisplayName } from "../member"
import { InvitationId } from "./id"

interface Props {
  id: InvitationId
  verified: boolean
  expiredAt: AppDateTime
  inviteeEmail: Email
  displayName: MemberDisplayName
}

export class Invitation extends Entity<Props> {
  static create(v: Props): Invitation {
    return new Invitation(v)
  }

  get id(): InvitationId {
    return this.value.id
  }

  get verified(): boolean {
    return this.value.verified
  }

  get expiredAt(): AppDateTime {
    return this.value.expiredAt
  }

  get inviteeEmail(): Email {
    return this.value.inviteeEmail
  }

  get displayName(): MemberDisplayName {
    return this.value.displayName
  }
}
