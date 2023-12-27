import { Email, Password } from "~/domain"
import { MeUseCaseInput } from "~/usecase"

export class MeController {
  constructor(private readonly useCase: MeUseCaseInput) {}

  async login(email: string, password: string): Promise<null | Error> {
    const emailObj = Email.create(email)
    if (emailObj.isErr) {
      return emailObj.error
    }
    const passwordObj = Password.create(password)
    if (passwordObj.isErr) {
      return passwordObj.error
    }
    return await this.useCase.login(emailObj.value, passwordObj.value)
  }

  async signOut(): Promise<null | Error> {
    return await this.useCase.signOut()
  }

  async find(): Promise<null | Error> {
    return await this.useCase.find()
  }

  async verifyEmail(): Promise<null | Error> {
    return await this.useCase.verifyEmail()
  }

  async sendEmailVerification(): Promise<null | Error> {
    return await this.useCase.sendEmailVerification()
  }

  async checkEmailVerified(): Promise<null | Error> {
    return await this.useCase.checkEmailVerified()
  }
}
