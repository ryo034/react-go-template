import { useContext, useEffect } from "react"
import { useNavigate } from "react-router-dom"
import { Button, Card, Separator, useToast } from "shared-ui"
import { InvitationId } from "~/domain"
import { ContainerContext } from "~/infrastructure/injector/context"
import { routeMap } from "~/infrastructure/route/path"

export const receivedInvitationsPageRoute = "/received-invitations"

export const ReceivedInvitationsPage = () => {
  const { store, controller, errorMessageProvider } = useContext(ContainerContext)
  const me = store.me((state) => state.me)
  const { toast } = useToast()
  const meIsLoading = store.me((state) => state.isLoading)
  const navigate = useNavigate()

  // if not receivedInvitations, navigate to home
  useEffect(() => {
    if (!meIsLoading && me?.hasNotReceivedInvitations === true) {
      navigate(routeMap.home)
      return
    }
  }, [me, meIsLoading])

  if (meIsLoading || me === null) {
    // TODO: Loading画面
    return null
  }

  const onClickJoinButton = async (invitationId: InvitationId) => {
    const err = await controller.me.acceptInvitation({ invitationId })
    if (err !== null) {
      toast({ title: errorMessageProvider.resolve(err), variant: "destructive" })
      return
    }
    navigate(routeMap.home)
  }

  return (
    <div className="flex justify-center items-center min-h-screen">
      <div className="mx-auto flex flex-wrap justify-center">
        <div className="space-y-2 text-center mb-12 px-8 w-full">
          <p className="text-4xl font-bold">You have access to these workspaces</p>
        </div>
        <Card className="w-[480px]">
          {me.receivedInvitations.values.map((invitation) => {
            return (
              <div key={invitation.invitation.id.value.asString}>
                <div className="flex justify-between">
                  <div className="flex space-x-2 items-center">
                    <div className="flex flex-col">
                      <p>{invitation.inviter.workspace.name.value}</p>
                      <span className="text-xs text-muted-foreground">
                        inviter: {invitation.inviter.self.profile.displayName?.value}
                      </span>
                    </div>
                  </div>
                  <Button onClick={() => onClickJoinButton(invitation.invitation.id)}>join</Button>
                </div>
                {me.receivedInvitations.values.length > 1 && <Separator />}
              </div>
            )
          })}
        </Card>
      </div>
    </div>
  )
}
