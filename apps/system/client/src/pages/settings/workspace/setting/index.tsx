import { useContext, useMemo, useState } from "react"
import { Separator, useToast } from "shared-ui"
import { ContainerContext } from "~/infrastructure/injector/context"
import { SettingsWorkspaceSettingForm, type SettingsWorkspaceSettingFormValues } from "./form"

export const settingsWorkspaceSettingPageRoute = "/settings/workspace/setting"

export const SettingsWorkspaceSettingPage = () => {
  const { store, controller } = useContext(ContainerContext)
  const { toast } = useToast()
  const me = store.me((state) => state.me)
  const meIsLoading = store.me((state) => state.isLoading)
  const [isUpdating, setIsUpdating] = useState(false)

  if (!me || meIsLoading || !me.workspace) return <></>

  const defaultValues = useMemo(() => {
    return {
      name: me.workspace?.name.value || ""
    }
  }, [me])

  const onSubmit = async (d: SettingsWorkspaceSettingFormValues) => {
    if (!me.workspace) return

    setIsUpdating(true)
    const err = await controller.workspace.updateWorkspace({
      name: d.name,
      subdomain: me.workspace?.subdomain.value,
      workspaceId: me.workspace?.id
    })
    setIsUpdating(false)
    if (err) {
      toast({ title: "Failed to update workspace" })
      return
    }
    toast({ title: "Workspace updated" })
  }

  return (
    <div className="space-y-6">
      <div>
        <h3 className="text-lg font-medium">Workspace Setting</h3>
        <p className="text-sm text-muted-foreground">
          Update your account settings. Set your preferred language and timezone.
        </p>
      </div>
      <Separator />
      <SettingsWorkspaceSettingForm isUpdating={isUpdating} onSubmit={onSubmit} defaultValues={defaultValues} />
    </div>
  )
}
