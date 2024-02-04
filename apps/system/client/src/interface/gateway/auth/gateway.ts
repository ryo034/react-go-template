import { Result } from "true-myth"
import { Email } from "~/domain"
import { AuthRepository, CustomToken, Otp } from "~/domain/auth"
import { AuthDriver, AuthProviderDriver } from "~/driver"
import { PromiseResult } from "~/infrastructure/shared/result"
import { AuthGatewayAdapter } from "~/interface/gateway/auth"

export class AuthGateway implements AuthRepository {
  constructor(
    private readonly driver: AuthDriver,
    private readonly apDriver: AuthProviderDriver,
    private readonly adapter: AuthGatewayAdapter
  ) {}

  async startWithEmail(email: Email): PromiseResult<null, Error> {
    const res = await this.driver.startWithEmail(email)
    if (res.isErr) {
      return Result.err(res.error)
    }
    return Result.ok(null)
  }

  async verifyOtp(email: Email, otp: Otp): PromiseResult<CustomToken, Error> {
    const res = await this.driver.verifyOtp(email, otp)
    if (res.isErr) {
      return Result.err(res.error)
    }
    return this.adapter.adaptJwt(res.value)
  }

  async signInWithCustomToken(customToken: CustomToken): PromiseResult<null, Error> {
    const res = await this.apDriver.signInWithCustomToken(customToken)
    if (res.isErr) {
      return Result.err(res.error)
    }
    return Result.ok(null)
  }
}
