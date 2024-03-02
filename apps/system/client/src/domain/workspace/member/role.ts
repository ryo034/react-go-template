export const MemberRoleList = {
  Owner: "owner",
  Admin: "admin",
  Member: "member",
  Guest: "guest"
} as const

export type MemberRole = (typeof MemberRoleList)[keyof typeof MemberRoleList]
