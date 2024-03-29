import { CheckIcon, ChevronDownIcon } from "lucide-react"
import { useContext, useLayoutEffect, useRef, useState } from "react"
import {
  Button,
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
  Command,
  CommandGroup,
  CommandItem,
  CommandList,
  CommandShortcut,
  Popover,
  PopoverContent,
  PopoverTrigger,
  Separator,
  useToast
} from "shared-ui"
import { AccountAvatar } from "~/components/account/avatar"
import { type Member, type SelectableRole, SelectableRoleList } from "~/domain"
import { useRole } from "~/infrastructure/hooks/role"
import { ContainerContext } from "~/infrastructure/injector/context"

export const settingsWorkspaceMembersPageRoute = "/settings/workspace/members"

export const SettingsWorkspaceMembersPage = () => {
  const { store, controller } = useContext(ContainerContext)
  const { toast } = useToast()
  const { translateRole } = useRole()
  const me = store.me((state) => state.me)
  const members = store.workspace((s) => s.members)
  const membersIsLoading = store.workspace((s) => s.membersIsLoading)
  const membersRef = useRef(members)

  const [isUpdating, setIsUpdating] = useState(false)

  useLayoutEffect(() => {
    store.workspace.subscribe((v) => {
      membersRef.current = v.members
    })

    const fetchMembers = async () => {
      if (me === null || me.workspace === undefined) {
        return null
      }
      await controller.workspace.findAllMembers()
    }
    fetchMembers()
  }, [])

  if (!me) return <></>

  const onSelectRole = async (member: Member, role: SelectableRole) => {
    setIsUpdating(true)
    const err = await controller.workspace.updateMemberRole(member.id, role)
    setIsUpdating(false)
    if (err) {
      toast({ title: "Failed to update role", variant: "destructive" })
      return
    }
  }

  return (
    <div className="space-y-6">
      <div>
        <h3 className="text-lg font-medium">Members</h3>
        <p className="text-sm text-muted-foreground">
          Invite members to your workspace. Manage their roles and permissions.
        </p>
      </div>
      <Separator />
      <Card className="p-0">
        <CardHeader>
          <CardTitle>Team Members</CardTitle>
          <CardDescription>Invite your team members to collaborate.</CardDescription>
        </CardHeader>
        <CardContent className="grid gap-6" data-testid="settingMembers">
          {membersIsLoading && <p>loading...</p>}
          {!membersIsLoading &&
            membersRef.current.values.map((m) => (
              <div
                className="flex items-center justify-between space-x-4"
                key={m.id.value.asString}
                data-testid={`settingMember-${m.user.email.value}`}
              >
                <div className="flex items-center space-x-4">
                  <AccountAvatar
                    alt={`${m.user.name?.value} avatar`}
                    url={m.user.photo?.photoURL ?? ""}
                    fallbackString={m.user.name?.firstTwoCharacters ?? ""}
                    size="sm"
                  />
                  <div className="max-w-[360px] w-full break-all">
                    <p className="text-sm font-medium leading-none">{m.user.name?.value}</p>
                    <p className="text-sm text-muted-foreground w-[360px]">{m.user.email.value}</p>
                  </div>
                </div>
                {m.isOwner || !me.member?.canEditRole ? (
                  <Button variant="outline" className="ml-auto" disabled data-testid="notSelectableRoleButton">
                    {translateRole(m.role)}
                    <ChevronDownIcon className="ml-2 h-4 w-4 text-muted-foreground" />
                  </Button>
                ) : (
                  <Popover>
                    <PopoverTrigger asChild>
                      <Button variant="outline" className="ml-auto" disabled={isUpdating}>
                        {translateRole(m.role)}
                        <ChevronDownIcon className="ml-2 h-4 w-4 text-muted-foreground" />
                      </Button>
                    </PopoverTrigger>
                    <PopoverContent className="p-0" align="end">
                      <Command>
                        <CommandList data-testid="selectMemberRole">
                          <CommandGroup>
                            {SelectableRoleList.map((role) => {
                              return (
                                <CommandItem
                                  className="px-4 py-2 cursor-pointer"
                                  onSelect={(v) => onSelectRole(m, role)}
                                  key={`${m.id.value.asString}-selectRole-${role}`}
                                >
                                  <span className="pr-2">{translateRole(role)}</span>
                                  {role !== m.role ? null : (
                                    <CommandShortcut>
                                      <CheckIcon size={16} />
                                    </CommandShortcut>
                                  )}
                                </CommandItem>
                              )
                            })}
                          </CommandGroup>
                        </CommandList>
                      </Command>
                    </PopoverContent>
                  </Popover>
                )}
              </div>
            ))}
        </CardContent>
      </Card>
    </div>
  )
}
