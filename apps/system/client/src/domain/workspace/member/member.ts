import { Entity } from "@/domain/shared"
import { User } from "~/domain/user"
import { MemberProfile } from "~/domain/workspace/member"

interface Props {
  user: User
  profile: MemberProfile
}

export class Member extends Entity<Props> {
  static create(v: Props): Member {
    return new Member(v)
  }

  get user(): User {
    return this.value.user
  }

  get profile(): MemberProfile {
    return this.value.profile
  }
}
