import { MeUseCaseInput } from "~/usecase"

export class MeController {
  constructor(private readonly useCase: MeUseCaseInput) {}

  async login(): Promise<null | Error> {
    return await this.useCase.login()
  }

  async signOut(): Promise<null | Error> {
    return await this.useCase.signOut()
  }

  async find(): Promise<null | Error> {
    return await this.useCase.find()
  }
}
