import { Result } from "true-myth"
import { Email, Me, MeRepository, Password } from "~/domain"
import { AuthProviderDriver } from "~/driver"
import { MeDriver } from "~/driver/me/driver"
import { AuthProviderCurrentUserNotFoundError } from "~/infrastructure/error"
import { PromiseResult } from "~/infrastructure/shared/result"
import { MeGatewayAdapter } from "~/interface/gateway/me/adapter"

export class MeGateway implements MeRepository {
  constructor(
    private readonly driver: MeDriver,
    private readonly apDriver: AuthProviderDriver,
    private readonly adapter: MeGatewayAdapter
  ) {}

  async login(): PromiseResult<Me, Error> {
    if (this.apDriver.currentUser === null) {
      return Result.err(new Error("user is not logged in"))
    }
    const res = await this.driver.login()
    if (res.isErr) {
      return Result.err(res.error)
    }
    return this.adapter.adapt(res.value.me)
  }

  async sendEmailVerification(): PromiseResult<null, Error> {
    const res = await this.apDriver.sendEmailVerification()
    if (res.isErr) {
      return Result.err(res.error)
    }
    return Result.ok(null)
  }

  async signOut(): PromiseResult<null, Error> {
    return this.apDriver.signOut()
  }

  async reloadAuth(): PromiseResult<Me, Error> {
    await this.apDriver.reload()
    if (this.apDriver.currentUser === null) {
      return Result.err(new AuthProviderCurrentUserNotFoundError("currentUser is null"))
    }
    const res = await this.driver.find()
    if (res.isErr) {
      return Result.err(res.error)
    }
    return this.adapter.adapt(res.value.me)
  }

  async signInWithEmailAndPassword(email: Email, password: Password): PromiseResult<null, Error> {
    const res = await this.apDriver.signInWithEmailAndPassword(email, password)
    if (res.isErr) {
      return Result.err(res.error)
    }
    return Result.ok(null)
  }

  async verifyEmail(): PromiseResult<null, Error> {
    const res = await this.driver.verifyEmail()
    if (res.isErr) {
      return Result.err(res.error)
    }
    return Result.ok(null)
  }

  async find(): PromiseResult<Me, Error> {
    if (this.apDriver.currentUser === null) {
      return Result.err(new Error("user is not logged in"))
    }
    const res = await this.driver.find()
    if (res.isErr) {
      return Result.err(res.error)
    }
    return this.adapter.adapt(res.value.me)
  }
}
