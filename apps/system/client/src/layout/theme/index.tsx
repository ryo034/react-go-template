import { useContext, useLayoutEffect } from "react"
import { Outlet } from "react-router-dom"
import { ContainerContext } from "~/infrastructure/injector/context"

export const ThemeLayout = () => {
  const { controller } = useContext(ContainerContext)

  useLayoutEffect(() => {
    controller.theme.init()
  }, [])

  return <Outlet />
}
