export const routeMap = {
  auth: "/",
  verifyOtp: "/verify-otp",
  account: "/account",
  home: "/home",
  items: "/items",
  transactions: "/transactions",
  analytics: "/analytics",
  creatures: "/creatures",
  settings: "/settings"
} as const

export const unauthenticatedRoutes = [routeMap.auth.toString(), routeMap.verifyOtp.toString()]

export const authRoutes = [
  routeMap.account.toString(),
  routeMap.home.toString(),
  routeMap.transactions.toString(),
  routeMap.items.toString(),
  routeMap.analytics.toString(),
  routeMap.creatures.toString(),
  routeMap.settings.toString()
]

export const unprotectedInitialPagePath = routeMap.auth
