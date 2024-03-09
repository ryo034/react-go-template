import { Outlet } from "react-router-dom"
import { Separator } from "shared-ui"
import { routeMap } from "~/infrastructure/route/path"
import { SidebarNav } from "../../../components/sidebar/settingsSidebarNav"

const sidebarNavAccountSettingItems = [
  {
    title: "Profile",
    href: routeMap.settingsProfile
  },
  {
    title: "Appearance",
    href: routeMap.settingsAppearance
  }
]

const sidebarNavWorkspaceSettingItems = [
  {
    title: "Setting",
    href: routeMap.settingsWorkspaceSetting
  },
  {
    title: "Account",
    href: routeMap.settingsWorkspaceAccount
  },
  {
    title: "Members",
    href: routeMap.settingsWorkspaceMembers
  },
  {
    title: "Invitations",
    href: routeMap.settingsWorkspaceInvitation
  }
]

interface Props {
  title: string
}

const SideNavSectionTitle = ({ title }: Props) => {
  return <h2 className="mb-2 px-4 text-lg font-semibold tracking-tight">{title}</h2>
}

export const SettingsLayout = () => {
  return (
    <>
      <h2 className="text-2xl font-bold tracking-tight" data-testid="pageTitle">
        Settings
      </h2>
      <Separator className="my-6" />
      <div className="flex flex-col space-y-8 lg:flex-row lg:space-x-12 lg:space-y-0">
        <aside className="-mx-4 lg:w-1/5 space-y-4">
          <div key="accountSettings">
            <SidebarNav items={sidebarNavAccountSettingItems} />
          </div>
          <div key="workspaceSettings">
            <SideNavSectionTitle title="Workspace" />
            <SidebarNav items={sidebarNavWorkspaceSettingItems} />
          </div>
        </aside>
        <div className="flex-1 lg:max-w-2xl">
          <Outlet />
        </div>
      </div>
    </>
  )
}
