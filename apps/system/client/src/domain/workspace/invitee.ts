import type { MemberDisplayName } from "~/domain"
import { type Email, Entity } from "~/domain/shared"

interface Props {
  email: Email
  name: MemberDisplayName
}

export class Invitee extends Entity<Props> {
  static create(v: Props): Invitee {
    return new Invitee(v)
  }

  get email(): Email {
    return this.value.email
  }

  get name(): MemberDisplayName {
    return this.value.name
  }
}
