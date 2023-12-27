import { useContext, useLayoutEffect, useRef } from "react"
import { useAuthState } from "react-firebase-hooks/auth"
import { Outlet, useNavigate } from "react-router-dom"
import { firebaseAuth } from "~/infrastructure/firebase"
import { ContainerContext } from "~/infrastructure/injector/context"
import { authRoutes, routeMap, unprotectedInitialPagePath } from "~/infrastructure/route/path"

export const AuthLayout = () => {
  const { controller, store, driver } = useContext(ContainerContext)
  const navigate = useNavigate()
  const me = store.me((state) => state.me)
  const meRef = useRef(me)
  const [_, loading] = useAuthState(firebaseAuth)

  useLayoutEffect(() => {
    // 認証済みルートを検証
    store.me.subscribe((state) => {
      const isAuthenticatedRoute = authRoutes.includes(window.location.pathname)
      meRef.current = state.me
      if (state.me === null || driver.firebase.currentUser === null) {
        return
      }
      if (isAuthenticatedRoute && state.me.emailNotVerified) {
        navigate(routeMap.confirmEmail)
      }
    })

    // 未認証ルートを検証
    const unsubscribed = firebaseAuth.onAuthStateChanged(async (user) => {
      const isAuthenticatedRoute = authRoutes.includes(window.location.pathname)
      if (user === null) {
        if (isAuthenticatedRoute) {
          navigate(unprotectedInitialPagePath)
          return
        }
      }
      await controller.me.find()
    })

    return () => unsubscribed()
  }, [])

  if (loading) {
    return <div />
  }

  return <Outlet />
}
