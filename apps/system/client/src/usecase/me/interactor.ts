import { MeRepository, User } from "~/domain"
import { MeUseCaseOutput } from "~/usecase"

export interface MeUseCase {
  signOut(): Promise<Error | null>
  find(): Promise<Error | null>
  updateProfile(i: UpdateProfileInput): Promise<Error | null>
}

export type UpdateProfileInput = {
  user: User
}

export class MeInteractor implements MeUseCase {
  constructor(private readonly repository: MeRepository, private readonly presenter: MeUseCaseOutput) {}

  async signOut(): Promise<Error | null> {
    const res = await this.repository.signOut()
    if (res.isErr) {
      return res.error
    }
    this.presenter.clear()
    return null
  }

  async find(): Promise<Error | null> {
    this.presenter.setIsLoading(true)
    const res = await this.repository.find()
    if (res.isErr) {
      this.presenter.setIsLoading(false)
      return res.error
    }
    this.presenter.set(res.value)
    this.presenter.setIsLoading(false)
    return null
  }

  async updateProfile(i: UpdateProfileInput): Promise<Error | null> {
    const res = await this.repository.updateProfile(i.user)
    if (res.isErr) {
      return res.error
    }
    this.presenter.set(res.value)
    return null
  }
}
