import { Result } from "true-myth"
import {
  Member,
  MemberDisplayName,
  MemberId,
  MemberIdNumber,
  MemberProfile,
  MemberRole,
  MemberRoleList,
  Members
} from "~/domain/workspace/member"
import { MemberBio } from "~/domain/workspace/member/bio"
import { components } from "~/generated/schema/openapi/systemApi"
import { AdapterError } from "~/infrastructure/error"
import { UserGatewayAdapter } from "~/interface/gateway/user"

export class MemberGatewayAdapter {
  constructor(private readonly userAdapter: UserGatewayAdapter) {}
  adaptRole(role: components["schemas"]["MemberRole"]): Result<MemberRole, Error> {
    if (role === undefined || role === null) {
      console.error(new AdapterError(MemberGatewayAdapter.name, this.adaptRole.name, "role is required"))
      return Result.err(new Error("Role is not found"))
    }

    switch (role) {
      case "OWNER":
        return Result.ok(MemberRoleList.Owner)
      case "ADMIN":
        return Result.ok(MemberRoleList.Admin)
      case "MEMBER":
        return Result.ok(MemberRoleList.Member)
      case "GUEST":
        return Result.ok(MemberRoleList.Guest)
      default:
        console.error(new AdapterError(MemberGatewayAdapter.name, this.adaptRole.name, `role is invalid: ${role}`))
        return Result.err(new Error("Role is invalid"))
    }
  }

  adapt(member: components["schemas"]["Member"]): Result<Member, Error> {
    if (member === undefined || member === null) {
      console.error(new AdapterError(MemberGatewayAdapter.name, this.adapt.name, "member is required"))
      return Result.err(new Error("Member is not found"))
    }

    const user = this.userAdapter.adapt(member.user)
    if (user.isErr) {
      return Result.err(user.error)
    }

    const id = MemberId.fromString(member.id)
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

    const bio = MemberBio.create(member.profile.bio ?? "")
    if (bio.isErr) {
      return Result.err(bio.error)
    }

    const profile = MemberProfile.create({
      displayName: displayName.value,
      idNumber: idNumber.value,
      bio: bio.value
    })

    const role = this.adaptRole(member.role)
    if (role.isErr) {
      return Result.err(role.error)
    }

    return Result.ok(Member.create({ id: id.value, user: user.value, profile, role: role.value }))
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
