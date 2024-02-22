import { WorkspaceRepository } from "~/domain"
import { CreateWorkspaceInput, InviteMembersInput, MeUseCase, WorkspaceUseCaseOutput } from "~/usecase"

export interface WorkspaceUseCase {
  create(i: CreateWorkspaceInput): Promise<Error | null>
  findAllMembers(): Promise<Error | null>
  inviteMembers(i: InviteMembersInput): Promise<Error | null>
  findAllInvitations(): Promise<Error | null>
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

  async inviteMembers(i: InviteMembersInput): Promise<Error | null> {
    const res = await this.repository.inviteMembers(i.invitees)
    if (res.isErr) {
      return res.error
    }
    return null
  }

  async findAllInvitations(): Promise<Error | null> {
    this.presenter.setInvitationsIsLoading(true)
    const res = await this.repository.findAllInvitations()
    if (res.isErr) {
      return res.error
    }
    this.presenter.setInvitations(res.value)
    this.presenter.setInvitationsIsLoading(false)
    return null
  }
}
