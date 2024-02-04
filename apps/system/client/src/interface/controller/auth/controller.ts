import { Email } from "~/domain"
import { Otp } from "~/domain/auth"
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

  async verifyOtp(email: Email, otp: string): Promise<null | Error> {
    const otpRes = Otp.create(otp)
    if (otpRes.isErr) {
      return otpRes.error
    }
    const input: VerifyOtpInput = { email, otp: otpRes.value }
    return await this.useCase.verifyOtp(input)
  }
}
