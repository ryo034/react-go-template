import { AccountId, AccountName, Email, User } from "~/domain"
import { MeUseCase } from "~/usecase"

export interface UpdateProfileInput {
  user: {
    userId: string
    email: string
    name: string
  }
}

export class MeController {
  constructor(private readonly useCase: MeUseCase) {}

  async signOut(): Promise<null | Error> {
    return await this.useCase.signOut()
  }

  async find(): Promise<null | Error> {
    return await this.useCase.find()
  }

  async updateProfile(i: UpdateProfileInput): Promise<null | Error> {
    const uId = AccountId.fromString(i.user.userId)
    if (uId.isErr) {
      return uId.error
    }
    const name = AccountName.create(i.user.name)
    if (name.isErr) {
      return name.error
    }
    const email = Email.create(i.user.email)
    if (email.isErr) {
      return email.error
    }
    return await this.useCase.updateProfile({
      user: User.create({ id: uId.value, name: name.value, email: email.value })
    })
  }
}
