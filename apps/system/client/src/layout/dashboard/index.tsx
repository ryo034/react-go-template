"use client"

import { useContext, useState } from "react"
import { Outlet } from "react-router-dom"
import {
  ResizableHandle,
  ResizablePanel,
  ResizablePanelGroup,
  ScrollArea,
  Separator,
  TooltipProvider,
  cn
} from "shared-ui"
import { Nav } from "~/components/sidebar/nav"
import { SidebarUserNav } from "~/components/sidebar/userNav"
import { WorkspaceSwitcher } from "~/components/sidebar/workspaceSwitcher"
import { ContainerContext } from "~/infrastructure/injector/context"

interface DashboardLayoutProps {
  defaultLayout?: number[] | undefined
  defaultCollapsed?: boolean
  navCollapsedSize?: number
}

export const DashboardLayout = ({
  defaultLayout = [265, 440, 655],
  defaultCollapsed = false,
  navCollapsedSize = 4
}: DashboardLayoutProps) => {
  const { store } = useContext(ContainerContext)
  const [isCollapsed, setIsCollapsed] = useState(defaultCollapsed)
  const me = store.me((state) => state.me)

  return (
    <TooltipProvider delayDuration={0}>
      <ResizablePanelGroup
        direction="horizontal"
        onLayout={(sizes: number[]) => {
          document.cookie = `react-resizable-panels:layout=${JSON.stringify(sizes)}`
        }}
        className="h-full max-h-full items-stretch"
      >
        <ResizablePanel
          defaultSize={defaultLayout[0]}
          collapsedSize={navCollapsedSize}
          collapsible={true}
          minSize={15}
          maxSize={20}
          onExpand={() => setIsCollapsed(false)}
          onCollapse={() => setIsCollapsed(true)}
          className={cn(isCollapsed && "min-w-[50px] transition-all duration-300 ease-in-out")}
        >
          <div className={cn("flex h-[52px] items-center justify-center", isCollapsed ? "h-[52px]" : "px-2")}>
            <WorkspaceSwitcher
              isCollapsed={isCollapsed}
              workspaces={me?.joinedWorkspaces}
              currentWorkspace={me?.workspace}
            />
          </div>
          <Separator />
          <Nav isCollapsed={isCollapsed} />
          <SidebarUserNav isCollapsed={isCollapsed} />
        </ResizablePanel>
        <ResizableHandle withHandle />
        <ResizablePanel defaultSize={defaultLayout[2]}>
          <ScrollArea className="h-screen">
            <div className="space-y-6 p-10 pb-16 block">
              <Outlet />
            </div>
          </ScrollArea>
        </ResizablePanel>
      </ResizablePanelGroup>
    </TooltipProvider>
  )
}
