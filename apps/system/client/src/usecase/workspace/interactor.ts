import { User, WorkspaceRepository } from "~/domain"
import { MeUseCase } from "~/usecase"
import { CreateWorkspaceInput } from "./input"

export interface WorkspaceUseCase {
  create(i: CreateWorkspaceInput): Promise<Error | null>
}

export type UpdateProfileInput = {
  user: User
}

export class WorkspaceInteractor implements WorkspaceUseCase {
  constructor(private readonly repository: WorkspaceRepository, private readonly meUseCase: MeUseCase) {}

  async create(i: CreateWorkspaceInput): Promise<Error | null> {
    const res = await this.repository.create(i)
    if (res.isErr) {
      return res.error
    }
    const meRes = await this.meUseCase.find()
    if (meRes) {
      return meRes
    }
    return null
  }
}
