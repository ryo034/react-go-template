import { Result } from "true-myth"
import type { AuthRepository, CustomToken, Me, Otp, ReceivedInvitation } from "~/domain"
import type { Email } from "~/domain/shared"
import type { AuthDriver, AuthProviderDriver } from "~/driver"
import type { PromiseResult } from "~/infrastructure/shared/result"
import type { AuthGatewayAdapter, MeGatewayAdapter } from "~/interface/gateway"

export class AuthGateway implements AuthRepository {
  constructor(
    private readonly driver: AuthDriver,
    private readonly apDriver: AuthProviderDriver,
    private readonly adapter: AuthGatewayAdapter,
    private readonly meAdapter: MeGatewayAdapter
  ) {}

  async startWithEmail(email: Email): PromiseResult<null, Error> {
    const res = await this.driver.startWithEmail(email)
    if (res.isErr) {
      return Result.err(res.error)
    }
    return Result.ok(null)
  }

  async startWithGoogle(): PromiseResult<null, Error> {
    const res = await this.apDriver.startWithGoogle()
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

  async findInvitationByToken(token: string): PromiseResult<ReceivedInvitation, Error> {
    const res = await this.driver.findInvitationByToken(token)
    if (res.isErr) {
      return Result.err(res.error)
    }
    return this.meAdapter.adaptReceivedInvitation(res.value)
  }

  async proceedInvitationByEmail(token: string, email: Email): PromiseResult<null, Error> {
    const res = await this.driver.proceedInvitationByEmail(token, email)
    if (res.isErr) {
      return Result.err(res.error)
    }
    return Result.ok(null)
  }

  async proceedInvitationByOAuth(token: string): PromiseResult<Me, Error> {
    const res = await this.driver.proceedInvitationByOAuth(token)
    if (res.isErr) {
      return Result.err(res.error)
    }
    return this.meAdapter.adapt(res.value)
  }

  async authByOAuth(): PromiseResult<Me, Error> {
    const res = await this.driver.authByOAuth()
    if (res.isErr) {
      return Result.err(res.error)
    }
    return this.meAdapter.adapt(res.value)
  }
}
