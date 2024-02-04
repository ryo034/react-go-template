import { useContext } from "react"
import { Card, useToast } from "shared-ui"
import { ContainerContext } from "~/infrastructure/injector/context"

export const homePageRoute = "/home"

export function HomePage() {
  const { controller, store } = useContext(ContainerContext)
  const me = store.me((state) => state.me)

  const { toast } = useToast()

  const onClick = async () => {
    const res = await controller.me.signOut()
    if (!res) {
      toast({ title: "ログアウトしました👋" })
      return
    }
  }

  // if (me === null || me.self === undefined || me.self.name === undefined) {
  //   return <></>
  // }

  return (
    <div className="flex justify-center items-center min-h-screen">
      <Card>🎉🎉🎉Home🎉🎉🎉</Card>
    </div>
  )
}
