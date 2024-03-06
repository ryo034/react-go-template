import { InvitationId, Member, MemberDisplayName } from "~/domain"
import { AppDateTime, Email, Entity } from "~/domain/shared"

interface Props {
  id: InvitationId
  accepted: boolean
  expiredAt: AppDateTime
  inviteeEmail: Email
  displayName: MemberDisplayName | null
  inviter: Member
}

export class Invitation extends Entity<Props> {
  static create(v: Props): Invitation {
    return new Invitation(v)
  }

  get id(): InvitationId {
    return this.value.id
  }

  get accepted(): boolean {
    return this.value.accepted
  }

  get expiredAt(): AppDateTime {
    return this.value.expiredAt
  }

  get inviteeEmail(): Email {
    return this.value.inviteeEmail
  }

  get displayName(): MemberDisplayName | null {
    return this.value.displayName
  }

  get inviter(): Member {
    return this.value.inviter
  }
}
