import { Result } from "true-myth"
import type {
  Invitation,
  Invitations,
  Invitees,
  Member,
  MemberId,
  Members,
  SelectableRole,
  Workspace,
  WorkspaceCreateInput,
  WorkspaceId,
  WorkspaceName,
  WorkspaceRepository,
  WorkspaceSubdomain
} from "~/domain"
import type { WorkspaceDriver } from "~/driver/workspace/driver"
import type { PromiseResult } from "~/infrastructure/shared/result"
import type { WorkspaceGatewayAdapter } from "./adapter"
import type { InvitationGatewayAdapter } from "./invitation"
import type { MemberGatewayAdapter } from "./member"

export class WorkspaceGateway implements WorkspaceRepository {
  constructor(
    private readonly driver: WorkspaceDriver,
    private readonly adapter: WorkspaceGatewayAdapter,
    private readonly memberAdapter: MemberGatewayAdapter,
    private readonly invitationAdapter: InvitationGatewayAdapter
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

  async inviteMembers(invitees: Invitees): PromiseResult<Invitations, Error> {
    const res = await this.driver.inviteMembers(invitees)
    if (res.isErr) {
      return Result.err(res.error)
    }
    return this.invitationAdapter.adaptAll(res.value)
  }

  async findAllInvitations(): PromiseResult<Invitations, Error> {
    const res = await this.driver.findAllInvitations()
    if (res.isErr) {
      return Result.err(res.error)
    }
    return this.invitationAdapter.adaptAll(res.value)
  }

  async resendInvitation(invitation: Invitation): PromiseResult<null, Error> {
    const res = await this.driver.resendInvitation(invitation)
    if (res.isErr) {
      return Result.err(res.error)
    }
    return Result.ok(null)
  }

  async revokeInvitation(invitation: Invitation): PromiseResult<Invitations, Error> {
    const res = await this.driver.revokeInvitation(invitation)
    if (res.isErr) {
      return Result.err(res.error)
    }
    return this.invitationAdapter.adaptAll(res.value)
  }

  async updateMemberRole(memberId: MemberId, role: SelectableRole): PromiseResult<Member, Error> {
    const res = await this.driver.updateMemberRole(memberId, role)
    if (res.isErr) {
      return Result.err(res.error)
    }
    return this.memberAdapter.adapt(res.value)
  }

  async updateWorkspace(
    workspaceId: WorkspaceId,
    name: WorkspaceName,
    subdomain: WorkspaceSubdomain
  ): PromiseResult<Workspace, Error> {
    const res = await this.driver.updateWorkspace(workspaceId, name, subdomain)
    if (res.isErr) {
      return Result.err(res.error)
    }
    return this.adapter.adapt(res.value)
  }

  async leaveWorkspace(memberId: MemberId): PromiseResult<null, Error> {
    const res = await this.driver.leaveWorkspace(memberId)
    if (res.isErr) {
      return Result.err(res.error)
    }
    return Result.ok(null)
  }
}
