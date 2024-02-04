import { Result } from "true-myth"
import { Email } from "~/domain"
import { AuthRepository, Jwt, Otp } from "~/domain/auth"
import { AuthDriver, AuthProviderDriver } from "~/driver"
import { PromiseResult } from "~/infrastructure/shared/result"
import { AuthGatewayAdapter } from "~/interface/gateway/auth"

export class AuthGateway implements AuthRepository {
  constructor(
    private readonly driver: AuthDriver,
    private readonly apDriver: AuthProviderDriver,
    private readonly adapter: AuthGatewayAdapter
  ) {}

  async startWithEmail(email: Email): PromiseResult<Otp, Error> {
    const res = await this.driver.startWithEmail(email)
    if (res.isErr) {
      return Result.err(res.error)
    }
    return this.adapter.adaptOtp(res.value)
  }

  async verifyOtp(email: Email, otp: Otp): PromiseResult<Jwt, Error> {
    const res = await this.driver.verifyOtp(email, otp)
    if (res.isErr) {
      return Result.err(res.error)
    }
    return this.adapter.adaptJwt(res.value)
  }

  async signInWithCustomToken(jwt: Jwt): PromiseResult<null, Error> {
    const res = await this.apDriver.signInWithCustomToken(jwt)
    if (res.isErr) {
      return Result.err(res.error)
    }
    return Result.ok(null)
  }
}
