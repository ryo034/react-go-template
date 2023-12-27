import { Entity } from "~/domain/shared"
import { User } from "~/domain/user"

interface Props {
  user: User
  emailVerified: boolean
}

export class Me extends Entity<Props> {
  static create(v: Props): Me {
    return new Me(v)
  }

  get user(): User {
    return this.value.user
  }

  get emailVerified(): boolean {
    return this.value.emailVerified
  }

  get emailNotVerified(): boolean {
    return !this.emailVerified
  }
}
