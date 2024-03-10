import type { ApiErrorHandler } from "shared-network"
import { Result } from "true-myth"
import type { Me } from "~/domain"
import type { MyCustomGA } from "~/infrastructure/logger/ga4"

export class LoggerDriver {
  constructor(
    private readonly gaClient: MyCustomGA,
    private readonly errorHandler: ApiErrorHandler
  ) {}

  initialize(): void {
    this.gaClient.initialize(import.meta.env.VITE_GA_MEASUREMENT_ID)
  }

  sendUser(me: Me): Result<null, Error> {
    try {
      this.gaClient.gtag("set", "user_properties", {
        user_id: me.self.id.value.asString
      })
      return Result.ok(null)
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }

  sendLocation(_page: string): void {
    this.gaClient.gtag("event", "page_view", {})
  }
}
