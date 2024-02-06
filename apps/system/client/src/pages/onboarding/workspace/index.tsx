import { useContext, useRef, useState } from "react"
import { SubmitHandler } from "react-hook-form"
import {
  OnboardingSettingWorkspacePageForm,
  OnboardingSettingWorkspacePageFormValues
} from "~/components/onboarding/workspace/form"
import { ContainerContext } from "~/infrastructure/injector/context"
import { useOnboardingSettingWorkspacePageMessage } from "./message"

export const onboardingSettingWorkspacePageRoute = "/onboarding/workspace"

export const OnboardingSettingWorkspacePage = () => {
  const { controller, store } = useContext(ContainerContext)
  const [errorMessage, setErrorMessage] = useState("")

  const me = store.me((state) => state.me)
  const meRef = useRef(me)

  const message = useOnboardingSettingWorkspacePageMessage()

  const onSubmit: SubmitHandler<OnboardingSettingWorkspacePageFormValues> = async (d) => {
    if (me === null || me.self === undefined || me.self.id === undefined || me.self.email === undefined) {
      setErrorMessage("Failed to update profile")
      return
    }
    controller.me.updateProfile({
      user: {
        userId: me.self.id.value.asString,
        name: d.subdomain,
        email: me.self.email.value
      }
    })
    console.log("onSubmit", d)
  }

  return (
    <div className="flex justify-center items-center min-h-screen">
      <div className="mx-auto space-y-6">
        <div className="space-y-2 text-center mb-12 px-8">
          <h1 className="text-4xl font-bold text-gray-800 mb-6">{message.title}</h1>
          <p className="text-gray-500 dark:text-gray-400">{message.description}</p>
        </div>
        <OnboardingSettingWorkspacePageForm onSubmit={onSubmit} errorMessage={errorMessage} isLoading={false} />
      </div>
    </div>
  )
}
