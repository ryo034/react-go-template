import { useContext, useEffect } from "react"
import { useAuthState } from "react-firebase-hooks/auth"
import { useNavigate } from "react-router-dom"
import { Button, Card, Separator, useToast } from "shared-ui"
import type { InvitationId } from "~/domain"
import { firebaseAuth } from "~/infrastructure/firebase"
import { useErrorHandler, useErrorMessageHandler } from "~/infrastructure/hooks/error"
import { ContainerContext } from "~/infrastructure/injector/context"
import { routeMap } from "~/infrastructure/route/path"
import { useReceivedInvitationsPageMessage } from "./message"

export const receivedInvitationsPageRoute = "/received-invitations"

export const ReceivedInvitationsPage = () => {
  const { store, controller, driver } = useContext(ContainerContext)
  const me = store.me((state) => state.me)
  const { toast } = useToast()
  const { handleError } = useErrorHandler()
  const { handleErrorMessage } = useErrorMessageHandler()
  const [_, loading] = useAuthState(driver.firebase.getClient)

  const meIsLoading = store.me((state) => state.isLoading)
  const navigate = useNavigate()

  const message = useReceivedInvitationsPageMessage()

  // if not receivedInvitations, navigate to home
  useEffect(() => {
    if (!meIsLoading && me?.hasNotReceivedInvitations === true) {
      navigate(routeMap.home)
      return
    }
  }, [me, meIsLoading])

  useEffect(() => {
    const unsubscribed = firebaseAuth.onAuthStateChanged(async (user) => {
      if (loading) {
        return
      }
      if (!user) {
        navigate(routeMap.auth)
        return
      }
      await controller.me.find()
    })
    return () => unsubscribed()
  }, [loading, navigate])

  const onClickJoinButton = async (invitationId: InvitationId) => {
    const err = await controller.me.acceptInvitation({ invitationId })
    if (err !== null) {
      toast({ title: handleErrorMessage(err), variant: "destructive" })
      handleError(err)
      return
    }
    navigate(routeMap.home)
  }

  if (meIsLoading || me === null) {
    return <></>
  }

  return (
    <div className="flex justify-center items-center min-h-screen" data-testid="receivedInvitationsPage">
      <div className="mx-auto flex flex-wrap justify-center">
        <div className="space-y-2 text-center mb-12 px-8 w-full">
          <p className="text-4xl font-bold">{message.title}</p>
        </div>
        <Card className="w-[480px]" data-testid="receivedInvitations">
          {me.receivedInvitations.values.map((invitation) => {
            return (
              <div key={invitation.invitation.id.value.asString} data-testid="invitation">
                <div className="flex justify-between">
                  <div className="flex space-x-2 items-center">
                    <div className="flex flex-col">
                      <p data-testid="workspaceName">{invitation.inviter.workspace.name.value}</p>
                      <span data-testid="inviterName" className="text-xs text-muted-foreground">
                        {message.word.inviter}: {invitation.inviter.self.profile.displayName?.value}
                      </span>
                    </div>
                  </div>
                  <Button onClick={() => onClickJoinButton(invitation.invitation.id)} data-testid="joinButton">
                    {message.action.join}
                  </Button>
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
