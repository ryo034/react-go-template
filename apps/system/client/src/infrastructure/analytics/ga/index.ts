import { useContext, useEffect } from "react"
import { ReactGAImplementation } from "react-ga4"
import { useLocation } from "react-router-dom"
import { ContainerContext } from "~/infrastructure/injector/context"

export class MyCustomGA extends ReactGAImplementation {}

export const usePageTracking = () => {
  const location = useLocation()
  const { driver } = useContext(ContainerContext)

  useEffect(() => {
    // TODO:ここの初期化は移す
    driver.ga.initialize()
    driver.ga.sendLocation(location.pathname + location.search)
  }, [location])
}
