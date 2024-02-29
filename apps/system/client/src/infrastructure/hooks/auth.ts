import { useContext, useRef } from "react"
import { ContainerContext } from "../injector/context"
import { authRoutes, onboardingRoutes, routeMap } from "../route/path"

export const useAuthenticator = () => {
  const { store } = useContext(ContainerContext)
  const me = store.me((state) => state.me)
  const meIsLoading = store.me((state) => state.isLoading)
  const meRef = useRef(me)
  const meIsLoadingRef = useRef(meIsLoading)
  const isAuthenticatedRoute = authRoutes.includes(window.location.pathname)
  const isOnboardingRoutes = onboardingRoutes.includes(window.location.pathname)

  const nextNavigate = (pathname: string, search: string): { pathname: string; search: string } | null => {
    if (meRef.current === null || meIsLoadingRef.current) {
      return null
    }
    if (meRef.current.self.hasNotName) {
      return { pathname: routeMap.onboardingSettingName, search: "" }
    }
    if (meRef.current.hasReceivedInvitations) {
      return { pathname: routeMap.receivedInvitations, search: "" }
    }
    if (meRef.current.hasNotWorkspace) {
      return { pathname: routeMap.onboardingSettingWorkspace, search: "" }
    }
    if (meRef.current.doneOnboarding) {
      if (isOnboardingRoutes) {
        return { pathname: routeMap.home, search }
      }
      if (isAuthenticatedRoute) {
        return null
      }
    }
    return null
  }

  return {
    nextNavigate
  }
}
