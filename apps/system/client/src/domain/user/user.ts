import { AccountFullName, AccountId } from "@/domain/account"
import { Email, Entity } from "@/domain/shared"

interface Props {
  id: AccountId
  name?: AccountFullName
  email: Email
}

export class User extends Entity<Props> {
  static create(v: Props): User {
    return new User(v)
  }

  get id(): AccountId {
    return this.value.id
  }

  get name(): AccountFullName | undefined {
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
