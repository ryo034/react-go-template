import { Result } from "true-myth"
import { Workspace, WorkspaceCreateInput, WorkspaceRepository } from "~/domain"
import { WorkspaceDriver } from "~/driver/workspace/driver"
import { PromiseResult } from "~/infrastructure/shared/result"
import { WorkspaceGatewayAdapter } from "./adapter"

export class WorkspaceGateway implements WorkspaceRepository {
  constructor(private readonly driver: WorkspaceDriver, private readonly adapter: WorkspaceGatewayAdapter) {}

  async create(i: WorkspaceCreateInput): PromiseResult<Workspace, Error> {
    const res = await this.driver.create(i)
    if (res.isErr) {
      return Result.err(res.error)
    }
    return this.adapter.adapt(res.value)
  }
}
