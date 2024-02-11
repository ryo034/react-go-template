import { useContext, useLayoutEffect, useRef } from "react"
import { useAuthState } from "react-firebase-hooks/auth"
import { Outlet, useNavigate } from "react-router-dom"
import { Loading } from "~/components/loading/loading"
import { firebaseAuth } from "~/infrastructure/firebase"
import { ContainerContext } from "~/infrastructure/injector/context"
import { authRoutes, routeMap, unprotectedInitialPagePath } from "~/infrastructure/route/path"

export const AuthLayout = () => {
  const { controller, store } = useContext(ContainerContext)
  const navigate = useNavigate()
  const me = store.me((state) => state.me)
  const meIsLoading = store.me((state) => state.isLoading)
  const meRef = useRef(me)
  const meIsLoadingRef = useRef(meIsLoading)
  const [_, loading] = useAuthState(firebaseAuth)

  useLayoutEffect(() => {
    store.me.subscribe((state) => {
      meRef.current = state.me
      meIsLoadingRef.current = state.isLoading
    })

    const unsubscribed = firebaseAuth.onAuthStateChanged(async (user) => {
      if (loading) {
        return
      }
      const isAuthenticatedRoute = authRoutes.includes(window.location.pathname)
      if (!user) {
        if (isAuthenticatedRoute) {
          navigate(unprotectedInitialPagePath)
          return
        }
        return
      }

      if (meRef.current !== null) {
        if (meRef.current.self.hasNotName) {
          navigate(routeMap.onboardingSettingName)
          return
        }
        if (meRef.current.hasNotWorkspace) {
          navigate(routeMap.onboardingSettingWorkspace)
          return
        }
        if (meRef.current.doneOnboarding) {
          if (isAuthenticatedRoute) {
            navigate(window.location.pathname)
            return
          }
          navigate(routeMap.home)
        }
        return
      }

      const res = await controller.me.find()
      if (!res) return
      if (res !== null) {
        await controller.me.signOut()
        navigate(unprotectedInitialPagePath)
      }
    })
    return () => unsubscribed()
  }, [loading, navigate, me])

  if (loading || meIsLoadingRef.current) {
    return <Loading />
  }

  return <Outlet />
}
