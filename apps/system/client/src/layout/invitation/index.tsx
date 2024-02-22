import { useContext, useEffect, useRef } from "react"
import { Outlet } from "react-router-dom"
import { Button } from "shared-ui"
import { firebaseAuth } from "~/infrastructure/firebase"
import { ContainerContext } from "~/infrastructure/injector/context"

const InvitationPageHeader = () => {
  const { store } = useContext(ContainerContext)

  const me = store.me((state) => state.me)
  if (me === null) return null

  return (
    <header className="fixed top-0 z-50 w-full flex flex-wrap justify-end p-6">
      <Button variant="ghost" className="h-14 flex flex-wrap text-right">
        <p className="w-full text-xs text-muted-foreground">Logged in as:</p>
        <p className="w-full font-bold">{me.self.email.value}</p>
      </Button>
    </header>
  )
}

export const InvitationLayout = () => {
  const { controller, store } = useContext(ContainerContext)
  const me = store.me((state) => state.me)
  const meIsLoading = store.me((state) => state.isLoading)
  const meRef = useRef(me)
  const meIsLoadingRef = useRef(meIsLoading)

  useEffect(() => {
    store.me.subscribe((state) => {
      meRef.current = state.me
      meIsLoadingRef.current = state.isLoading
    })
    const unsubscribed = firebaseAuth.onAuthStateChanged(async (user) => await controller.me.find())
    return () => unsubscribed()
  }, [])

  return (
    <>
      {!meIsLoading && me !== null && <InvitationPageHeader />}
      <Outlet />
    </>
  )
}
