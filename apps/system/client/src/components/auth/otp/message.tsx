import { useContext } from "react"
import { i18nKeys } from "~/infrastructure/i18n"
import { ContainerContext } from "~/infrastructure/injector/context"

export const useVerifyOtpPageFormMessage = () => {
  const { i18n } = useContext(ContainerContext)
  return {
    word: {
      submit: i18n.translate(`${i18nKeys.word.submit}`)
    }
  }
}
