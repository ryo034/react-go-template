import { WorkspaceSubdomain } from "~/domain"
import { WorkspaceUseCase } from "~/usecase/workspace"

export interface CreateWorkspaceInput {
  subdomain: string
}

export class WorkspaceController {
  constructor(private readonly useCase: WorkspaceUseCase) {}

  async create(i: CreateWorkspaceInput): Promise<null | Error> {
    const subdomain = WorkspaceSubdomain.create(i.subdomain)
    if (subdomain.isErr) {
      return subdomain.error
    }
    return await this.useCase.create({ subdomain: subdomain.value })
  }
}
