import { RouterProvider, createBrowserRouter } from "react-router-dom"
import { routeMap } from "~/infrastructure/route/path"
import { AuthLayout } from "~/layout/auth"
import { DashboardLayout } from "~/layout/dashboard"
import { LoadingLayout } from "~/layout/loading"
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
              { path: routeMap.startInvitationPageRoute, element: <StartInvitationPage /> },
              {
                element: <AuthLayout />,
                children: [
                  // invitation
                  // auth
                  { path: routeMap.auth, element: <AuthPage /> },
                  { path: routeMap.verifyOtp, element: <VerifyOtpPage /> },
                  { path: routeMap.onboardingSettingName, element: <OnboardingSettingNamePage /> },
                  { path: routeMap.onboardingSettingWorkspace, element: <OnboardingSettingWorkspacePage /> },
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
