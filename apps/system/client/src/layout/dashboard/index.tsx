"use client"

import * as React from "react"
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
import { AccountSwitcher } from "~/components/sidebar/accountSwitcher"
import { Nav } from "~/components/sidebar/nav"

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
  const [isCollapsed, setIsCollapsed] = React.useState(defaultCollapsed)

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
          onExpand={() => {
            setIsCollapsed(false)
          }}
          onCollapse={() => {
            setIsCollapsed(true)
          }}
          className={cn(isCollapsed && "min-w-[50px] transition-all duration-300 ease-in-out")}
        >
          <div className={cn("flex h-[52px] items-center justify-center", isCollapsed ? "h-[52px]" : "px-2")}>
            {/* <AccountSwitcher isCollapsed={isCollapsed} accounts={[]} /> */}
          </div>
          <Separator />
          <Nav isCollapsed={isCollapsed} />
        </ResizablePanel>
        <ResizableHandle withHandle />
        <ResizablePanel defaultSize={defaultLayout[2]}>
          <ScrollArea className="h-screen">
            <Outlet />
          </ScrollArea>
        </ResizablePanel>
      </ResizablePanelGroup>
    </TooltipProvider>
  )
}
