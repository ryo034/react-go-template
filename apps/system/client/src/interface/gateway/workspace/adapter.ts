import { Result } from "true-myth"
import { Workspace, WorkspaceId, WorkspaceName, WorkspaceSubdomain, Workspaces } from "~/domain/workspace"
import { components } from "~/generated/schema/openapi/systemApi"
import { AdapterError } from "~/infrastructure/error"

export class WorkspaceGatewayAdapter {
  adapt(workspace: components["schemas"]["Workspace"]): Result<Workspace, Error> {
    if (workspace === undefined || workspace === null) {
      console.error(new AdapterError(WorkspaceGatewayAdapter.name, this.adapt.name, "workspace is required"))
      return Result.err(new Error("Workspace is not found"))
    }

    const id = WorkspaceId.fromString(workspace.workspaceId)
    if (id.isErr) {
      return Result.err(id.error)
    }

    const name = WorkspaceName.create(workspace.name)
    if (name.isErr) {
      return Result.err(name.error)
    }

    const subdomain = WorkspaceSubdomain.create(workspace.subdomain)
    if (subdomain.isErr) {
      return Result.err(subdomain.error)
    }

    return Result.ok(Workspace.create({ id: id.value, name: name.value, subdomain: subdomain.value }))
  }

  adaptAll(workspaces: components["schemas"]["Workspaces"]): Workspaces {
    const vs = workspaces.map((workspace) => {
      const res = this.adapt(workspace)
      if (res.isErr) {
        throw res.error
      }
      return res.value
    })
    return Workspaces.create(vs)
  }
}
