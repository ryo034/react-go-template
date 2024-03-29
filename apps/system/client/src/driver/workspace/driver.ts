import type { ApiErrorHandler } from "shared-network"
import { Result } from "true-myth"
import type {
  Invitation,
  Invitees,
  MemberId,
  SelectableRole,
  WorkspaceCreateInput,
  WorkspaceId,
  WorkspaceName,
  WorkspaceSubdomain
} from "~/domain"
import type { components } from "~/generated/schema/openapi/systemApi"
import type { SystemAPIClient } from "~/infrastructure/openapi/client"
import type { PromiseResult } from "~/infrastructure/shared/result"
import type { FirebaseDriver } from "../firebase/driver"

export class WorkspaceDriver {
  constructor(
    private readonly client: SystemAPIClient,
    private readonly errorHandler: ApiErrorHandler,
    private readonly fbDriver: FirebaseDriver
  ) {}

  async create(i: WorkspaceCreateInput): PromiseResult<components["schemas"]["Workspace"], Error> {
    try {
      const res = await this.client.POST("/api/v1/workspaces", {
        body: {
          subdomain: i.subdomain.value
        }
      })
      return res.data ? Result.ok(res.data.workspace) : Result.err(this.errorHandler.adapt(res))
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }

  async findAllMembers(): PromiseResult<components["schemas"]["Members"], Error> {
    try {
      const res = await this.client.GET("/api/v1/members")
      return res.data ? Result.ok(res.data.members) : Result.err(this.errorHandler.adapt(res))
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }

  async inviteMembers(invitees: Invitees): PromiseResult<components["schemas"]["Invitations"], Error> {
    try {
      const res = await this.client.POST("/api/v1/members/invitations/bulk", {
        body: {
          invitees: invitees.values.map((i) => ({
            email: i.email.value,
            name: i.name.value
          }))
        }
      })
      return res.data ? Result.ok(res.data.successfulInvitations) : Result.err(this.errorHandler.adapt(res))
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }

  async findAllInvitations(): PromiseResult<components["schemas"]["Invitations"], Error> {
    try {
      const res = await this.client.GET("/api/v1/invitations", { params: {} })
      return res.data ? Result.ok(res.data.invitations) : Result.err(this.errorHandler.adapt(res))
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }

  async resendInvitation(invitation: Invitation): PromiseResult<components["schemas"]["Invitation"], Error> {
    try {
      const res = await this.client.POST("/api/v1/members/invitations/{invitationId}/resend", {
        params: { path: { invitationId: invitation.id.value.asString } }
      })
      return res.data ? Result.ok(res.data) : Result.err(this.errorHandler.adapt(res))
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }

  async revokeInvitation(invitation: Invitation): PromiseResult<components["schemas"]["Invitations"], Error> {
    try {
      const res = await this.client.POST("/api/v1/members/invitations/{invitationId}/revoke", {
        params: { path: { invitationId: invitation.id.value.asString } }
      })
      return res.data ? Result.ok(res.data) : Result.err(this.errorHandler.adapt(res))
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }

  async updateMemberRole(
    memberId: MemberId,
    role: SelectableRole
  ): PromiseResult<components["schemas"]["Member"], Error> {
    try {
      const res = await this.client.PUT("/api/v1/members/{memberId}/role", {
        params: { path: { memberId: memberId.value.asString } },
        body: { role }
      })
      await this.fbDriver.refreshToken()
      return res.data ? Result.ok(res.data.member) : Result.err(this.errorHandler.adapt(res))
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }

  async updateWorkspace(
    workspaceId: WorkspaceId,
    name: WorkspaceName,
    subdomain: WorkspaceSubdomain
  ): PromiseResult<components["schemas"]["Workspace"], Error> {
    try {
      const res = await this.client.PUT("/api/v1/workspaces/{workspaceId}", {
        params: { path: { workspaceId: workspaceId.value.asString } },
        body: { name: name.value, subdomain: subdomain.value }
      })
      return res.data ? Result.ok(res.data.workspace) : Result.err(this.errorHandler.adapt(res))
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }

  async leaveWorkspace(memberId: MemberId): PromiseResult<null, Error> {
    try {
      const res = await this.client.DELETE("/api/v1/members/{memberId}", {
        params: { path: { memberId: memberId.value.asString } }
      })
      return res.data ? Result.ok(null) : Result.err(this.errorHandler.adapt(res))
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }
}
