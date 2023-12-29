import { MeRepository } from "~/domain"
import { MeUseCaseInput, MeUseCaseLoginInputData, MeUseCaseOutput } from "~/usecase"

export class MeInteractor implements MeUseCaseInput {
  constructor(private readonly repository: MeRepository, private readonly presenter: MeUseCaseOutput) {}

  private async loginWithEmailVerify(): Promise<Error | null> {
    const res = await this.repository.login()
    if (res.isErr) {
      return res.error
    }
    this.presenter.set(res.value)
    if (res.value.emailNotVerified) {
      const sendEmailVerificationRes = await this.repository.sendEmailVerification()
      if (sendEmailVerificationRes.isErr) {
        return sendEmailVerificationRes.error
      }
    }
    return null
  }

  async login(data: MeUseCaseLoginInputData): Promise<Error | null> {
    this.presenter.setIsLoading(true)
    const fbRes = await this.repository.signInWithEmailAndPassword(data.email, data.password)
    if (fbRes.isErr) {
      this.presenter.setIsLoading(false)
      return fbRes.error
    }
    this.presenter.setIsLoading(false)
    await this.loginWithEmailVerify()
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

  async verifyEmail(): Promise<Error | null> {
    const res = await this.repository.verifyEmail()
    if (res.isErr) {
      return res.error
    }
    return null
  }

  async checkEmailVerified(): Promise<Error | null> {
    const res = await this.repository.reloadAuth()
    if (res.isErr) {
      return res.error
    }
    this.presenter.set(res.value)
    return null
  }

  async sendEmailVerification(): Promise<Error | null> {
    const res = await this.repository.sendEmailVerification()
    if (res.isErr) {
      return res.error
    }
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
