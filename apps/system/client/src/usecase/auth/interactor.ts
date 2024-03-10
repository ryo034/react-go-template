import type { AuthRepository } from "~/domain/auth"
import type {
  AuthUseCaseOutput,
  FindInvitationByTokenInput,
  ProceedInvitationByEmailInput,
  ProceedInvitationByOAuthInput,
  StartWithEmailInput,
  VerifyOtpInput
} from "~/usecase/auth"
import type { MeUseCase, MeUseCaseOutput } from "~/usecase/me"

export interface AuthUseCase {
  startWithEmail(i: StartWithEmailInput): Promise<Error | null>
  createByOAuth(): Promise<Error | null>
  verifyOtp(i: VerifyOtpInput): Promise<Error | null>
  findInvitationByToken(i: FindInvitationByTokenInput): Promise<Error | null>
  proceedInvitationByEmail(i: ProceedInvitationByEmailInput): Promise<Error | null>
  proceedInvitationByOAuth(i: ProceedInvitationByOAuthInput): Promise<Error | null>
}

export class AuthInteractor implements AuthUseCase {
  constructor(
    private readonly repository: AuthRepository,
    private readonly meUseCase: MeUseCase,
    private readonly presenter: AuthUseCaseOutput,
    private readonly mePresenter: MeUseCaseOutput
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

  async startWithGoogle(): Promise<Error | null> {
    this.presenter.setIsLoading(true)
    const res = await this.repository.startWithGoogle()
    if (res.isErr) {
      this.presenter.setIsLoading(false)
      return res.error
    }
    return null
  }

  async createByOAuth(): Promise<Error | null> {
    this.presenter.setIsLoading(true)
    const res = await this.repository.authByOAuth()
    this.presenter.setIsLoading(false)
    if (res.isErr) {
      return res.error
    }
    this.mePresenter.set(res.value)
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

  async findInvitationByToken(i: FindInvitationByTokenInput): Promise<Error | null> {
    this.presenter.setReceivedInvitationIsLoading(true)
    const res = await this.repository.findInvitationByToken(i.token)
    this.presenter.setReceivedInvitationIsLoading(false)
    if (res.isErr) {
      return res.error
    }
    this.presenter.setReceivedInvitation(res.value)
    return null
  }

  async proceedInvitationByEmail(i: ProceedInvitationByEmailInput): Promise<Error | null> {
    this.presenter.setIsLoading(true)
    const res = await this.repository.proceedInvitationByEmail(i.token, i.email)
    this.presenter.setIsLoading(false)
    if (res.isErr) {
      return res.error
    }
    return null
  }

  async proceedInvitationByOAuth(i: ProceedInvitationByOAuthInput): Promise<Error | null> {
    this.presenter.setIsLoading(true)
    const res = await this.repository.proceedInvitationByOAuth(i.token)
    this.presenter.setIsLoading(false)
    if (res.isErr) {
      return res.error
    }
    this.mePresenter.set(res.value)
    return null
  }
}
