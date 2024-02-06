import { useContext, useRef, useState } from "react"
import { SubmitHandler } from "react-hook-form"
import { useToast } from "shared-ui"
import { OnboardingSettingNamePageForm, OnboardingSettingNamePageFormValues } from "~/components/onboarding/name/form"
import { ContainerContext } from "~/infrastructure/injector/context"
import { useOnboardingSettingNamePageMessage } from "./message"

export const onboardingSettingNamePageRoute = "/onboarding/name"

export const OnboardingSettingNamePage = () => {
  const { controller, store } = useContext(ContainerContext)
  const [errorMessage, setErrorMessage] = useState("")

  const message = useOnboardingSettingNamePageMessage()
  const me = store.me((state) => state.me)
  const meRef = useRef(me)

  const onSubmit: SubmitHandler<OnboardingSettingNamePageFormValues> = async (d) => {
    console.log("onSubmit", d)
    if (me === null || me.self === undefined || me.self.id === undefined || me.self.email === undefined) {
      setErrorMessage("Failed to update profile")
      return
    }
    const res = await controller.me.updateProfile({
      user: {
        userId: me.self.id.value.asString,
        name: d.name,
        email: me.self.email.value
      }
    })
    if (!res) {
      setErrorMessage("Failed to update name")
      return
    }
  }

  return (
    <div className="flex justify-center items-center min-h-screen">
      <div className="mx-auto space-y-6">
        <div className="space-y-2 text-center mb-12 px-8">
          <h1 className="text-4xl font-bold">{message.title}</h1>
          <p className="text-gray-500 dark:text-gray-400">{message.description}</p>
        </div>
        <OnboardingSettingNamePageForm onSubmit={onSubmit} errorMessage={errorMessage} isLoading={false} />
      </div>
    </div>
  )
}
