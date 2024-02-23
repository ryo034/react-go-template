"use client"

import { Settings, Users2, UsersRound } from "lucide-react"
import { useContext } from "react"
import { useLocation } from "react-router-dom"
import { Tooltip, TooltipContent, TooltipTrigger } from "shared-ui"
import { ContainerContext } from "~/infrastructure/injector/context"
import { isSettingsPage, routeMap } from "~/infrastructure/route/path"
import { NavItem, SidebarListItem, SidebarListItemCollapsed } from "./listItem"

interface NavProps {
  isCollapsed: boolean
}

export const Nav = ({ isCollapsed }: NavProps) => {
  const { i18n } = useContext(ContainerContext)
  const location = useLocation()
  const links: NavItem[] = [
    {
      title: i18n.translate("router.home"),
      icon: Users2,
      variant: location.pathname === routeMap.home ? "default" : "ghost",
      to: routeMap.home
    },
    {
      title: i18n.translate("router.member"),
      icon: UsersRound,
      variant: location.pathname === routeMap.members ? "default" : "ghost",
      to: routeMap.members
    },
    {
      title: i18n.translate("router.setting"),
      icon: Settings,
      variant: isSettingsPage(location.pathname) ? "default" : "ghost",
      to: routeMap.settingsProfile
    }
  ]

  return (
    <div data-collapsed={isCollapsed} className="w-full group flex flex-col gap-4 py-2 data-[collapsed=true]:py-2">
      <nav className="grid gap-1 px-2 group-[[data-collapsed=true]]:justify-center group-[[data-collapsed=true]]:px-2">
        {links.map((link) =>
          isCollapsed ? (
            <Tooltip delayDuration={0}>
              <TooltipTrigger asChild>
                <SidebarListItemCollapsed {...link} />
              </TooltipTrigger>
              <TooltipContent side="right" className="flex items-center gap-4">
                {link.title}
                {link.label && <span className="ml-auto text-muted-foreground">{link.label}</span>}
              </TooltipContent>
            </Tooltip>
          ) : (
            <SidebarListItem {...link} key={`${link.title}-${link.to}`} />
          )
        )}
      </nav>
    </div>
  )
}
