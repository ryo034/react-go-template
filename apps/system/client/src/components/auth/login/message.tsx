import { useContext } from "react"
import { Email } from "~/domain"
import { i18nKeys } from "~/infrastructure/i18n"
import { ContainerContext } from "~/infrastructure/injector/context"

export const useLoginPageFormMessage = () => {
  const { i18n } = useContext(ContainerContext)
  return {
    forgotPassword: i18n.translate(`${i18nKeys.page.login.forgotPassword}`),
    word: {
      email: i18n.translate(`${i18nKeys.word.email}`)
    },
    action: {
      login: i18n.translate(`${i18nKeys.action.login}`),
      sendOneTimeCode: i18n.translate(`${i18nKeys.action.sendOneTimeCode}`)
    },
    form: {
      validation: {
        email: {
          required: i18n.translate(`${i18nKeys.form.required}`, { field: i18n.translate(`${i18nKeys.word.email}`) }),
          regex: i18n.translate(`${i18nKeys.form.regex}`, { field: i18n.translate(`${i18nKeys.word.email}`) }),
          max: i18n.translate(`${i18nKeys.form.max}`, {
            field: i18n.translate(`${i18nKeys.word.email}`),
            max: Email.max
          })
        }
      }
    }
  }
}
