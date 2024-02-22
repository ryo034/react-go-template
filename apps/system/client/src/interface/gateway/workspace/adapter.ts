import { Result } from "true-myth"
import { Inviter, Workspace, WorkspaceId, WorkspaceName, WorkspaceSubdomain, Workspaces } from "~/domain"
import { components } from "~/generated/schema/openapi/systemApi"
import { AdapterError } from "~/infrastructure/error"
import { MemberGatewayAdapter } from "./member"

export class WorkspaceGatewayAdapter {
  constructor(private readonly memberAdapter: MemberGatewayAdapter) {}
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

  adaptAll(workspaces: components["schemas"]["Workspaces"]): Result<Workspaces, Error> {
    const vs: Workspace[] = []
    for (const w of workspaces) {
      const res = this.adapt(w)
      if (res.isErr) {
        return Result.err(res.error)
      }
      vs.push(res.value)
    }
    return Result.ok(Workspaces.create(vs))
  }

  adaptInviter(inviter: components["schemas"]["Inviter"]): Result<Inviter, Error> {
    if (inviter === undefined || inviter === null) {
      console.error(new AdapterError(WorkspaceGatewayAdapter.name, this.adaptInviter.name, "inviter is required"))
      return Result.err(new Error("Inviter is not found"))
    }
    const self = this.memberAdapter.adapt(inviter.member)
    if (self.isErr) {
      return Result.err(self.error)
    }
    const workspace = this.adapt(inviter.workspace)
    if (workspace.isErr) {
      return Result.err(workspace.error)
    }
    return Result.ok(Inviter.create({ self: self.value, workspace: workspace.value }))
  }
}
