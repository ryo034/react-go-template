import { AlertCircle } from "lucide-react"
import { useContext, useMemo, useState } from "react"
import {
  Alert,
  AlertDescription,
  AlertDialog,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogDestructiveAction,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
  AlertTitle,
  Button,
  Separator,
  useToast
} from "shared-ui"
import { ContainerContext } from "~/infrastructure/injector/context"
import { SettingsWorkspaceSettingForm, type SettingsWorkspaceSettingFormValues } from "./form"

export const settingsWorkspaceSettingPageRoute = "/settings/workspace/setting"

const LeaveWorkspaceAlertDialog = () => {
  const { controller } = useContext(ContainerContext)
  const { toast } = useToast()

  const onClickLeaveWorkspaceButton = async () => {
    const err = await controller.me.leaveWorkspace()
    if (err) {
      toast({ title: "Failed to leave workspace" })
      return
    }
    toast({ title: "Left workspace" })
  }

  return (
    <AlertDialog>
      <AlertDialogTrigger asChild>
        <Button variant="outline">ワークスペースから退出</Button>
      </AlertDialogTrigger>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>このワークスペースを退出してもよろしいですか？</AlertDialogTitle>
          <AlertDialogDescription>退出後、一度ログアウトされます。</AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel>キャンセル</AlertDialogCancel>
          <AlertDialogDestructiveAction onClick={onClickLeaveWorkspaceButton} data-testid="leaveWorkspaceExecButton">
            退出
          </AlertDialogDestructiveAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  )
}

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
      <Separator />
      <Alert variant="warning">
        <AlertCircle className="h-4 w-4" />
        <AlertTitle>注意</AlertTitle>
        <AlertDescription>
          オーナーはワークスペースに所属しているのがオーナー1人の場合のみ退出できます。オーナーを退出させる場合は、他のメンバーに権限を委譲後に退出するようにしてください。
          ワークスペースに再参加するには、メンバーに招待してもらう必要があります。
          所属していた際に作成されたデータは削除されません。
        </AlertDescription>
      </Alert>

      <LeaveWorkspaceAlertDialog />
    </div>
  )
}
