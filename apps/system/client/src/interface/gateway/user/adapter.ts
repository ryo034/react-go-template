import { Result } from "true-myth"
import { AccountFullName, AccountId } from "~/domain/account"
import { Email } from "~/domain/shared"
import { Photo } from "~/domain/shared/photo"
import { User } from "~/domain/user"
import type { components } from "~/generated/schema/openapi/systemApi"
import { AdapterError } from "~/infrastructure/error"

export class UserGatewayAdapter {
  adapt(user: components["schemas"]["User"]): Result<User, Error> {
    if (user === undefined || user === null) {
      console.error(new AdapterError(UserGatewayAdapter.name, this.adapt.name, "user is required"))
      return Result.err(new Error("User is not found"))
    }

    const id = AccountId.fromString(user.userId)
    if (id.isErr) {
      return Result.err(id.error)
    }

    const email = Email.create(user.email)
    if (email.isErr) {
      return Result.err(email.error)
    }

    let name: AccountFullName | undefined = undefined
    if (user.name) {
      const tmpName = AccountFullName.create(user.name)
      if (tmpName.isErr) {
        return Result.err(tmpName.error)
      }
      name = tmpName.value
    }

    let photo: Photo | undefined = undefined
    if (user.photo) {
      const tmpPhoto = Photo.fromString(user.photo)
      if (tmpPhoto.isErr) {
        return Result.err(tmpPhoto.error)
      }
      photo = tmpPhoto.value
    }

    return Result.ok(User.create({ id: id.value, email: email.value, name, photo }))
  }
}
