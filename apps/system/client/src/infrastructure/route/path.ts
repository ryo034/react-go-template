import { authPageRoute } from "~/pages/auth"
import { homePageRoute } from "~/pages/home"
import { startInvitationPageRoute } from "~/pages/invitation"
import { membersPageRoute } from "~/pages/members"
import { onboardingSettingNamePageRoute } from "~/pages/onboarding/name"
import { onboardingSettingWorkspacePageRoute } from "~/pages/onboarding/workspace"
import { verifyOtpPageRoute } from "~/pages/otp"
import { receivedInvitationsPageRoute } from "~/pages/receivedInvitation"
import { settingPageRoute } from "~/pages/settings"

export const routeMap = {
  auth: authPageRoute,
  verifyOtp: verifyOtpPageRoute,
  startInvitationPageRoute: startInvitationPageRoute,
  receivedInvitationsPageRoute: receivedInvitationsPageRoute,
  onboardingSettingName: onboardingSettingNamePageRoute,
  onboardingSettingWorkspace: onboardingSettingWorkspacePageRoute,
  home: homePageRoute,
  settings: settingPageRoute,
  members: membersPageRoute
} as const

export const unauthenticatedRoutes = [routeMap.auth.toString(), routeMap.verifyOtp.toString()]

export const authRoutes = [routeMap.home.toString(), routeMap.settings.toString(), routeMap.members.toString()]

export const unprotectedInitialPagePath = routeMap.auth
