import { useContext } from "react"
import { i18nKeys } from "~/infrastructure/i18n"
import { ContainerContext } from "~/infrastructure/injector/context"

export const useOnboardingSettingWorkspacePageMessage = () => {
  const { i18n } = useContext(ContainerContext)
  return {
    title: i18n.translate(`${i18nKeys.page.onboardingSettingWorkspace.title}`),
    description: i18n.translate(`${i18nKeys.page.onboardingSettingWorkspace.description}`),
    error: {
      alreadyExist: i18n.translate(`${i18nKeys.page.onboardingSettingWorkspace.error.alreadyExist}`)
    }
  }
}
