import type { WorkspaceId, WorkspaceName, WorkspaceSubdomain } from "~/domain"
import { Entity } from "~/domain/shared"

interface Props {
  id: WorkspaceId
  name: WorkspaceName
  subdomain: WorkspaceSubdomain
}

export class Workspace extends Entity<Props> {
  static create(v: Props): Workspace {
    return new Workspace(v)
  }

  get id(): WorkspaceId {
    return this.value.id
  }

  get name(): WorkspaceName {
    return this.value.name
  }

  get subdomain(): WorkspaceSubdomain {
    return this.value.subdomain
  }
}
