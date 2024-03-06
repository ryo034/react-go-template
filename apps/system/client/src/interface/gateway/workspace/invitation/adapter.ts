import { Result } from "true-myth"
import { Invitation, InvitationId, Invitations, MemberDisplayName } from "~/domain"
import { AppDateTime, Email } from "~/domain/shared"
import { components } from "~/generated/schema/openapi/systemApi"
import { AdapterError } from "~/infrastructure/error"
import { MemberGatewayAdapter } from "../member"

export class InvitationGatewayAdapter {
  constructor(private readonly memberAdapter: MemberGatewayAdapter) {}
  adapt(invitation: components["schemas"]["Invitation"]): Result<Invitation, Error> {
    if (invitation === undefined || invitation === null) {
      console.error(new AdapterError(InvitationGatewayAdapter.name, this.adapt.name, "Invitation is required"))
      return Result.err(new Error("Invitation is required"))
    }

    const id = InvitationId.fromString(invitation.id)
    if (id.isErr) {
      return Result.err(id.error)
    }

    let displayName: MemberDisplayName | null = null
    if (invitation.displayName !== null && invitation.displayName.length > 0) {
      const displayNameRes = MemberDisplayName.create(invitation.displayName)
      if (displayNameRes.isErr) {
        return Result.err(displayNameRes.error)
      }
      displayName = displayNameRes.value
    }

    const inviteeEmail = Email.create(invitation.inviteeEmail)
    if (inviteeEmail.isErr) {
      return Result.err(inviteeEmail.error)
    }

    const expiredAt = AppDateTime.fromString(invitation.expiredAt)
    if (expiredAt.isErr) {
      return Result.err(expiredAt.error)
    }

    const inviter = this.memberAdapter.adapt(invitation.inviter)
    if (inviter.isErr) {
      return Result.err(inviter.error)
    }

    return Result.ok(
      Invitation.create({
        id: id.value,
        displayName,
        inviteeEmail: inviteeEmail.value,
        accepted: invitation.accepted,
        expiredAt: expiredAt.value,
        inviter: inviter.value
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
