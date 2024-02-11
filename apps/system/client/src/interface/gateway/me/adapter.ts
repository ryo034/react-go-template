import { Result } from "true-myth"
import { Me } from "~/domain/me"
import { Workspace } from "~/domain/workspace"
import { Member } from "~/domain/workspace/member"
import { components } from "~/generated/schema/openapi/systemApi"
import { AdapterError, AuthProviderCurrentUserNotFoundError } from "~/infrastructure/error"
import { UserGatewayAdapter } from "~/interface/gateway/user"
import { WorkspaceGatewayAdapter } from "~/interface/gateway/workspace"
import { MemberGatewayAdapter } from "~/interface/gateway/workspace/member"

export class MeGatewayAdapter {
  constructor(
    private readonly userAdapter: UserGatewayAdapter,
    private readonly memberAdapter: MemberGatewayAdapter,
    private readonly workspaceAdapter: WorkspaceGatewayAdapter
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

    return Result.ok(Me.create({ self: user.value, workspace, member, joinedWorkspaces: joinedWorkspaces.value }))
  }
}
