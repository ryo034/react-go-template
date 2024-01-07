import { Result } from "true-myth"
import { Me, MeRepository } from "~/domain"
import { AuthProviderDriver } from "~/driver"
import { MeDriver } from "~/driver/me/driver"
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
    return this.adapter.adapt(res.value)
  }

  async signUp(): PromiseResult<null, Error> {
    return this.apDriver.signUp()
  }

  async signOut(): PromiseResult<null, Error> {
    return this.apDriver.signOut()
  }

  async find(): PromiseResult<Me, Error> {
    if (this.apDriver.currentUser === null) {
      return Result.err(new Error("user is not logged in"))
    }
    const res = await this.driver.find()
    if (res.isErr) {
      return Result.err(res.error)
    }
    return this.adapter.adapt(res.value)
  }
}
