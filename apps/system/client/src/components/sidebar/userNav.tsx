import { useContext } from "react"
import {
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
import { AccountAvatar } from "../account/avatar"

interface Props {
  isCollapsed: boolean
}

export const SidebarUserNav = ({ isCollapsed }: Props) => {
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
          {isCollapsed ? (
            <AccountAvatar
              alt="avatar"
              url={me.self.photo?.photoURL || ""}
              fallbackString={me.self.name?.firstTwoCharacters || ""}
              size="sm"
              data-testid="collapsedAvatarOnSidebar"
            />
          ) : (
            <Card className="p-2 cursor-pointer" date-testid="userNavigationOnSidebar">
              <div className="flex items-center space-x-2">
                <AccountAvatar
                  alt="avatar"
                  url={me.self.photo?.photoURL || ""}
                  fallbackString={me.self.name?.firstTwoCharacters || ""}
                  size="sm"
                  data-testid="avatarOnSidebar"
                />
                <p className="text-sm text-gray-500 truncate dark:text-gray-400" data-testid="displayNameOnSidebar">
                  {me.member?.profile.displayName?.value}
                </p>
              </div>
            </Card>
          )}
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
