import { AccountName, Me, User } from "~/domain"
import { AuthProviderCurrentUserNotFoundError } from "~/infrastructure/error"
import { MeUseCase } from "~/usecase"

export interface UpdateProfileNameInput {
  current: Me | null
  user: {
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

  async updateProfileName(i: UpdateProfileNameInput): Promise<null | Error> {
    if (i.current === null) {
      return new AuthProviderCurrentUserNotFoundError("current user not found")
    }
    const name = AccountName.create(i.user.name)
    if (name.isErr) {
      return name.error
    }
    return await this.useCase.updateProfile({
      user: User.create({ id: i.current?.self.id, name: name.value, email: i.current.self.email })
    })
  }
}
