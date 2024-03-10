import { Result } from "true-myth"
import { Me, ReceivedInvitation, ReceivedInvitations } from "~/domain/me"
import type { Workspace } from "~/domain/workspace"
import type { Member } from "~/domain/workspace/member"
import type { components } from "~/generated/schema/openapi/systemApi"
import { AdapterError, AuthProviderCurrentUserNotFoundError } from "~/infrastructure/error"
import type { UserGatewayAdapter } from "~/interface/gateway/user"
import type { WorkspaceGatewayAdapter } from "~/interface/gateway/workspace"
import type { MemberGatewayAdapter } from "~/interface/gateway/workspace/member"
import type { AuthGatewayAdapter } from "../auth"
import type { InvitationGatewayAdapter } from "../workspace/invitation"

export class MeGatewayAdapter {
  constructor(
    private readonly authAdapter: AuthGatewayAdapter,
    private readonly userAdapter: UserGatewayAdapter,
    private readonly memberAdapter: MemberGatewayAdapter,
    private readonly workspaceAdapter: WorkspaceGatewayAdapter,
    private readonly invitationAdapter: InvitationGatewayAdapter
  ) {}

  adapt(me: components["schemas"]["Me"]): Result<Me, Error> {
    if (me === undefined || me === null) {
      console.error(new AdapterError(MeGatewayAdapter.name, this.adapt.name, "me is required"))
      return Result.err(new AuthProviderCurrentUserNotFoundError("User is not found"))
    }

    let member: Member | undefined = undefined
    if (me.member !== undefined && me.member !== null) {
      const tmpMember = this.memberAdapter.adapt(me.member)
      if (tmpMember.isErr) {
        return Result.err(tmpMember.error)
      }
      member = tmpMember.value
    }

    let workspace: Workspace | undefined = undefined
    if (me.currentWorkspace !== undefined && me.currentWorkspace !== null) {
      const tmpWorkspace = this.workspaceAdapter.adapt(me.currentWorkspace)
      if (tmpWorkspace.isErr) {
        return Result.err(tmpWorkspace.error)
      }
      workspace = tmpWorkspace.value
    }

    const joinedWorkspaces = this.workspaceAdapter.adaptAll(me.joinedWorkspaces)
    if (joinedWorkspaces.isErr) {
      return Result.err(joinedWorkspaces.error)
    }

    const user = this.userAdapter.adapt(me.self)
    if (user.isErr) {
      return Result.err(user.error)
    }

    const rivs = []
    if (me.receivedInvitations !== undefined) {
      for (const ari of me.receivedInvitations) {
        const ri = this.adaptReceivedInvitation(ari)
        if (ri.isErr) {
          return Result.err(ri.error)
        }
        rivs.push(ri.value)
      }
    }

    const providers = this.authAdapter.adaptAllAuthProvider(me.providers)
    if (providers.isErr) {
      return Result.err(providers.error)
    }

    return Result.ok(
      Me.create({
        self: user.value,
        workspace,
        member,
        joinedWorkspaces: joinedWorkspaces.value,
        receivedInvitations: ReceivedInvitations.create(rivs),
        providers: providers.value
      })
    )
  }

  adaptReceivedInvitation(
    receivedInvitation: components["schemas"]["ReceivedInvitation"]
  ): Result<ReceivedInvitation, Error> {
    const invitation = this.invitationAdapter.adapt(receivedInvitation.invitation)
    if (invitation.isErr) {
      return Result.err(invitation.error)
    }
    const inviter = this.workspaceAdapter.adaptInviter(receivedInvitation.inviter)
    if (inviter.isErr) {
      return Result.err(inviter.error)
    }
    return Result.ok(
      ReceivedInvitation.create({
        invitation: invitation.value,
        inviter: inviter.value
      })
    )
  }
}
