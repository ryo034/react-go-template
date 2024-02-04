import { LucideBarChart4, LucideCog, LucideDollarSign, LucideFish, Package } from "lucide-react"
import { useContext, useLayoutEffect } from "react"
import { Outlet } from "react-router-dom"
import { NavItem } from "~/components/sidebar/listItem"
import { Sidebar } from "~/components/sidebar/sidebar"
import { ContainerContext } from "~/infrastructure/injector/context"
import { routeMap } from "~/infrastructure/route/path"

export const sideMenu: NavItem[] = [
  {
    label: "ホーム",
    icon: <LucideBarChart4 className="w-6 h-6" />,
    link: routeMap.home
  },
  {
    label: "設定",
    icon: <LucideCog className="w-6 h-6" />,
    link: routeMap.settings
  }
]

export const DashboardLayout = () => {
  const { store } = useContext(ContainerContext)
  useLayoutEffect(() => {
    store.me.subscribe((state) => {
      if (state.me === null) {
        return
      }
      return () => {}
    })
  })
  return (
    <div className="overflow-hidden w-full h-full relative flex z-0">
      <Sidebar menus={sideMenu} />
      <main className="relative flex h-full max-w-full flex-1 overflow-hidden">
        <div className="flex-1 overflow-y-auto">
          <Outlet />
        </div>
      </main>
    </div>
  )
}
