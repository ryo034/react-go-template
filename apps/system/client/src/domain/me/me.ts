import { Entity } from "~/domain/shared"
import { User } from "~/domain/user"
import { Workspace, Workspaces } from "~/domain/workspace"
import { Member } from "~/domain/workspace/member"

interface Props {
  self: User
  workspace?: Workspace
  member?: Member
  joinedWorkspaces: Workspaces
}

export class Me extends Entity<Props> {
  static create(v: Props): Me {
    return new Me(v)
  }

  get self(): User {
    return this.value.self
  }

  get doneOnboarding(): boolean {
    return this.value.self.hasName && this.hasWorkspace
  }

  get workspace(): Workspace | undefined {
    return this.value.workspace
  }

  get member(): Member | undefined {
    return this.value.member
  }

  get hasWorkspace(): boolean {
    return !!this.value.workspace
  }

  get hasNotWorkspace(): boolean {
    return !this.hasWorkspace
  }

  get hasMember(): boolean {
    return this.value.member !== undefined
  }

  get joinedWorkspaces(): Workspaces {
    return this.value.joinedWorkspaces
  }
}
