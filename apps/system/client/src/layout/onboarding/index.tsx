import { useContext, useEffect, useRef } from "react"
import { useAuthState } from "react-firebase-hooks/auth"
import { Outlet, useNavigate } from "react-router-dom"
import { firebaseAuth } from "~/infrastructure/firebase"
import { ContainerContext } from "~/infrastructure/injector/context"
import { routeMap } from "~/infrastructure/route/path"

export const OnboardingLayout = () => {
  const { controller, store } = useContext(ContainerContext)
  const navigate = useNavigate()
  const me = store.me((state) => state.me)
  const meIsLoading = store.me((state) => state.isLoading)
  const meRef = useRef(me)
  const meIsLoadingRef = useRef(meIsLoading)
  const [_, loading] = useAuthState(firebaseAuth)

  useEffect(() => {
    store.me.subscribe((state) => {
      meRef.current = state.me
      meIsLoadingRef.current = state.isLoading
    })
    const unsubscribed = firebaseAuth.onAuthStateChanged(async (user) => {
      if (!loading) {
        if (user) {
          await controller.me.find()
          return
        }
        await controller.me.signOut()
        navigate(routeMap.auth)
      }
    })
    return () => unsubscribed()
  }, [loading, navigate])

  return <Outlet />
}
