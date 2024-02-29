import { authPageRoute } from "~/pages/auth"
import { homePageRoute } from "~/pages/home"
import { startInvitationPageRoute } from "~/pages/invitation"
import { membersPageRoute } from "~/pages/members"
import { onboardingSettingNamePageRoute } from "~/pages/onboarding/name"
import { onboardingSettingWorkspacePageRoute } from "~/pages/onboarding/workspace"
import { verifyOtpPageRoute } from "~/pages/otp"
import { receivedInvitationsPageRoute } from "~/pages/receivedInvitation"
import { settingsAccountPageRoute } from "~/pages/settings/account"
import { settingsAppearancePageRoute } from "~/pages/settings/appearance"
import { settingsProfilePageRoute } from "~/pages/settings/profile"

export const routeMap = {
  auth: authPageRoute,
  // onboarding
  verifyOtp: verifyOtpPageRoute,
  startInvitation: startInvitationPageRoute,
  receivedInvitations: receivedInvitationsPageRoute,
  onboardingSettingName: onboardingSettingNamePageRoute,
  onboardingSettingWorkspace: onboardingSettingWorkspacePageRoute,
  // dashboard
  home: homePageRoute,
  settingsAppearance: settingsAppearancePageRoute,
  settingsAccount: settingsAccountPageRoute,
  settingsProfile: settingsProfilePageRoute,
  members: membersPageRoute
} as const

export const isSettingsPage = (path: string) => {
  return (
    path.includes(routeMap.settingsAppearance) ||
    path.includes(routeMap.settingsProfile) ||
    path.includes(routeMap.settingsAccount)
  )
}

export const unauthenticatedRoutes = [routeMap.auth.toString(), routeMap.verifyOtp.toString()]
export const invitationRoutes = [routeMap.startInvitation.toString(), routeMap.receivedInvitations.toString()]
export const onboardingRoutes = [
  routeMap.onboardingSettingName.toString(),
  routeMap.onboardingSettingWorkspace.toString()
]

export const authRoutes = [
  routeMap.home.toString(),
  routeMap.members.toString(),
  routeMap.settingsAppearance.toString(),
  routeMap.settingsProfile.toString()
]

export const unprotectedInitialPagePath = routeMap.auth
