import { Entity } from "@/domain/shared"
import { MemberId, MemberProfile, MemberRole, User } from "~/domain"

interface Props {
  id: MemberId
  user: User
  profile: MemberProfile
  role: MemberRole
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

  get role(): MemberRole {
    return this.value.role
  }

  get isOwner(): boolean {
    return this.value.role === "owner"
  }
}
