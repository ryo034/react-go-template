import { RouterProvider, createBrowserRouter } from "react-router-dom"
import { routeMap } from "~/infrastructure/route/path"
import { AuthLayout, AuthenticatedLayout } from "~/layout/auth"
import { DashboardLayout } from "~/layout/dashboard"
import { InvitationLayout } from "~/layout/invitation"
import { LoadingLayout } from "~/layout/loading"
import { OnboardingLayout } from "~/layout/onboarding"
import { ThemeLayout } from "~/layout/theme"
import { TrackingLayout } from "~/layout/tracking"
import { AuthPage } from "~/pages/auth"
import { NotFoundPage } from "~/pages/error/notFound"
import { HomePage } from "~/pages/home"
import { StartInvitationPage } from "~/pages/invitation"
import { MembersPage } from "~/pages/members"
import { OnboardingSettingNamePage } from "~/pages/onboarding/name"
import { OnboardingSettingWorkspacePage } from "~/pages/onboarding/workspace"
import { VerifyOtpPage } from "~/pages/otp"
import { ReceivedInvitationsPage } from "~/pages/receivedInvitation"
import { SettingsPage } from "~/pages/settings"

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
                  { path: routeMap.startInvitationPageRoute, element: <StartInvitationPage /> },
                  { path: routeMap.receivedInvitationsPageRoute, element: <ReceivedInvitationsPage /> }
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
                element: <AuthLayout />,
                children: [
                  {
                    element: <DashboardLayout />,
                    children: [
                      { path: routeMap.home, element: <HomePage /> },
                      { path: routeMap.members, element: <MembersPage /> },
                      { path: routeMap.settings, element: <SettingsPage /> }
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
