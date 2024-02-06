import { Result } from "true-myth"
import { AccountName, Me, MeRepository, User } from "~/domain"
import { AuthProviderDriver, MeDriver } from "~/driver"
import { PromiseResult } from "~/infrastructure/shared/result"
import { MeGatewayAdapter } from "~/interface/gateway/me/adapter"

export class MeGateway implements MeRepository {
  constructor(
    private readonly driver: MeDriver,
    private readonly apDriver: AuthProviderDriver,
    private readonly adapter: MeGatewayAdapter
  ) {}

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

  async updateProfile(user: User): PromiseResult<Me, Error> {
    const res = await this.driver.updateProfile(user)
    if (res.isErr) {
      return Result.err(res.error)
    }
    return this.adapter.adapt(res.value)
  }
}
