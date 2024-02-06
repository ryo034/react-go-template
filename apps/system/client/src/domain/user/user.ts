import { AccountId, AccountName } from "@/domain/account"
import { Email, Entity } from "@/domain/shared"

interface Props {
  id: AccountId
  name?: AccountName
  email: Email
}

export class User extends Entity<Props> {
  static create(v: Props): User {
    return new User(v)
  }

  get id(): AccountId {
    return this.value.id
  }

  get name(): AccountName | undefined {
    return this.value.name
  }

  get hasName(): boolean {
    return !!this.value.name
  }

  get hasNotName(): boolean {
    return !this.hasName
  }

  get email(): Email {
    return this.value.email
  }
}
