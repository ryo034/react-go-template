import { WorkspaceRepository } from "~/domain"
import { CreateWorkspaceInput, MeUseCase, WorkspaceUseCaseOutput } from "~/usecase"

export interface WorkspaceUseCase {
  create(i: CreateWorkspaceInput): Promise<Error | null>
  findAllMembers(): Promise<Error | null>
}

export class WorkspaceInteractor implements WorkspaceUseCase {
  constructor(
    private readonly repository: WorkspaceRepository,
    private readonly meUseCase: MeUseCase,
    private readonly presenter: WorkspaceUseCaseOutput
  ) {}

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

  async findAllMembers(): Promise<Error | null> {
    this.presenter.setMembersIsLoading(true)
    const res = await this.repository.findAllMembers()
    if (res.isErr) {
      this.presenter.setMembersIsLoading(false)
      return res.error
    }
    this.presenter.setMembers(res.value)
    this.presenter.setMembersIsLoading(false)
    return null
  }
}