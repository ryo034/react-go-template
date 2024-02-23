import { Outlet } from "react-router-dom"
import { Separator } from "shared-ui"
import { routeMap } from "~/infrastructure/route/path"
import { SidebarNav } from "../../../components/sidebar/settingsSidebarNav"

const sidebarNavItems = [
  {
    title: "Profile",
    href: routeMap.settingsProfile
  },
  {
    title: "Account",
    href: routeMap.settingsAccount
  },
  {
    title: "Appearance",
    href: routeMap.settingsAppearance
  }
]

export const SettingsLayout = () => {
  return (
    <>
      <div className="hidden space-y-6 p-10 pb-16 md:block">
        <div className="space-y-0.5">
          <h2 className="text-2xl font-bold tracking-tight">Settings</h2>
          <p className="text-muted-foreground">Manage your account settings and set e-mail preferences.</p>
        </div>
        <Separator className="my-6" />
        <div className="flex flex-col space-y-8 lg:flex-row lg:space-x-12 lg:space-y-0">
          <aside className="-mx-4 lg:w-1/5 space-y-4">
            <div key="settingsAccount">
              <SidebarNav items={sidebarNavItems} />
            </div>
            {/* <div key="settingsWorkspace">
              <h2 className="mb-2 px-4 text-lg font-semibold tracking-tight">Workspace</h2>
              <SidebarNav items={sidebarNavItems} />
            </div> */}
          </aside>
          <div className="flex-1 lg:max-w-2xl">
            <Outlet />
          </div>
        </div>
      </div>
    </>
  )
}
