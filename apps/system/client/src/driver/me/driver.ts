import { Result } from "true-myth"
import { AccountName } from "~/domain/account"
import { components } from "~/generated/schema/openapi/systemApi"
import { ErrorHandler } from "~/infrastructure/error/handler"
import { SystemAPIClient } from "~/infrastructure/openapi/client"
import { PromiseResult } from "~/infrastructure/shared/result"

export class MeDriver {
  constructor(private readonly client: SystemAPIClient) {}

  async login(): PromiseResult<components["schemas"]["Me"], Error> {
    try {
      const res = await this.client.POST("/login")
      return res.data ? Result.ok(res.data) : Result.err(ErrorHandler.adapt(res.error))
    } catch (e) {
      return Result.err(ErrorHandler.adapt(e))
    }
  }

  async signUp(firstName: AccountName, lastName: AccountName): PromiseResult<components["schemas"]["Me"], Error> {
    try {
      const res = await this.client.POST("/sign_up", { params: { query: { firstName, lastName } } })
      return res.data ? Result.ok(res.data) : Result.err(ErrorHandler.adapt(res.error))
    } catch (e) {
      return Result.err(ErrorHandler.adapt(e))
    }
  }

  async find(): PromiseResult<components["schemas"]["Me"], Error> {
    try {
      const res = await this.client.GET("/me")
      return res.data ? Result.ok(res.data) : Result.err(ErrorHandler.adapt(res.error))
    } catch (e) {
      return Result.err(ErrorHandler.adapt(e))
    }
  }
}
