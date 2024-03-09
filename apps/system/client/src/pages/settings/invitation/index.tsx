import { HTMLAttributes, useContext, useLayoutEffect, useRef, useState } from "react"
import { Button, Card, CardContent, CardDescription, CardHeader, CardTitle, Separator, useToast } from "shared-ui"
import { Invitation, Me } from "~/domain"
import { ContainerContext } from "~/infrastructure/injector/context"

export const settingsWorkspaceInvitationsPageRoute = "/settings/workspace/invitations"

export interface Props extends HTMLAttributes<HTMLDivElement> {
  isUpdating?: boolean
  me: Me
  onClickResend: (invitation: Invitation) => void
  onClickRevoke: (invitation: Invitation) => void
  invitation: Invitation
}

const InviteeListItem = ({ isUpdating, me, invitation, onClickResend, onClickRevoke, ...props }: Props) => {
  return (
    <div className="flex items-center justify-between space-x-4" {...props}>
      <div className="flex items-center justify-between space-x-4 w-full">
        <div className="">
          <p className="text-sm font-medium leading-none">{invitation.displayName?.value || "(unknown)"}</p>
          <p className="text-sm text-muted-foreground truncate">{invitation.inviteeEmail.value}</p>
          <p className="text-sm text-muted-foreground truncate">inviter: {invitation.inviter.user.name?.value}</p>
        </div>
        {me.self.id.value.asString === invitation.inviter.user.id.value.asString && (
          <div className="flex space-x-4">
            <Button variant="outline" onClick={() => onClickResend(invitation)} disabled={isUpdating}>
              再送
            </Button>
            <Button variant="outline" onClick={() => onClickRevoke(invitation)} disabled={isUpdating}>
              取消
            </Button>
          </div>
        )}
      </div>
    </div>
  )
}

export const SettingsWorkspaceInvitationsPage = () => {
  const { store, controller } = useContext(ContainerContext)
  const { toast } = useToast()
  const me = store.me((state) => state.me)
  const invitations = store.invitations((state) => state.invitations)
  const invitationsLoading = store.invitations((state) => state.invitationsIsLoading)
  const invitationsRef = useRef(invitations)
  const invitationsLoadingRef = useRef(invitationsLoading)
  // const { toast } = useToast()

  const [isUpdating, setIsUpdating] = useState(false)

  useLayoutEffect(() => {
    store.invitations.subscribe((state) => {
      invitationsRef.current = state.invitations
      invitationsLoadingRef.current = state.invitationsIsLoading
    })

    const findAllInvitations = async () => {
      const err = await controller.workspace.findAllInvitations()
      if (err) {
        toast({ title: "Failed to get invitations" })
        return
      }
    }
    findAllInvitations()
  }, [])

  const onClickResend = async (invitation: Invitation) => {
    setIsUpdating(true)
    const err = await controller.workspace.resendInvitation(invitation)
    setIsUpdating(false)
    if (err) {
      toast({ title: "Failed to resend invitation" })
      return
    }
    toast({ title: "Invitation resent" })
  }

  const onClickRevoke = async (invitation: Invitation) => {
    setIsUpdating(true)
    const err = await controller.workspace.revokeInvitation(invitation)
    setIsUpdating(false)
    if (err) {
      toast({ title: "Failed to revoke invitation" })
      return
    }
    toast({ title: "Invitation revoked" })
  }

  if (!me) return <></>

  return (
    <div className="space-y-6">
      <div>
        <h3 className="text-lg font-medium">Invitation</h3>
        <p className="text-sm text-muted-foreground">
          Update your account settings. Set your preferred language and timezone.
        </p>
      </div>
      <Separator />

      <Card className="p-0">
        <CardHeader>
          <CardTitle>Invitation</CardTitle>
          <CardDescription>招待一覧です。取り消された招待はここには表示されません。</CardDescription>
        </CardHeader>
        {invitationsLoadingRef.current && (
          <CardContent>
            <p>loading...</p>
          </CardContent>
        )}
        {!invitationsLoadingRef.current && invitations.isEmpty && (
          <CardContent>
            <p>有効な招待はありません</p>
          </CardContent>
        )}
        {!invitationsLoadingRef.current && invitations.isNotEmpty && (
          <CardContent className="grid gap-6">
            {invitations.values.map((invitation) => (
              <InviteeListItem
                isUpdating={isUpdating}
                me={me}
                key={invitation.id.value.asString}
                invitation={invitation}
                onClickResend={onClickResend}
                onClickRevoke={onClickRevoke}
              />
            ))}
          </CardContent>
        )}
      </Card>
    </div>
  )
}
