import type { Invitation, MemberId, SelectableRole, WorkspaceRepository } from "~/domain"
import type {
  CreateWorkspaceInput,
  InviteMembersInput,
  MeUseCase,
  UpdateWorkspaceInput,
  WorkspaceUseCaseOutput
} from "~/usecase"

export interface WorkspaceUseCase {
  create(i: CreateWorkspaceInput): Promise<Error | null>
  findAllMembers(): Promise<Error | null>
  inviteMembers(i: InviteMembersInput): Promise<Error | null>
  findAllInvitations(): Promise<Error | null>
  resendInvitation(i: Invitation): Promise<Error | null>
  revokeInvitation(i: Invitation): Promise<Error | null>
  updateMemberRole(memberId: MemberId, role: SelectableRole): Promise<Error | null>
  updateWorkspace(i: UpdateWorkspaceInput): Promise<Error | null>
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
    const meErr = await this.meUseCase.find()
    if (meErr) {
      return meErr
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
    this.presenter.addInvitations(res.value)
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

  async resendInvitation(i: Invitation): Promise<Error | null> {
    const res = await this.repository.resendInvitation(i)
    if (res.isErr) {
      return res.error
    }
    return null
  }

  async revokeInvitation(i: Invitation): Promise<Error | null> {
    const res = await this.repository.revokeInvitation(i)
    if (res.isErr) {
      return res.error
    }
    this.presenter.setInvitations(res.value)
    return null
  }

  async updateMemberRole(memberId: MemberId, role: SelectableRole): Promise<Error | null> {
    const res = await this.repository.updateMemberRole(memberId, role)
    if (res.isErr) {
      return res.error
    }
    this.presenter.updateMember(res.value)
    return null
  }

  async updateWorkspace(i: UpdateWorkspaceInput): Promise<Error | null> {
    const res = await this.repository.updateWorkspace(i.workspaceId, i.name, i.subdomain)
    if (res.isErr) {
      return res.error
    }
    const meErr = await this.meUseCase.refetch()
    if (meErr) {
      return meErr
    }
    return null
  }
}
