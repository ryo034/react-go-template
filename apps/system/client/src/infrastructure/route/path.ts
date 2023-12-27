export const routeMap = {
  login: "/",
  confirmEmail: "/confirm-email",
  account: "/account",
  home: "/home",
  items: "/items",
  transactions: "/transactions",
  analytics: "/analytics",
  creatures: "/creatures",
  settings: "/settings"
} as const

export const unauthenticatedRoutes = [routeMap.login.toString()]

export const authRoutes = [
  routeMap.account.toString(),
  routeMap.confirmEmail.toString(),
  routeMap.home.toString(),
  routeMap.transactions.toString(),
  routeMap.items.toString(),
  routeMap.analytics.toString(),
  routeMap.creatures.toString(),
  routeMap.settings.toString()
]

export const unprotectedInitialPagePath = routeMap.login
