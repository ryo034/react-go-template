import { AccountFullName, InvitationId, Me, MemberDisplayName, MemberIdNumber, User } from "~/domain"
import { MemberBio } from "~/domain/workspace/member/bio"
import { AuthProviderCurrentUserNotFoundError } from "~/infrastructure/error"
import { MeUseCase } from "~/usecase"

interface AcceptInvitationInput {
  invitationId: InvitationId
}

export interface UpdateProfileNameInput {
  current: Me | null
  user: {
    name: string
  }
}

export interface UpdateMemberProfileInput {
  displayName: string
  bio: string
  idNumber: string
}

export class MeController {
  constructor(private readonly useCase: MeUseCase) {}

  async signOut(): Promise<null | Error> {
    return await this.useCase.signOut()
  }

  async find(): Promise<null | Error> {
    return await this.useCase.find()
  }

  async acceptInvitation(i: AcceptInvitationInput): Promise<null | Error> {
    return await this.useCase.acceptInvitation({ invitationId: i.invitationId })
  }

  async updateProfileName(i: UpdateProfileNameInput): Promise<null | Error> {
    if (i.current === null) {
      return new AuthProviderCurrentUserNotFoundError("current user not found")
    }
    const name = AccountFullName.create(i.user.name)
    if (name.isErr) {
      return name.error
    }
    return await this.useCase.updateProfile({
      user: User.create({ id: i.current?.self.id, name: name.value, email: i.current.self.email })
    })
  }

  async updateMemberProfile(i: UpdateMemberProfileInput): Promise<null | Error> {
    const bio = MemberBio.create(i.bio)
    if (bio.isErr) {
      return bio.error
    }

    let displayName: MemberDisplayName | undefined = undefined
    if (i.displayName) {
      const dn = MemberDisplayName.create(i.displayName)
      if (dn.isErr) {
        return dn.error
      }
      displayName = dn.value
    }

    let idNumber: MemberIdNumber | undefined = undefined
    if (i.idNumber) {
      const idn = MemberIdNumber.create(i.idNumber)
      if (idn.isErr) {
        return idn.error
      }
      idNumber = idn.value
    }

    return await this.useCase.updateMemberProfile({
      bio: bio.value,
      displayName,
      idNumber
    })
  }
}
