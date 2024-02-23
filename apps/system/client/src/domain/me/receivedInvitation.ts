import { Invitation, Inviter } from "~/domain"
import { Entity } from "~/domain/shared"

interface Props {
  invitation: Invitation
  inviter: Inviter
}

export class ReceivedInvitation extends Entity<Props> {
  static create(v: Props): ReceivedInvitation {
    return new ReceivedInvitation(v)
  }

  get invitation(): Invitation {
    return this.value.invitation
  }

  get inviter(): Inviter {
    return this.value.inviter
  }
}
