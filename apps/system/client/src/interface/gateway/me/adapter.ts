import { Result } from "true-myth"
import { AccountId, AccountName } from "~/domain/account"
import { Me } from "~/domain/me"
import { Email } from "~/domain/shared"
import { User } from "~/domain/user"
import { components } from "~/generated/schema/openapi/systemApi"
import { AdapterError, AuthProviderCurrentUserNotFoundError } from "~/infrastructure/error"

export class MeGatewayAdapter {
  adapt(me: components["schemas"]["Me"]): Result<Me, Error> {
    if (me === undefined || me === null) {
      console.error(new AdapterError(MeGatewayAdapter.name, this.adapt.name, "me is required"))
      return Result.err(new AuthProviderCurrentUserNotFoundError("User is not found"))
    }

    const { userId, email, firstName, lastName } = me.user

    const id = AccountId.fromString(userId)
    if (id.isErr) {
      return Result.err(id.error)
    }

    const fn = AccountName.create(firstName)
    if (fn.isErr) {
      return Result.err(fn.error)
    }

    const ln = AccountName.create(lastName)
    if (ln.isErr) {
      return Result.err(ln.error)
    }

    const em = Email.create(email)
    if (em.isErr) {
      return Result.err(em.error)
    }
    const user = User.create({ id: id.value, firstName: fn.value, lastName: ln.value, email: em.value })
    const emailVerified = me.emailVerified === undefined ? false : me.emailVerified

    return Result.ok(Me.create({ user, emailVerified }))
  }
}
