import { useContext } from "react"
import {
  Avatar,
  AvatarFallback,
  AvatarImage,
  Card,
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuShortcut,
  DropdownMenuTrigger
} from "shared-ui"
import { ContainerContext } from "~/infrastructure/injector/context"

export const SidebarUserNav = () => {
  const { controller, store } = useContext(ContainerContext)
  const onClickLogoutButton = async () => {
    await controller.me.signOut()
  }

  const me = store.me((state) => state.me)

  if (me === null) {
    return null
  }

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <div className="w-full mt-auto p-2">
          <Card className="p-2 cursor-pointer" data-testid="userNavigationOnSidebar">
            <div className="flex items-center space-x-2">
              <Avatar className="h-8 w-8">
                <AvatarImage src="/avatars/01.png" alt="@shadcn" />
                <AvatarFallback>{me.member?.profile.displayName?.firstTwoCharacters}</AvatarFallback>
              </Avatar>
              <p className="text-sm text-gray-500 dark:text-gray-400">{me.member?.profile.displayName?.value}</p>
            </div>
          </Card>
        </div>
      </DropdownMenuTrigger>
      <DropdownMenuContent className="w-56" align="end" forceMount>
        <DropdownMenuLabel className="font-normal">
          <div className="flex flex-col space-y-1">
            <p className="text-sm font-medium leading-none">{me.self.name?.value}</p>
            <p className="text-xs leading-none text-muted-foreground">{me.self.email.value}</p>
          </div>
        </DropdownMenuLabel>
        <DropdownMenuSeparator />
        <DropdownMenuGroup>
          <DropdownMenuItem>
            Profile
            <DropdownMenuShortcut>⇧⌘P</DropdownMenuShortcut>
          </DropdownMenuItem>
          <DropdownMenuItem>
            Billing
            <DropdownMenuShortcut>⌘B</DropdownMenuShortcut>
          </DropdownMenuItem>
          <DropdownMenuItem>
            Settings
            <DropdownMenuShortcut>⌘S</DropdownMenuShortcut>
          </DropdownMenuItem>
        </DropdownMenuGroup>
        <DropdownMenuSeparator />
        <DropdownMenuItem onClick={onClickLogoutButton} data-testid="logoutOnSidebar">
          ログアウト
          <DropdownMenuShortcut>⇧⌘Q</DropdownMenuShortcut>
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}
