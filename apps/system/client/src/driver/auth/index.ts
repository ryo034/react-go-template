import { ApiErrorHandler } from "shared-network"
import { Result } from "true-myth"
import { Email } from "~/domain"
import { Otp } from "~/domain/auth"
import { components } from "~/generated/schema/openapi/systemApi"
import { SystemAPIClient } from "~/infrastructure/openapi/client"
import { PromiseResult } from "~/infrastructure/shared"

export class AuthDriver {
  constructor(private readonly client: SystemAPIClient, private readonly errorHandler: ApiErrorHandler) {}

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
}
