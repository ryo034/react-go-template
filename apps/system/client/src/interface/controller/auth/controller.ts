import { Otp } from "~/domain/auth"
import { Email } from "~/domain/shared"
import { AuthUseCase, StartWithEmailInput, VerifyOtpInput } from "~/usecase"

export class AuthController {
  constructor(private readonly useCase: AuthUseCase) {}

  async startWithEmail(email: string): Promise<null | Error> {
    const e = Email.create(email)
    if (e.isErr) {
      return e.error
    }
    const input: StartWithEmailInput = { email: e.value }
    return await this.useCase.startWithEmail(input)
  }

  async createByOAuth(): Promise<null | Error> {
    return await this.useCase.createByOAuth()
  }

  async verifyOtp(email: Email, otp: string): Promise<null | Error> {
    const otpRes = Otp.create(otp)
    if (otpRes.isErr) {
      return otpRes.error
    }
    const input: VerifyOtpInput = { email, otp: otpRes.value }
    return await this.useCase.verifyOtp(input)
  }

  async findInvitationByToken(token: string): Promise<null | Error> {
    return await this.useCase.findInvitationByToken({ token })
  }

  async proceedToInvitation(token: string, email: Email): Promise<null | Error> {
    return await this.useCase.proceedToInvitation({ token, email })
  }
}
