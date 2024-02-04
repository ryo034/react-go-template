import { AuthRepository } from "~/domain/auth"
import { AuthUseCaseOutput, StartWithEmailInput, VerifyOtpInput } from "~/usecase/auth"
import { MeUseCase } from "~/usecase/me"

export interface AuthUseCase {
  startWithEmail(i: StartWithEmailInput): Promise<Error | null>
  verifyOtp(i: VerifyOtpInput): Promise<Error | null>
}

export class AuthInteractor implements AuthUseCase {
  constructor(
    private readonly repository: AuthRepository,
    private readonly meUseCase: MeUseCase,
    private readonly presenter: AuthUseCaseOutput
  ) {}

  async startWithEmail(i: StartWithEmailInput): Promise<Error | null> {
    this.presenter.setIsLoading(true)
    this.presenter.setEmail(i.email)
    const res = await this.repository.startWithEmail(i.email)
    if (res.isErr) {
      this.presenter.setIsLoading(false)
      return res.error
    }
    this.presenter.setIsLoading(false)
    return null
  }

  async verifyOtp(i: VerifyOtpInput): Promise<Error | null> {
    this.presenter.setIsLoading(true)
    const res = await this.repository.verifyOtp(i.email, i.otp)
    if (res.isErr) {
      this.presenter.setIsLoading(false)
      return res.error
    }
    const singInRes = await this.repository.signInWithCustomToken(res.value)
    this.presenter.setIsLoading(false)
    if (singInRes.isErr) {
      return singInRes.error
    }
    return await this.meUseCase.find()
  }
}
