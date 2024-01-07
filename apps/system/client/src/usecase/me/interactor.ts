import { MeRepository } from "~/domain"
import { MeUseCaseInput, MeUseCaseOutput, MeUseCaseSignUpInputData } from "~/usecase"

export class MeInteractor implements MeUseCaseInput {
  constructor(private readonly repository: MeRepository, private readonly presenter: MeUseCaseOutput) {}

  async login(): Promise<Error | null> {
    this.presenter.setIsLoading(true)
    const fbRes = await this.repository.login()
    if (fbRes.isErr) {
      this.presenter.setIsLoading(false)
      return fbRes.error
    }
    this.presenter.setIsLoading(false)
    return null
  }

  async signUp(data: MeUseCaseSignUpInputData): Promise<Error | null> {
    const res = await this.repository.signUp()
    if (res.isErr) {
      return res.error
    }
    this.presenter.clear()
    return null
  }

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
}
