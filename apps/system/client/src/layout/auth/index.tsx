import { useContext, useLayoutEffect, useRef } from "react"
import { useAuthState } from "react-firebase-hooks/auth"
import { Outlet, useNavigate } from "react-router-dom"
import { firebaseAuth } from "~/infrastructure/firebase"
import { ContainerContext } from "~/infrastructure/injector/context"
import { authRoutes, routeMap, unprotectedInitialPagePath } from "~/infrastructure/route/path"

export const AuthLayout = () => {
  const { controller, store } = useContext(ContainerContext)
  const navigate = useNavigate()
  const me = store.me((state) => state.me)
  const meRef = useRef(me)
  const [_, loading] = useAuthState(firebaseAuth)

  useLayoutEffect(() => {
    store.me.subscribe((state) => {
      meRef.current = state.me
    })

    const unsubscribed = firebaseAuth.onAuthStateChanged(async (user) => {
      if (loading) {
        return
      }
      const isAuthenticatedRoute = authRoutes.includes(window.location.pathname)

      if (!user && isAuthenticatedRoute) {
        navigate(unprotectedInitialPagePath)
      } else if (user && !isAuthenticatedRoute && meRef.current !== null) {
        if (meRef.current.self.hasNotName) {
          navigate(routeMap.onboardingSettingName)
          return
        }
        navigate(routeMap.home)
        return
      } else if (user && isAuthenticatedRoute && meRef.current !== null) {
        if (meRef.current.self.hasNotName) {
          navigate(routeMap.onboardingSettingName)
          return
        }
        if (meRef.current.hasNotWorkspace) {
          navigate(routeMap.onboardingSettingWorkspace)
          return
        }
        return
      }
      await controller.me.find()
    })
    return () => unsubscribed()
  }, [loading, navigate, location.pathname, me])

  if (loading) {
    return <div />
  }

  return <Outlet />
}
