import { Email, Entity, Member, MemberDisplayName, Workspace } from "~/domain"

interface Props {
  self: Member
  workspace: Workspace
}

export class Inviter extends Entity<Props> {
  static create(v: Props): Inviter {
    return new Inviter(v)
  }

  get self(): Member {
    return this.value.self
  }

  get workspace(): Workspace {
    return this.value.workspace
  }
}
