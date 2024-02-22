import { useContext, useLayoutEffect, useRef, useState } from "react"
import { Trans } from "react-i18next"
import { createSearchParams, useNavigate, useSearchParams } from "react-router-dom"
import { isBadRequestError } from "shared-network"
import { Button, LoadingButton, useToast } from "shared-ui"
import { ReceivedInvitation } from "~/domain"
import {
  isAlreadyAcceptedInvitationError,
  isAlreadyExpiredInvitationError,
  isAlreadyRevokeInvitationError
} from "~/domain/workspace/invitation/error"
import { firebaseAuth } from "~/infrastructure/firebase"
import { ContainerContext } from "~/infrastructure/injector/context"
import { routeMap } from "~/infrastructure/route/path"
import { useStartInvitationPageMessage } from "./message"

export const startInvitationPageRoute = "/invitation"

const errorMessageTypeList = {
  unknown: "unknown",
  invalidToken: "invalidToken",
  alreadyExpired: "alreadyExpired",
  alreadyRevoked: "alreadyRevoked",
  alreadyAccepted: "alreadyAccepted"
} as const

type ErrorMessageType = (typeof errorMessageTypeList)[keyof typeof errorMessageTypeList]

type ErrorMessage = {
  title: string
  description: string
}

export const StartInvitationPage = () => {
  const { store, controller } = useContext(ContainerContext)
  const [searchParams] = useSearchParams()
  const [errorMessageType, setErrorMessageType] = useState<ErrorMessageType | null>(null)

  const message = useStartInvitationPageMessage()
  const meIsLoading = store.me((state) => state.isLoading)
  const receivedInvitation = store.receivedInvitation((s) => s.invitation)
  const me = store.me((s) => s.me)
  const meRef = useRef(me)
  const meIsLoadingRef = useRef(meIsLoading)

  useLayoutEffect(() => {
    store.me.subscribe((state) => {
      meRef.current = state.me
      meIsLoadingRef.current = state.isLoading
    })

    const unsubscribed = firebaseAuth.onAuthStateChanged(async () => {
      await controller.me.find()
    })
    return () => unsubscribed()
  }, [])

  useLayoutEffect(() => {
    const fetchInvitationByToken = async () => {
      const token = searchParams.get("token") || ""
      const err = await controller.auth.findInvitationByToken(token)
      if (err === null) {
        return
      }

      if (isBadRequestError(err)) {
        setErrorMessageType("invalidToken")
        return
      }
      if (isAlreadyExpiredInvitationError(err)) {
        setErrorMessageType("alreadyExpired")
        return
      }
      if (isAlreadyRevokeInvitationError(err)) {
        setErrorMessageType("alreadyRevoked")
        return
      }
      if (isAlreadyAcceptedInvitationError(err)) {
        setErrorMessageType("alreadyAccepted")
        return
      }
    }
    fetchInvitationByToken()
  }, [])

  return (
    <div className="flex justify-center items-center min-h-screen">
      <InvitationSection
        titleI18nKey={message.title}
        receivedInvitation={receivedInvitation}
        description={message.description("uni.nashi.034+1@gmail.com")}
        startButtonLabel={message.action.start}
        errorMessageType={errorMessageType}
      />
    </div>
  )
}

const InvitationSection = ({
  titleI18nKey,
  description,
  receivedInvitation,
  startButtonLabel,
  errorMessageType
}: {
  titleI18nKey: string
  receivedInvitation: ReceivedInvitation | null
  description: string
  startButtonLabel: string
  errorMessageType: ErrorMessageType | null
}) => {
  const [searchParams] = useSearchParams()
  const { controller, store } = useContext(ContainerContext)
  const navigate = useNavigate()
  const { toast } = useToast()

  const isInvitationProcessing = store.receivedInvitation((s) => s.isInvitationProcessing)

  if (errorMessageType !== null) {
    return <ErrorMessageSection errorMessageType={errorMessageType} />
  }

  if (receivedInvitation === null) {
    return <div>Loading...</div>
  }

  const onClickStartButton = async () => {
    const token = searchParams.get("token") || ""
    const err = await controller.auth.proceedToInvitation(token, receivedInvitation.invitation.inviteeEmail)
    if (err !== null) {
      toast({
        title: "招待の受諾に失敗しました。お手数ですが、しばらくしてから再度お試しください",
        variant: "destructive"
      })
      return
    }
    navigate({
      pathname: routeMap.verifyOtp,
      search: createSearchParams({ email: receivedInvitation.invitation.inviteeEmail.value }).toString()
    })
  }

  return (
    <div className="mx-auto space-y-6 flex flex-wrap justify-center">
      <div className="space-y-2 text-center mb-12 px-8 w-full">
        <h1 className="text-4xl font-bold">
          <Trans
            i18nKey={titleI18nKey}
            values={{
              inviterName: receivedInvitation.inviter.self.profile.displayName?.value,
              workspaceName: receivedInvitation.inviter.workspace.name.value
            }}
          />
        </h1>
        <p className="text-muted-foreground">{description}</p>
      </div>
      {isInvitationProcessing ? (
        <LoadingButton className="w-[256px]" />
      ) : (
        <Button className="w-[256px]" onClick={onClickStartButton} data-testid="startButtonFromInvitation">
          {startButtonLabel}
        </Button>
      )}
    </div>
  )
}

const ErrorMessageSection = ({ errorMessageType }: { errorMessageType: ErrorMessageType | null }) => {
  const message = useStartInvitationPageMessage()
  const navigate = useNavigate()

  const onClickGoBackButton = () => {
    navigate("/")
  }

  if (errorMessageType === null) return null
  let errorMessage: ErrorMessage | null = null
  switch (errorMessageType) {
    case "invalidToken":
      errorMessage = {
        title: message.error.invitationInvalidTokenTitle,
        description: message.error.invitationInvalidTokenDescription
      }
      break
    case "alreadyExpired":
      errorMessage = {
        title: message.error.invitationAlreadyExpiredTitle,
        description: message.error.invitationAlreadyExpiredDescription
      }
      break
    case "alreadyRevoked":
      errorMessage = {
        title: message.error.invitationAlreadyRevokedTitle,
        description: message.error.invitationAlreadyRevokedDescription
      }
      break
    case "alreadyAccepted":
      errorMessage = {
        title: message.error.invitationAlreadyAcceptedTitle,
        description: message.error.invitationAlreadyAcceptedDescription
      }
      break
    case "unknown":
      errorMessage = {
        title: message.error.unknownTitle,
        description: message.error.unknownDescription
      }
      break
  }
  return (
    <div className="mx-auto space-y-6 flex flex-wrap justify-center">
      <div className="space-y-2 text-center mb-12 px-8 w-full justify-center">
        <h1 className="text-4xl font-bold">{errorMessage.title}</h1>
        <p className="text-muted-foreground max-w-[512px] m-auto">{errorMessage.description}</p>
      </div>
      <Button className="w-64" onClick={onClickGoBackButton}>
        {message.word.goBack}
      </Button>
    </div>
  )
}
