import { Result } from "true-myth"
import { Member, MemberDisplayName, MemberId, MemberIdNumber, MemberProfile, Members } from "~/domain/workspace/member"
import { components } from "~/generated/schema/openapi/systemApi"
import { AdapterError } from "~/infrastructure/error"
import { UserGatewayAdapter } from "~/interface/gateway/user"

export class MemberGatewayAdapter {
  constructor(private readonly userAdapter: UserGatewayAdapter) {}
  adapt(member: components["schemas"]["Member"]): Result<Member, Error> {
    if (member === undefined || member === null) {
      console.error(new AdapterError(MemberGatewayAdapter.name, this.adapt.name, "member is required"))
      return Result.err(new Error("Member is not found"))
    }

    const user = this.userAdapter.adapt(member.user)
    if (user.isErr) {
      return Result.err(user.error)
    }

    const id = MemberId.fromString(member.profile.id)
    if (id.isErr) {
      return Result.err(id.error)
    }

    const displayName = MemberDisplayName.create(member.profile.displayName ?? "")
    if (displayName.isErr) {
      return Result.err(displayName.error)
    }

    const idNumber = MemberIdNumber.create(member.profile.idNumber ?? "")
    if (idNumber.isErr) {
      return Result.err(idNumber.error)
    }

    const profile = MemberProfile.create({
      id: id.value,
      displayName: displayName.value,
      idNumber: idNumber.value
    })

    return Result.ok(Member.create({ user: user.value, profile }))
  }

  adaptAll(members: components["schemas"]["Members"]): Result<Members, Error> {
    const vs: Member[] = []
    for (const m of members) {
      const res = this.adapt(m)
      if (res.isErr) {
        return Result.err(res.error)
      }
      vs.push(res.value)
    }
    return Result.ok(Members.create(vs))
  }
}
