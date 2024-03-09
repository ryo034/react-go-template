import { authPageRoute } from "~/pages/auth"
import { homePageRoute } from "~/pages/home"
import { startInvitationPageRoute } from "~/pages/invitation"
import { membersPageRoute } from "~/pages/members"
import { onboardingSettingNamePageRoute } from "~/pages/onboarding/name"
import { onboardingSettingWorkspacePageRoute } from "~/pages/onboarding/workspace"
import { verifyOtpPageRoute } from "~/pages/otp"
import { receivedInvitationsPageRoute } from "~/pages/receivedInvitation"
import { settingsWorkspaceAccountPageRoute } from "~/pages/settings/account"
import { settingsAppearancePageRoute } from "~/pages/settings/appearance"
import { settingsWorkspaceInvitationsPageRoute } from "~/pages/settings/invitation"
import { settingsWorkspaceMembersPageRoute } from "~/pages/settings/members"
import { settingsProfilePageRoute } from "~/pages/settings/profile"
import { settingsWorkspaceSettingPageRoute } from "~/pages/settings/workspace/setting"

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
  settingsProfile: settingsProfilePageRoute,
  settingsAppearance: settingsAppearancePageRoute,
  settingsWorkspaceAccount: settingsWorkspaceAccountPageRoute,
  settingsWorkspaceInvitation: settingsWorkspaceInvitationsPageRoute,
  settingsWorkspaceMembers: settingsWorkspaceMembersPageRoute,
  settingsWorkspaceSetting: settingsWorkspaceSettingPageRoute,
  members: membersPageRoute
} as const

export const isSettingsPage = (path: string) => {
  return (
    path.includes(routeMap.settingsAppearance) ||
    path.includes(routeMap.settingsProfile) ||
    path.includes(routeMap.settingsWorkspaceAccount) ||
    path.includes(routeMap.settingsWorkspaceInvitation) ||
    path.includes(routeMap.settingsWorkspaceMembers) ||
    path.includes(routeMap.settingsWorkspaceSetting)
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
