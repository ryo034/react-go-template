import { AccountId, AccountName, Entity, WorkspaceId, WorkspaceName, WorkspaceSubdomain } from "~/domain"

interface Props {
  id: WorkspaceId
  name: WorkspaceName
  subdomain: WorkspaceSubdomain
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

  get subdomain(): WorkspaceSubdomain {
    return this.value.subdomain
  }
}
