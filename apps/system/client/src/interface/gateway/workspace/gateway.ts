import { Result } from "true-myth"
import {
  Invitation,
  Invitations,
  Invitees,
  Members,
  Workspace,
  WorkspaceCreateInput,
  WorkspaceRepository
} from "~/domain"
import { WorkspaceDriver } from "~/driver/workspace/driver"
import { PromiseResult } from "~/infrastructure/shared/result"
import { WorkspaceGatewayAdapter } from "./adapter"
import { InvitationGatewayAdapter } from "./invitation"
import { MemberGatewayAdapter } from "./member"

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

  async inviteMembers(invitees: Invitees): PromiseResult<null, Error> {
    const res = await this.driver.inviteMembers(invitees)
    if (res.isErr) {
      return Result.err(res.error)
    }
    return Result.ok(null)
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
}
