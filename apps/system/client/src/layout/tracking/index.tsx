import { Outlet } from "react-router-dom"
import { usePageTracking } from "~/infrastructure/analytics/ga"

export const TrackingLayout = () => {
  usePageTracking()
  return <Outlet />
}
