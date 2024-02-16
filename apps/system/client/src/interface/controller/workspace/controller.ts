import { Email, Invitee, Invitees, MemberDisplayName, WorkspaceSubdomain } from "~/domain"
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
}
