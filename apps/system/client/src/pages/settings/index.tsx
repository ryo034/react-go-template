import { LogOut } from "lucide-react"
import { useContext } from "react"
import { Button, Card, useToast } from "shared-ui"
import { ContainerContext } from "~/infrastructure/injector/context"

export const settingPageRoute = "/settings"

export const SettingsPage = () => {
  const { controller } = useContext(ContainerContext)
  // const me = store.me((state) => state.me)

  const { toast } = useToast()

  const onClick = async () => {
    const res = await controller.me.signOut()
    if (!res) {
      toast({ title: "ログアウトしました👋" })
      return
    }
  }

  // if (me === null) {
  //   return <></>
  // }

  return (
    <div className="flex justify-center items-center min-h-screen">
      <Card>
        <Button className="w-full" onClick={onClick}>
          <LogOut className="mr-2 h-4 w-4" /> Logout
        </Button>
      </Card>
    </div>
  )
}
