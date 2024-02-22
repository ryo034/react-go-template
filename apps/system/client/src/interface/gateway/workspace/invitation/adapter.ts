import { Result } from "true-myth"
import { AppDateTime, Email, Invitation, InvitationId, Invitations, MemberDisplayName } from "~/domain"
import { components } from "~/generated/schema/openapi/systemApi"
import { AdapterError } from "~/infrastructure/error"

export class InvitationGatewayAdapter {
  adapt(invitation: components["schemas"]["Invitation"]): Result<Invitation, Error> {
    if (invitation === undefined || invitation === null) {
      console.error(new AdapterError(InvitationGatewayAdapter.name, this.adapt.name, "Invitation is required"))
      return Result.err(new Error("Invitation is required"))
    }

    const id = InvitationId.fromString(invitation.id)
    if (id.isErr) {
      return Result.err(id.error)
    }

    const displayName = MemberDisplayName.create(invitation.displayName)
    if (displayName.isErr) {
      return Result.err(displayName.error)
    }

    const inviteeEmail = Email.create(invitation.inviteeEmail)
    if (inviteeEmail.isErr) {
      return Result.err(inviteeEmail.error)
    }

    const expiredAt = AppDateTime.fromString(invitation.expiredAt)
    if (expiredAt.isErr) {
      return Result.err(expiredAt.error)
    }

    return Result.ok(
      Invitation.create({
        id: id.value,
        displayName: displayName.value,
        inviteeEmail: inviteeEmail.value,
        verified: invitation.verified,
        expiredAt: expiredAt.value
      })
    )
  }

  adaptAll(invitations: components["schemas"]["Invitations"]): Result<Invitations, Error> {
    const vs: Invitation[] = []
    for (const i of invitations) {
      const res = this.adapt(i)
      if (res.isErr) {
        return Result.err(res.error)
      }
      vs.push(res.value)
    }
    return Result.ok(Invitations.create(vs))
  }
}
