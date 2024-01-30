import { Entity } from "~/domain/shared"
import { User } from "~/domain/user"
import { Workspace } from "~/domain/workspace"
import { Member } from "~/domain/workspace/member"

interface Props {
  self: User
  workspace?: Workspace
  member?: Member
}

export class Me extends Entity<Props> {
  static create(v: Props): Me {
    return new Me(v)
  }

  get self(): User {
    return this.value.self
  }

  get workspace(): Workspace | undefined {
    return this.value.workspace
  }

  get member(): Member | undefined {
    return this.value.member
  }

  get hasWorkspace(): boolean {
    return this.value.workspace !== undefined
  }

  get hasMember(): boolean {
    return this.value.member !== undefined
  }
}
