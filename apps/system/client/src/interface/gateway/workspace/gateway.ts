import { Result } from "true-myth"
import { Members, Workspace, WorkspaceCreateInput, WorkspaceRepository } from "~/domain"
import { WorkspaceDriver } from "~/driver/workspace/driver"
import { PromiseResult } from "~/infrastructure/shared/result"
import { WorkspaceGatewayAdapter } from "./adapter"
import { MemberGatewayAdapter } from "./member"

export class WorkspaceGateway implements WorkspaceRepository {
  constructor(
    private readonly driver: WorkspaceDriver,
    private readonly adapter: WorkspaceGatewayAdapter,
    private readonly memberAdapter: MemberGatewayAdapter
  ) {}

  async create(i: WorkspaceCreateInput): PromiseResult<Workspace, Error> {
    const res = await this.driver.create(i)
    if (res.isErr) {
      return Result.err(res.error)
    }
    return this.adapter.adapt(res.value)
  }

  async findAllMembers(): PromiseResult<Members, Error> {
    const res = await this.driver.findAllMembers()
    if (res.isErr) {
      return Result.err(res.error)
    }
    return this.memberAdapter.adaptAll(res.value)
  }
}
