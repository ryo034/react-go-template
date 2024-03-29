import { Result } from "true-myth"
import type { AccountFullName, InvitationId, Me, MeRepository, MemberProfile } from "~/domain"
import type { AuthProviderDriver, MeDriver } from "~/driver"
import { AuthProviderCurrentUserNotFoundError } from "~/infrastructure/error"
import type { PromiseResult } from "~/infrastructure/shared/result"
import type { MeGatewayAdapter } from "~/interface/gateway/me/adapter"

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
      return Result.err(new AuthProviderCurrentUserNotFoundError("current user not found"))
    }
    const res = await this.driver.find()
    if (res.isErr) {
      return Result.err(res.error)
    }
    return this.adapter.adapt(res.value)
  }

  async acceptInvitation(invitationId: InvitationId): PromiseResult<Me, Error> {
    const res = await this.driver.acceptInvitation(invitationId)
    if (res.isErr) {
      return Result.err(res.error)
    }
    return this.adapter.adapt(res.value)
  }

  async updateProfile(name: AccountFullName): PromiseResult<Me, Error> {
    const res = await this.driver.updateProfile(name)
    if (res.isErr) {
      return Result.err(res.error)
    }
    return this.adapter.adapt(res.value)
  }

  async updatePhoto(file: File): PromiseResult<Me, Error> {
    const res = await this.driver.updatePhoto(file)
    if (res.isErr) {
      return Result.err(res.error)
    }
    return this.adapter.adapt(res.value)
  }

  async removePhoto(): PromiseResult<Me, Error> {
    const res = await this.driver.removePhoto()
    if (res.isErr) {
      return Result.err(res.error)
    }
    return this.adapter.adapt(res.value)
  }

  async updateMemberProfile(profile: MemberProfile): PromiseResult<Me, Error> {
    const res = await this.driver.updateMemberProfile(profile)
    if (res.isErr) {
      return Result.err(res.error)
    }
    return this.adapter.adapt(res.value)
  }

  async leaveWorkspace(): PromiseResult<null, Error> {
    return this.driver.leaveWorkspace()
  }
}
