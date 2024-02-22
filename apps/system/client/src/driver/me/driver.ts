import { ApiErrorHandler } from "shared-network"
import { Result } from "true-myth"
import { User } from "~/domain"
import { components } from "~/generated/schema/openapi/systemApi"
import { SystemAPIClient } from "~/infrastructure/openapi/client"
import { PromiseResult } from "~/infrastructure/shared/result"

export class MeDriver {
  constructor(private readonly client: SystemAPIClient, private readonly errorHandler: ApiErrorHandler) {}

  async find(): PromiseResult<components["schemas"]["Me"], Error> {
    try {
      const res = await this.client.GET("/api/v1/me")
      return res.data ? Result.ok(res.data.me) : Result.err(this.errorHandler.adapt(res))
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }

  async updateProfile(user: User): PromiseResult<components["schemas"]["Me"], Error> {
    try {
      const res = await this.client.PUT("/api/v1/me/profile", {
        body: {
          user: {
            userId: user.id.value.asString,
            email: user.email.value,
            name: user.name?.value || "",
            phoneNumber: ""
          }
        }
      })
      return res.data ? Result.ok(res.data.me) : Result.err(this.errorHandler.adapt(res))
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }
}
