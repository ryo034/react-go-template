export const routeMap = {
  auth: "/",
  verifyOtp: "/verify-otp",
  account: "/account",
  home: "/home",
  settings: "/settings"
} as const

export const unauthenticatedRoutes = [routeMap.auth.toString(), routeMap.verifyOtp.toString()]

export const authRoutes = [routeMap.account.toString(), routeMap.home.toString(), routeMap.settings.toString()]

export const unprotectedInitialPagePath = routeMap.auth
