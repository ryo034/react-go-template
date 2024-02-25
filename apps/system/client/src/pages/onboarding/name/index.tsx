import { useContext, useState } from "react"
import { SubmitHandler } from "react-hook-form"
import { useNavigate } from "react-router-dom"
import { OnboardingSettingNamePageForm, OnboardingSettingNamePageFormValues } from "~/components/onboarding/name/form"
import { AuthProviderUserNotFoundError } from "~/infrastructure/error"
import { ContainerContext } from "~/infrastructure/injector/context"
import { receivedInvitationsPageRoute } from "~/pages/receivedInvitation"
import { onboardingSettingWorkspacePageRoute } from "../workspace"
import { useOnboardingSettingNamePageMessage } from "./message"

export const onboardingSettingNamePageRoute = "/onboarding/name"

export const OnboardingSettingNamePage = () => {
  const { controller, store, errorMessageProvider } = useContext(ContainerContext)
  const navigate = useNavigate()
  const [errorMessage, setErrorMessage] = useState("")

  const message = useOnboardingSettingNamePageMessage()
  const me = store.me((state) => state.me)
  const isLoading = store.me((state) => state.isLoading)

  const onSubmit: SubmitHandler<OnboardingSettingNamePageFormValues> = async (d) => {
    const res = await controller.me.updateProfile({
      current: me,
      user: { name: d.name }
    })
    if (res) {
      if (res instanceof AuthProviderUserNotFoundError) {
        navigate(onboardingSettingNamePageRoute)
        return
      }
      setErrorMessage(errorMessageProvider.resolve(res))
      return
    }
    if (me?.hasReceivedInvitations === true) {
      navigate(receivedInvitationsPageRoute)
      return
    }
    if (me?.hasNotWorkspace === true) {
      navigate(onboardingSettingWorkspacePageRoute)
    }
  }

  return (
    <div className="flex justify-center items-center min-h-screen">
      <div className="mx-auto space-y-6">
        <div className="space-y-2 text-center mb-12 px-8">
          <h1 className="text-4xl font-bold">{message.title}</h1>
          <p className="text-gray-500 dark:text-gray-400">{message.description}</p>
        </div>
        <OnboardingSettingNamePageForm onSubmit={onSubmit} errorMessage={errorMessage} isLoading={isLoading} />
      </div>
    </div>
  )
}
