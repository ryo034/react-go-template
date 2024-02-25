import { Entity } from "@/domain/shared"
import { MemberDisplayName, MemberId, MemberIdNumber } from "~/domain/workspace/member"
import { MemberBio } from "./bio"

interface Props {
  id: MemberId
  displayName?: MemberDisplayName
  idNumber?: MemberIdNumber
  bio: MemberBio
}

export class MemberProfile extends Entity<Props> {
  static create(v: Props): MemberProfile {
    return new MemberProfile(v)
  }

  get id(): MemberId {
    return this.value.id
  }

  get displayName(): MemberDisplayName | undefined {
    return this.value.displayName
  }

  get idNumber(): MemberIdNumber | undefined {
    return this.value.idNumber
  }

  get bio(): MemberBio {
    return this.value.bio
  }

  get hasDisplayName(): boolean {
    return this.displayName !== undefined
  }

  get hasIdNumber(): boolean {
    return this.idNumber !== undefined
  }
}
