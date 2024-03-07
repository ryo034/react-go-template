export const MemberRoleList = {
  Owner: "owner",
  Admin: "admin",
  Member: "member",
  Guest: "guest"
} as const

export type MemberRole = (typeof MemberRoleList)[keyof typeof MemberRoleList]

export const SelectableRoleList = [MemberRoleList.Admin, MemberRoleList.Member, MemberRoleList.Guest] as const

export type SelectableRole = (typeof SelectableRoleList)[number]
