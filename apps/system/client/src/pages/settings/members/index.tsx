import { ChevronDownIcon } from "lucide-react"
import { useContext, useLayoutEffect, useRef, useState } from "react"
import {
  Button,
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
  Command,
  CommandEmpty,
  CommandGroup,
  CommandItem,
  CommandList,
  Popover,
  PopoverContent,
  PopoverTrigger,
  Separator,
  useToast
} from "shared-ui"
import { AccountAvatar } from "~/components/account/avatar"
import { ContainerContext } from "~/infrastructure/injector/context"

export const settingsMembersPageRoute = "/settings/members"

export const SettingsMembersPage = () => {
  const { store, controller } = useContext(ContainerContext)
  const { toast } = useToast()
  const me = store.me((state) => state.me)
  const members = store.workspace((s) => s.members)
  const membersIsLoading = store.workspace((s) => s.membersIsLoading)
  const membersRef = useRef(members)
  // const { toast } = useToast()

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
        <CardContent className="grid gap-6">
          {membersIsLoading && <p>loading...</p>}
          {!membersIsLoading &&
            membersRef.current.values.map((m) => (
              <div className="flex items-center justify-between space-x-4" key={m.id.value.asString}>
                <div className="flex items-center space-x-4">
                  <AccountAvatar
                    alt={`${m.user.name?.value} avatar`}
                    url={m.user.photo?.photoURL ?? ""}
                    fallbackString={m.user.name?.firstTwoCharacters ?? ""}
                    size="sm"
                  />
                  <div className="max-w-[420px] w-full break-all">
                    <p className="text-sm font-medium leading-none">{m.user.name?.value}</p>
                    <p className="text-sm text-muted-foreground w-[420px]">{m.user.email.value}</p>
                  </div>
                </div>
                {m.isOwner ? (
                  <Button variant="outline" className="ml-auto" disabled>
                    Owner
                  </Button>
                ) : (
                  <Popover>
                    <PopoverTrigger asChild>
                      <Button variant="outline" className="ml-auto">
                        {m.role} <ChevronDownIcon className="ml-2 h-4 w-4 text-muted-foreground" />
                      </Button>
                    </PopoverTrigger>
                    <PopoverContent className="p-0" align="end">
                      <Command>
                        <CommandList>
                          <CommandEmpty>No roles found.</CommandEmpty>
                          <CommandGroup>
                            <CommandItem className="teamaspace-y-1 flex flex-col items-start px-4 py-2">
                              <p>Guest</p>
                              <p className="text-sm text-muted-foreground">Can view and comment.</p>
                            </CommandItem>
                            <CommandItem className="teamaspace-y-1 flex flex-col items-start px-4 py-2">
                              <p>Member</p>
                              <p className="text-sm text-muted-foreground">Can view, comment and edit.</p>
                            </CommandItem>
                            <CommandItem className="teamaspace-y-1 flex flex-col items-start px-4 py-2">
                              <p>Admin</p>
                              <p className="text-sm text-muted-foreground">Can view, comment and manage billing.</p>
                            </CommandItem>
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
