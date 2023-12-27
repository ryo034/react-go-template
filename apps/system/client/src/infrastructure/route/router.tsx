import { RouterProvider, createBrowserRouter } from "react-router-dom"
import { routeMap } from "~/infrastructure/route/path"
import { AuthLayout } from "~/layout/auth"
import { DashboardLayout } from "~/layout/dashboard"
import { LoadingLayout } from "~/layout/loading"
import { ThemeLayout } from "~/layout/theme"
import { TrackingLayout } from "~/layout/tracking"
import { AccountPage } from "~/pages/account"
import { ConfirmEmailPage } from "~/pages/confirmEmail"
import { NotFoundPage } from "~/pages/error/notFound"
import { HomePage } from "~/pages/home"
import { LoginPage } from "~/pages/login"
import { SettingsPage } from "~/pages/settings"

export const accountInitialPagePath = routeMap.creatures

// const LazyLoadedPage = (
//   pageName: string,
// ): React.LazyExoticComponent<React.ComponentType> => {
//   return lazy(() => import(`@/pages/${pageName}`));
// };

// {
//   path: routeMap.login, lazy: async () => {
//     // @ts-ignore
//     return { Component: (await import("~/pages/login/index.tsx")).LoginPage }
//   }
// },

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
                element: <AuthLayout />,
                children: [
                  { path: routeMap.login, element: <LoginPage /> },
                  { path: routeMap.confirmEmail, element: <ConfirmEmailPage /> },
                  {
                    element: <DashboardLayout />,
                    children: [
                      { path: routeMap.home, element: <HomePage /> },
                      { path: routeMap.account, element: <AccountPage /> },
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
