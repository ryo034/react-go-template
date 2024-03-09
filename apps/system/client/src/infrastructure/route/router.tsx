import { RouterProvider, createBrowserRouter } from "react-router-dom"
import { MemberRoleList } from "~/domain"
import { routeMap } from "~/infrastructure/route/path"
import { AuthenticatedLayout, AuthenticationLayout } from "~/layout/authentication"
import { AuthorizationLayout } from "~/layout/authorization"
import { DashboardLayout } from "~/layout/dashboard"
import { SettingsLayout } from "~/layout/dashboard/setting"
import { InvitationLayout } from "~/layout/invitation"
import { LoadingLayout } from "~/layout/loading"
import { OnboardingLayout } from "~/layout/onboarding"
import { ThemeLayout } from "~/layout/theme"
import { TrackingLayout } from "~/layout/tracking"
import { AuthPage } from "~/pages/auth"
import { ForbiddenPage } from "~/pages/error/forbidden"
import { NotFoundPage } from "~/pages/error/notFound"
import { HomePage } from "~/pages/home"
import { StartInvitationPage } from "~/pages/invitation"
import { MembersPage } from "~/pages/members"
import { OnboardingSettingNamePage } from "~/pages/onboarding/name"
import { OnboardingSettingWorkspacePage } from "~/pages/onboarding/workspace"
import { VerifyOtpPage } from "~/pages/otp"
import { ReceivedInvitationsPage } from "~/pages/receivedInvitation"
import { SettingsAppearancePage } from "~/pages/settings/appearance"
import { SettingsProfilePage } from "~/pages/settings/profile"
import { SettingsWorkspaceAccountPage } from "~/pages/settings/workspace/account"
import { SettingsWorkspaceInvitationsPage } from "~/pages/settings/workspace/invitations"
import { SettingsWorkspaceMembersPage } from "~/pages/settings/workspace/members"
import { SettingsWorkspaceSettingPage } from "~/pages/settings/workspace/setting"

export const accountInitialPagePath = routeMap.home

const router = createBrowserRouter([
  {
    path: "/",
    element: <TrackingLayout />,
    children: [
      {
        element: <ThemeLayout />,
        children: [
          {
            element: <LoadingLayout />,
            children: [
              {
                element: <AuthenticatedLayout />,
                children: [
                  { path: routeMap.auth, element: <AuthPage /> },
                  { path: routeMap.verifyOtp, element: <VerifyOtpPage /> }
                ]
              },
              {
                element: <InvitationLayout />,
                children: [
                  { path: routeMap.startInvitation, element: <StartInvitationPage /> },
                  { path: routeMap.receivedInvitations, element: <ReceivedInvitationsPage /> }
                ]
              },
              {
                element: <OnboardingLayout />,
                children: [
                  { path: routeMap.onboardingSettingName, element: <OnboardingSettingNamePage /> },
                  { path: routeMap.onboardingSettingWorkspace, element: <OnboardingSettingWorkspacePage /> }
                ]
              },
              {
                element: <AuthenticationLayout />,
                children: [
                  {
                    element: <DashboardLayout />,
                    children: [
                      { path: routeMap.home, element: <HomePage /> },
                      { path: routeMap.members, element: <MembersPage /> },
                      {
                        element: <SettingsLayout />,
                        children: [
                          { path: routeMap.settingsProfile, element: <SettingsProfilePage /> },
                          { path: routeMap.settingsAppearance, element: <SettingsAppearancePage /> },
                          {
                            element: (
                              <AuthorizationLayout
                                roles={[MemberRoleList.Owner, MemberRoleList.Admin]}
                                fallback={<ForbiddenPage />}
                              />
                            ),
                            children: [
                              { path: routeMap.settingsWorkspaceAccount, element: <SettingsWorkspaceAccountPage /> },
                              {
                                path: routeMap.settingsWorkspaceInvitation,
                                element: <SettingsWorkspaceInvitationsPage />
                              },
                              { path: routeMap.settingsWorkspaceMembers, element: <SettingsWorkspaceMembersPage /> },
                              { path: routeMap.settingsWorkspaceSetting, element: <SettingsWorkspaceSettingPage /> }
                            ]
                          }
                        ]
                      }
                    ]
                  }
                ]
              },
              { path: "*", element: <NotFoundPage /> }
            ]
          }
        ]
      }
    ]
  }
])

export const Router = () => {
  return <RouterProvider router={router} />
}
