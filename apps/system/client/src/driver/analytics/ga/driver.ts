import { ApiErrorHandler } from "shared-network"
import { Result } from "true-myth"
import { Me } from "~/domain"
import { MyCustomGA } from "~/infrastructure/analytics"

export class GoogleAnalyticsDriver {
  constructor(private readonly client: MyCustomGA, private readonly errorHandler: ApiErrorHandler) {}

  initialize(): void {
    this.client.initialize(import.meta.env.VITE_GA_MEASUREMENT_ID)
  }

  sendUser(me: Me): Result<null, Error> {
    try {
      this.client.gtag("set", "user_properties", {
        user_id: me.self.id.value.asString
      })
      return Result.ok(null)
    } catch (e) {
      return Result.err(this.errorHandler.adapt(e))
    }
  }

  sendLocation(_page: string): void {
    this.client.gtag("event", "page_view", {})
  }
}
