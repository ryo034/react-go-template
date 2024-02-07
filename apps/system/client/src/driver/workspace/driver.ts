import { ApiErrorHandler } from "shared-network"
import { Result } from "true-myth"
import { WorkspaceCreateInput } from "~/domain"
import { components } from "~/generated/schema/openapi/systemApi"
import { SystemAPIClient } from "~/infrastructure/openapi/client"
import { PromiseResult } from "~/infrastructure/shared/result"

export class WorkspaceDriver {
  constructor(private readonly client: SystemAPIClient, private readonly errorHandler: ApiErrorHandler) {}

  async create(i: WorkspaceCreateInput): PromiseResult<components["schemas"]["Workspace"], Error> {
    try {
      const res = await this.client.POST("/api/v1/workspaces", {
        body: {
          name: "",
          subdomain: i.subdomain.value
        }
      })
      return res.data ? Result.ok(res.data) : Result.err(this.errorHandler.adapt(res))
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }
}
