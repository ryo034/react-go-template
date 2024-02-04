import { useContext } from "react"
import { Button } from "shared-ui"
import { ContainerContext } from "~/infrastructure/injector/context"

export const settingPageRoute = "/settings"

export const SettingsPage = () => {
  const { store, controller } = useContext(ContainerContext)
  const me = store.me((state) => state.me)

  if (me === null) {
    return <></>
  }

  const onClickLogout = async () => {
    await controller.me.signOut()
  }

  return (
    <div className="flex justify-center items-center min-h-screen">
      <Button onClick={onClickLogout}>ログアウト</Button>
    </div>
  )
}
