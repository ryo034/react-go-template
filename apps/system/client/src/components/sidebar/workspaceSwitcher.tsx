"use client"

import { CaretSortIcon } from "@radix-ui/react-icons"
import { CheckIcon } from "lucide-react"
import { ComponentPropsWithoutRef, useContext, useState } from "react"
import {
  Avatar,
  AvatarFallback,
  AvatarImage,
  Button,
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
  Dialog,
  Popover,
  PopoverContent,
  PopoverTrigger,
  cn
} from "shared-ui"
import { Workspace, Workspaces } from "~/domain"
import { ContainerContext } from "~/infrastructure/injector/context"

type PopoverTriggerProps = ComponentPropsWithoutRef<typeof PopoverTrigger>

type TeamSwitcherProps = PopoverTriggerProps & {
  isCollapsed: boolean
  workspaces?: Workspaces
  currentWorkspace?: Workspace
}

export const WorkspaceSwitcher = ({ className, isCollapsed, workspaces, currentWorkspace }: TeamSwitcherProps) => {
  const [open, setOpen] = useState(false)
  const [showNewWorkspaceDialog, setShowNewWorkspaceDialog] = useState(false)

  const { store } = useContext(ContainerContext)
  const me = store.me((state) => state.me)

  if (!currentWorkspace || !workspaces || !me) {
    return null
  }
  const [selectedTeam, setSelectedTeam] = useState<Workspace>(currentWorkspace)

  return (
    <Dialog open={showNewWorkspaceDialog} onOpenChange={setShowNewWorkspaceDialog}>
      <Popover open={open} onOpenChange={setOpen}>
        <PopoverTrigger asChild className="w-full">
          <Button
            fullWidth
            variant="outline"
            role="combobox"
            aria-expanded={open}
            aria-label="ワークスペースを選択"
            className={cn(`justify-between ${isCollapsed && "w-10"}`, className)}
          >
            <Avatar className="mr-2 h-5 w-5">
              <AvatarImage src={""} alt={selectedTeam.name.value} className="grayscale" />
              <AvatarFallback>{selectedTeam.subdomain.value.slice(0, 1).toUpperCase()}</AvatarFallback>
            </Avatar>
            {!isCollapsed && <p className="truncate">{selectedTeam.name.value}</p>}
            {!isCollapsed && <CaretSortIcon className="ml-auto h-4 w-4 shrink-0 opacity-50" />}
          </Button>
        </PopoverTrigger>
        <PopoverContent className="p-0 w-full">
          <Command>
            <CommandInput placeholder="Search workspace..." />
            <CommandList>
              <CommandEmpty>No workspace found.</CommandEmpty>
              {workspaces.values.map((w) => (
                <CommandGroup key={w.id.value.asString} heading={w.name.value}>
                  <CommandItem
                    key={w.id.value.asString}
                    onSelect={() => {
                      setSelectedTeam(w)
                      setOpen(false)
                    }}
                    className="text-sm"
                  >
                    <Avatar className="mr-2 h-5 w-5">
                      <AvatarImage src={""} alt={w.name.value} className="grayscale" />
                      <AvatarFallback>{w.subdomain.value.slice(0, 1).toUpperCase()}</AvatarFallback>
                    </Avatar>
                    {w.name.value}
                    <CheckIcon
                      className={cn("ml-auto h-4 w-4", selectedTeam.equals(w) ? "opacity-100" : "opacity-0")}
                    />
                  </CommandItem>
                </CommandGroup>
              ))}
            </CommandList>
            {/* <CommandSeparator />
            <CommandList>
              <CommandGroup>
                <DialogTrigger asChild>
                  <CommandItem
                    onSelect={() => {
                      setOpen(false)
                      setShowNewWorkspaceDialog(true)
                    }}
                  >
                    <PlusSquareIcon className="mr-2 h-5 w-5" />
                    ワークスペースを作成
                  </CommandItem>
                </DialogTrigger>
              </CommandGroup>
            </CommandList> */}
          </Command>
        </PopoverContent>
      </Popover>
      {/* <DialogContent>
        <DialogHeader>
          <DialogTitle>ワークスペースを新規作成</DialogTitle>
        </DialogHeader>
        <div>
          <div className="space-y-4 py-2 pb-4">
            <div className="space-y-2">
              <Label htmlFor="name">URLを入力</Label>
              <Input id="name" placeholder="Acme Inc." />
            </div>
          </div>
        </div>
        <DialogFooter>
          <Button variant="outline" onClick={() => setShowNewWorkspaceDialog(false)}>
            キャンセル
          </Button>
          <Button type="submit">作成</Button>
        </DialogFooter>
      </DialogContent> */}
    </Dialog>
  )
}
