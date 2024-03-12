export const MembershipStatusList = {
  Active: "active",
  Left: "left"
} as const

export type MembershipStatus = (typeof MembershipStatusList)[keyof typeof MembershipStatusList]
