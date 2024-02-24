import { LogOut } from "lucide-react"
import { useContext, useEffect, useRef } from "react"
import { useAuthState } from "react-firebase-hooks/auth"
import { Outlet, useNavigate } from "react-router-dom"
import {
  Button,
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
  useToast
} from "shared-ui"
import { firebaseAuth } from "~/infrastructure/firebase"
import { ContainerContext } from "~/infrastructure/injector/context"
import { routeMap } from "~/infrastructure/route/path"

const InvitationPageHeader = () => {
  const { store, controller } = useContext(ContainerContext)
  const { toast } = useToast()
  const navigate = useNavigate()
  const me = store.me((state) => state.me)

  const onClickLogout = async () => {
    const err = await controller.me.signOut()
    if (err) {
      return
    }
    toast({ title: "ログアウトしました" })
    navigate(routeMap.auth)
  }
  if (me === null) return <></>

  return (
    <header className="fixed top-0 z-50 w-full flex flex-wrap justify-end p-6">
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button variant="ghost" className="h-14 flex flex-wrap text-right">
            <p className="w-full text-xs text-muted-foreground">Logged in as:</p>
            <p className="w-full font-bold" data-testid="loggedInBy">
              {me.self.email.value}
            </p>
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent className="w-56">
          <DropdownMenuLabel>{me.self.name === null ? "No name" : me.self.name?.value}</DropdownMenuLabel>
          <DropdownMenuSeparator />
          <DropdownMenuItem onClick={onClickLogout} data-testid="logoutButtonOnInvitationHeader">
            <LogOut className="mr-2 h-4 w-4" />
            <span>Log out</span>
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </header>
  )
}

export const InvitationLayout = () => {
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
        if (!user) {
          return
        }
        await controller.me.find()
      }
    })
    return () => unsubscribed()
  }, [loading, navigate])

  return (
    <>
      {!meIsLoading && me !== null && <InvitationPageHeader />}
      <Outlet />
    </>
  )
}
