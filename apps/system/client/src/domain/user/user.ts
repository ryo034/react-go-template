import type { AccountFullName, AccountId } from "~/domain/account"
import { type Email, Entity, type Photo } from "~/domain/shared"

interface Props {
  id: AccountId
  email: Email
  name?: AccountFullName
  photo?: Photo
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
    return this.value.name !== undefined && this.value.name.value !== ""
  }

  get hasNotName(): boolean {
    return !this.hasName
  }

  get email(): Email {
    return this.value.email
  }

  get photo(): Photo | undefined {
    return this.value.photo
  }

  get hasPhoto(): boolean {
    return this.value.photo !== undefined
  }
}
