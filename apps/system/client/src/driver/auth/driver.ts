import type { ApiErrorHandler } from "shared-network"
import { Result } from "true-myth"
import type { Otp } from "~/domain/auth"
import type { Email } from "~/domain/shared"
import type { components } from "~/generated/schema/openapi/systemApi"
import type { SystemAPIClient } from "~/infrastructure/openapi/client"
import type { PromiseResult } from "~/infrastructure/shared"

export class AuthDriver {
  constructor(
    private readonly client: SystemAPIClient,
    private readonly errorHandler: ApiErrorHandler
  ) {}

  async startWithEmail(email: Email): PromiseResult<null, Error> {
    try {
      const res = await this.client.POST("/api/v1/auth/otp", { body: { email: email.value } })
      return res.data ? Result.ok(res) : Result.err(this.errorHandler.adapt(res))
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }

  async verifyOtp(email: Email, otp: Otp): PromiseResult<components["schemas"]["JwtToken"], Error> {
    try {
      const res = await this.client.POST("/api/v1/auth/otp/verify", { body: { email: email.value, otp: otp.value } })
      return res.data ? Result.ok(res.data) : Result.err(this.errorHandler.adapt(res))
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }

  async findInvitationByToken(token: string): PromiseResult<components["schemas"]["ReceivedInvitation"], Error> {
    try {
      const res = await this.client.GET("/api/v1/auth/invitations", { params: { query: { token } } })
      return res.data ? Result.ok(res.data.receivedInvitation) : Result.err(this.errorHandler.adapt(res))
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }

  async proceedInvitationByEmail(token: string, email: Email): PromiseResult<null, Error> {
    try {
      const res = await this.client.POST("/api/v1/auth/invitations/process/email", {
        body: { token, email: email.value }
      })
      return res.data ? Result.ok(null) : Result.err(this.errorHandler.adapt(res))
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }

  async proceedInvitationByOAuth(token: string): PromiseResult<components["schemas"]["Me"], Error> {
    try {
      const res = await this.client.POST("/api/v1/auth/invitations/process/oauth", { body: { token } })
      return res.data ? Result.ok(res.data) : Result.err(this.errorHandler.adapt(res))
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }

  async authByOAuth(): PromiseResult<components["schemas"]["Me"], Error> {
    try {
      const res = await this.client.POST("/api/v1/auth/oauth")
      return res.data ? Result.ok(res.data) : Result.err(this.errorHandler.adapt(res))
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }
}
