import { useContext } from "react"
import { i18nKeys } from "~/infrastructure/i18n"
import { ContainerContext } from "~/infrastructure/injector/context"

export const usePasswordInputComponentMessage = () => {
  const { i18n } = useContext(ContainerContext)
  return {
    action: {
      showPassword: i18n.translate(`${i18nKeys.action.showItem}`, {
        field: i18n.translate(`${i18nKeys.word.password}`)
      })
    }
  }
}
