import { useContext, useState } from "react"
import { SubmitHandler } from "react-hook-form"
import { useNavigate } from "react-router-dom"
import { AlreadyExistError } from "shared-network"
import { useToast } from "shared-ui"
import {
  OnboardingSettingWorkspacePageForm,
  OnboardingSettingWorkspacePageFormValues
} from "~/components/onboarding/workspace/form"
import { ContainerContext } from "~/infrastructure/injector/context"
import { routeMap } from "~/infrastructure/route/path"
import { useOnboardingSettingWorkspacePageMessage } from "./message"

export const onboardingSettingWorkspacePageRoute = "/onboarding/workspace"

export const OnboardingSettingWorkspacePage = () => {
  const { controller } = useContext(ContainerContext)
  const [errorMessage, setErrorMessage] = useState("")
  const navigate = useNavigate()
  const { toast } = useToast()
  const message = useOnboardingSettingWorkspacePageMessage()

  const onSubmit: SubmitHandler<OnboardingSettingWorkspacePageFormValues> = async (d) => {
    const res = await controller.workspace.create({
      subdomain: d.subdomain
    })
    if (res) {
      if (res instanceof AlreadyExistError) {
        setErrorMessage(message.error.alreadyExist)
        return
      }
      setErrorMessage(res.message)
      return
    }
    toast({ title: "ワークスペースを作成しました🎉" })
    navigate(routeMap.home)
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
