import { HTMLAttributes } from "react"
import { ScrollArea } from "shared-ui"
import Logo from "~/assets/react.svg"
import { NavItem, SidebarListItem } from "~/components/sidebar/listItem"

export interface SidebarProps extends HTMLAttributes<HTMLDivElement> {
  menus: NavItem[]
}

export const Sidebar = ({ menus }: SidebarProps) => {
  return (
    <aside className="hidden md:sticky md:block h-full min-h-full w-64 flex-col overflow-y-auto border-r border-slate-200 bg-sidebar py-4 dark:border-slate-700 dark:bg-slate-900">
      <ScrollArea className="flex h-full flex-col border-slate-200 dark:border-slate-700 dark:bg-slate-900">
        <div className="mb-10 flex items-center rounded-lg px-3 py-2 text-slate-900 dark:text-white">
          <img src={Logo} alt="Logo" />
          <span className="ml-3 text-lg font-semibold">Taxonomy</span>
        </div>
        <ul data-testid="sidebarMenuList" className="space-y-2 text-sm font-medium">
          {menus?.map((menu, i) => (
            <SidebarListItem key={`${menu.link}-${i}`} menu={menu} />
          ))}
        </ul>
      </ScrollArea>
    </aside>
  )
}
