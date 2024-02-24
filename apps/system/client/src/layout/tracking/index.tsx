import { useContext, useEffect } from "react"
import { Outlet, useLocation } from "react-router-dom"
import { ContainerContext } from "~/infrastructure/injector/context"

export const TrackingLayout = () => {
  const location = useLocation()
  const { driver } = useContext(ContainerContext)

  useEffect(() => {
    driver.logger.initialize()
    driver.logger.sendLocation(location.pathname + location.search)
  }, [location])

  return <Outlet />
}
