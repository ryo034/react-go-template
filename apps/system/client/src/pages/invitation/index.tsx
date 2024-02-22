import { CheckIcon } from "lucide-react"
import { useContext, useEffect, useLayoutEffect, useMemo, useRef, useState } from "react"
import { Trans } from "react-i18next"
import { useLocation, useNavigate, useSearchParams } from "react-router-dom"
import { Button, Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle, cn } from "shared-ui"
import { firebaseAuth } from "~/infrastructure/firebase"
import { ContainerContext } from "~/infrastructure/injector/context"
import { useStartInvitationPageMessage } from "./message"

export const startInvitationPageRoute = "/invitation"

const errorMessageTypeList = {
  alreadyAccepted: "alreadyAccepted",
  invalidToken: "invalidToken",
  alreadyRevoked: "alreadyRevoked",
  unknown: "unknown"
} as const

type ErrorMessageType = (typeof errorMessageTypeList)[keyof typeof errorMessageTypeList]

type ErrorMessage = {
  title: string
  description: string
}

export const StartInvitationPage = () => {
  const { store, controller } = useContext(ContainerContext)
  const navigate = useNavigate()
  const [searchParams] = useSearchParams()
  const [errorMessageType, setErrorMessageType] = useState<ErrorMessageType | null>(null)

  const message = useStartInvitationPageMessage()
  const meIsLoading = store.me((state) => state.isLoading)
  const me = store.me((s) => s.me)
  const meRef = useRef(me)
  const meIsLoadingRef = useRef(meIsLoading)

  useLayoutEffect(() => {
    store.me.subscribe((state) => {
      meRef.current = state.me
      meIsLoadingRef.current = state.isLoading
    })

    const unsubscribed = firebaseAuth.onAuthStateChanged(async () => {
      console.log("onAuthStateChanged")
      await controller.me.find()
    })
    return () => unsubscribed()
  }, [])

  useMemo(() => {
    // TODO: call API to get info from token
    //  クエリパラメータにtokenがない場合はエラーメッセージを表示する
    console.log("searchParams", searchParams)
    console.log("searchParams", searchParams.get("token"))
    if (!searchParams.get("token")) {
      setErrorMessageType("alreadyAccepted")
    }
  }, [searchParams])

  const onClickStartButton = () => {
    // TODO: call API to verfiy token
    navigate("/")
  }

  return (
    <>
      <InvitationPageHeader />
      <div className="flex justify-center items-center min-h-screen">
        <InvitationSection
          titleI18nKey={message.title}
          workspaceName="workspaceName"
          inviterName="inviterName"
          description={message.description("uni.nashi.034+1@gmail.com")}
          startButtonLabel={message.action.start}
          onClick={onClickStartButton}
          errorMessageType={errorMessageType}
        />
      </div>
    </>
  )
}

const InvitationPageHeader = () => {
  const { store } = useContext(ContainerContext)

  const me = store.me((state) => state.me)
  if (me === null) return null

  return (
    <header className="fixed top-0 z-50 w-full flex flex-wrap justify-end p-6">
      <Button variant="ghost" className="h-14 flex flex-wrap text-right">
        <p className="w-full text-xs text-muted-foreground">Logged in as:</p>
        <p className="w-full font-bold">{me.self.email.value}</p>
      </Button>
    </header>
  )
}

const InvitationSection = ({
  titleI18nKey,
  description,
  workspaceName,
  inviterName,
  startButtonLabel,
  onClick,
  errorMessageType
}: {
  titleI18nKey: string
  workspaceName: string
  inviterName: string
  description: string
  startButtonLabel: string
  onClick: () => void
  errorMessageType: ErrorMessageType | null
}) => {
  if (errorMessageType !== null) {
    return <ErrorMessageSection errorMessageType={errorMessageType} />
  }
  return (
    <div className="mx-auto space-y-6 flex flex-wrap justify-center">
      <div className="space-y-2 text-center mb-12 px-8 w-full">
        <h1 className="text-4xl font-bold">
          <Trans i18nKey={titleI18nKey} values={{ inviterName, workspaceName }} />
        </h1>
        <p className="text-muted-foreground">{description}</p>
      </div>
      <Button className="w-[256px]" onClick={onClick} data-testid="startButtonFromInvitation">
        {startButtonLabel}
      </Button>
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
    case "alreadyAccepted":
      errorMessage = {
        title: message.error.invitationAlreadyAcceptedTitle,
        description: message.error.invitationAlreadyAcceptedDescription
      }
      break
    case "invalidToken":
      errorMessage = {
        title: message.error.invitationInvalidTokenTitle,
        description: message.error.invitationInvalidTokenDescription
      }
      break
    case "alreadyRevoked":
      errorMessage = {
        title: message.error.invitationAlreadyRevokedTitle,
        description: message.error.invitationAlreadyRevokedDescription
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
        <p className="text-muted-foreground">{errorMessage.description}</p>
      </div>
      <Button className="w-64" onClick={onClickGoBackButton}>
        {message.word.goBack}
      </Button>
    </div>
  )
}
