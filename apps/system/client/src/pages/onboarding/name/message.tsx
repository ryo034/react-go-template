import { useContext } from "react"
import { i18nKeys } from "~/infrastructure/i18n"
import { ContainerContext } from "~/infrastructure/injector/context"

export const useOnboardingSettingNamePageMessage = () => {
  const { i18n } = useContext(ContainerContext)
  return {
    title: i18n.translate(`${i18nKeys.page.onboardingSettingName.title}`),
    description: i18n.translate(`${i18nKeys.page.onboardingSettingName.description}`)
  }
}
