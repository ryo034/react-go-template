import { Result } from "true-myth"
import { AccountId, AccountName } from "~/domain/account"
import { Me } from "~/domain/me"
import { Email } from "~/domain/shared"
import { User } from "~/domain/user"
import * as mePb from "~/generated/schema/api/me/v1/me_pb"
import { AdapterError, AuthProviderCurrentUserNotFoundError } from "~/infrastructure/error"

export class MeGatewayAdapter {
  adapt(me?: mePb.Me): Result<Me, Error> {
    if (me === undefined || me === null) {
      console.error(new AdapterError(MeGatewayAdapter.name, this.adapt.name, "me is required"))
      return Result.err(new AuthProviderCurrentUserNotFoundError("User is not found"))
    }

    if (me.role.value === undefined || me.role.value === null) {
      console.error(new AdapterError(MeGatewayAdapter.name, this.adapt.name, "me.role.value is invalid"))
      return Result.err(new AuthProviderCurrentUserNotFoundError("User is not found"))
    }

    if (me.role.value.user === undefined || me.role.value.user === null) {
      console.error(new AdapterError(MeGatewayAdapter.name, this.adapt.name, "me.role.value.user is invalid"))
      return Result.err(new AuthProviderCurrentUserNotFoundError("User is not found"))
    }

    const { firstName, lastName, email, userId } = me.role.value.user

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
