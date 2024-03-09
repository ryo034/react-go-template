import {
  Invitation,
  Invitee,
  Invitees,
  MemberDisplayName,
  MemberId,
  SelectableRole,
  WorkspaceId,
  WorkspaceName,
  WorkspaceSubdomain
} from "~/domain"
import { Email } from "~/domain/shared"
import { WorkspaceUseCase } from "~/usecase/workspace"

export interface CreateWorkspaceInput {
  subdomain: string
}

interface InviteMembersInput {
  invitees: {
    email: string
    name: string
  }[]
}

interface UpdateWorkspaceInput {
  workspaceId: WorkspaceId
  name: string
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

  async findAllMembers(): Promise<null | Error> {
    return await this.useCase.findAllMembers()
  }

  async inviteMembers(i: InviteMembersInput): Promise<null | Error> {
    let invitees = Invitees.empty()
    for (const inv of i.invitees) {
      const e = Email.create(inv.email)
      if (e.isErr) {
        return e.error
      }
      const md = MemberDisplayName.create(inv.name)
      if (md.isErr) {
        return md.error
      }
      invitees = invitees.add(Invitee.create({ email: e.value, name: md.value }))
    }
    return await this.useCase.inviteMembers({ invitees })
  }

  async findAllInvitations(): Promise<null | Error> {
    return await this.useCase.findAllInvitations()
  }

  async resendInvitation(i: Invitation): Promise<null | Error> {
    return await this.useCase.resendInvitation(i)
  }

  async revokeInvitation(i: Invitation): Promise<null | Error> {
    return await this.useCase.revokeInvitation(i)
  }

  async updateMemberRole(memberId: MemberId, role: SelectableRole): Promise<null | Error> {
    return await this.useCase.updateMemberRole(memberId, role)
  }

  async updateWorkspace(i: UpdateWorkspaceInput): Promise<null | Error> {
    const n = WorkspaceName.create(i.name)
    if (n.isErr) {
      return n.error
    }
    const s = WorkspaceSubdomain.create(i.subdomain)
    if (s.isErr) {
      return s.error
    }
    return await this.useCase.updateWorkspace({ workspaceId: i.workspaceId, name: n.value, subdomain: s.value })
  }
}
