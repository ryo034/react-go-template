import type { ApiErrorHandler } from "shared-network"
import { Result } from "true-myth"
import { type AccountFullName, type InvitationId, type MemberProfile, User, type WorkspaceId } from "~/domain"
import type { components } from "~/generated/schema/openapi/systemApi"
import type { SystemAPIClient } from "~/infrastructure/openapi/client"
import type { PromiseResult } from "~/infrastructure/shared/result"
import type { FirebaseDriver } from "../firebase/driver"

export class MeDriver {
  constructor(
    private readonly client: SystemAPIClient,
    private readonly fbDriver: FirebaseDriver,
    private readonly errorHandler: ApiErrorHandler
  ) {}

  async find(): PromiseResult<components["schemas"]["Me"], Error> {
    try {
      const res = await this.client.GET("/api/v1/me")
      return res.data ? Result.ok(res.data.me) : Result.err(this.errorHandler.adapt(res))
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }

  async acceptInvitation(invitationId: InvitationId): PromiseResult<components["schemas"]["Me"], Error> {
    try {
      const res = await this.client.POST("/api/v1/members/invitations/{invitationId}/accept", {
        params: { path: { invitationId: invitationId.value.asString } }
      })
      return res.data ? Result.ok(res.data.me) : Result.err(this.errorHandler.adapt(res))
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }

  async updateProfile(name: AccountFullName): PromiseResult<components["schemas"]["Me"], Error> {
    try {
      const res = await this.client.PUT("/api/v1/me/profile", {
        body: {
          profile: {
            name: name.value
          }
        }
      })
      return res.data ? Result.ok(res.data.me) : Result.err(this.errorHandler.adapt(res))
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }

  async updatePhoto(file: File): PromiseResult<components["schemas"]["Me"], Error> {
    try {
      const res = await this.client.PUT("/api/v1/me/profile/photo", {
        body: { photo: "" },
        bodySerializer(body) {
          const fd = new FormData()
          fd.append("photo", file)
          return fd
        }
      })
      return res.data ? Result.ok(res.data.me) : Result.err(this.errorHandler.adapt(res))
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }

  async removePhoto(): PromiseResult<components["schemas"]["Me"], Error> {
    try {
      const res = await this.client.DELETE("/api/v1/me/profile/photo")
      return res.data ? Result.ok(res.data.me) : Result.err(this.errorHandler.adapt(res))
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }

  async updateMemberProfile(profile: MemberProfile): PromiseResult<components["schemas"]["Me"], Error> {
    try {
      const res = await this.client.PUT("/api/v1/me/member/profile", {
        body: {
          memberProfile: {
            displayName: profile.displayName?.value || "",
            bio: profile.bio.value,
            idNumber: profile.idNumber ? profile.idNumber.value : ""
          }
        }
      })
      return res.data ? Result.ok(res.data.me) : Result.err(this.errorHandler.adapt(res))
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }

  async leaveWorkspace(): PromiseResult<null, Error> {
    try {
      const res = await this.client.POST("/api/v1/me/workspace/leave")
      await this.fbDriver.refreshToken()
      return res.data ? Result.ok(null) : Result.err(this.errorHandler.adapt(res))
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }
}
