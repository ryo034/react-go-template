import { ApiErrorHandler } from "shared-network"
import { Result } from "true-myth"
import { Invitees, WorkspaceCreateInput, WorkspaceId } from "~/domain"
import { components } from "~/generated/schema/openapi/systemApi"
import { SystemAPIClient } from "~/infrastructure/openapi/client"
import { PromiseResult } from "~/infrastructure/shared/result"

export class WorkspaceDriver {
  constructor(private readonly client: SystemAPIClient, private readonly errorHandler: ApiErrorHandler) {}

  async create(i: WorkspaceCreateInput): PromiseResult<components["schemas"]["Workspace"], Error> {
    try {
      const res = await this.client.POST("/api/v1/workspaces", {
        body: {
          subdomain: i.subdomain.value
        }
      })
      return res.data ? Result.ok(res.data) : Result.err(this.errorHandler.adapt(res))
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }

  async findAllMembers(): PromiseResult<components["schemas"]["Members"], Error> {
    try {
      const res = await this.client.GET("/api/v1/members")
      return res.data ? Result.ok(res.data) : Result.err(this.errorHandler.adapt(res))
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }

  async inviteMembers(invitees: Invitees): PromiseResult<null, Error> {
    try {
      const res = await this.client.POST("/api/v1/members/invitations/bulk", {
        body: {
          invitees: invitees.values.map((i) => ({
            email: i.email.value,
            name: i.name.value
          }))
        }
      })
      return res.data ? Result.ok(null) : Result.err(this.errorHandler.adapt(res))
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }
}
