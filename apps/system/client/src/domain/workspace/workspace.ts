import { AccountId, AccountName } from "@/domain/account"
import { Entity } from "@/domain/shared"
import { WorkspaceId } from "./id"
import { WorkspaceName } from "./name"

interface Props {
  id: WorkspaceId
  name: WorkspaceName
}

export class Workspace extends Entity<Props> {
  static create(v: Props): Workspace {
    return new Workspace(v)
  }

  get id(): AccountId {
    return this.value.id
  }

  get name(): AccountName | null {
    return this.value.name
  }
}
