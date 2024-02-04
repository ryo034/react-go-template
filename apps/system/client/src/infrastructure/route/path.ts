import { authPageRoute } from "~/pages/auth"
import { homePageRoute } from "~/pages/home"
import { verifyOtpPageRoute } from "~/pages/otp"
import { settingPageRoute } from "~/pages/settings"

export const routeMap = {
  auth: authPageRoute,
  verifyOtp: verifyOtpPageRoute,
  home: homePageRoute,
  settings: settingPageRoute
} as const

export const unauthenticatedRoutes = [routeMap.auth.toString(), routeMap.verifyOtp.toString()]

export const authRoutes = [routeMap.home.toString(), routeMap.settings.toString()]

export const unprotectedInitialPagePath = routeMap.auth
