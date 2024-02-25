import { useContext, useLayoutEffect, useRef } from "react"
import { useAuthState } from "react-firebase-hooks/auth"
import { Outlet, useNavigate } from "react-router-dom"
import { Loading } from "~/components/loading/loading"
import { firebaseAuth } from "~/infrastructure/firebase"
import { ContainerContext } from "~/infrastructure/injector/context"
import { unprotectedInitialPagePath } from "~/infrastructure/route/path"

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
      if (!loading) {
        if (!user) {
          await controller.me.signOut()
          navigate(unprotectedInitialPagePath)
          return
        }

        if (!meIsLoadingRef.current && meRef.current !== null) {
          return
        }
        const res = await controller.me.find()
        if (!res) return
        if (res !== null) {
          await controller.me.signOut()
          navigate(unprotectedInitialPagePath)
        }
      }
    })
    return () => unsubscribed()
  }, [loading, navigate, me])

  // loading can not early return
  if (!loading) {
    if (meRef.current !== null) {
      return <Outlet />
    }
  }
  return <Loading />
}

export const AuthenticatedLayout = () => {
  const { store } = useContext(ContainerContext)
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

    // const unsubscribed = firebaseAuth.onAuthStateChanged(async (user) => {
    //   if (!loading) {
    //     if (!user) {
    //       return
    //     }
    //   }
    // })
    // return () => unsubscribed()
  }, [loading, me])

  return <Outlet />
}
