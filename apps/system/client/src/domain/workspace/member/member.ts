import { Entity } from "@/domain/shared"
import { MemberId, MemberProfile, User } from "~/domain"

interface Props {
  id: MemberId
  user: User
  profile: MemberProfile
}

export class Member extends Entity<Props> {
  static create(v: Props): Member {
    return new Member(v)
  }

  get id(): MemberId {
    return this.value.id
  }

  get user(): User {
    return this.value.user
  }

  get profile(): MemberProfile {
    return this.value.profile
  }
}
