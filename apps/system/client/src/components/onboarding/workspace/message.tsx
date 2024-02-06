import { useContext } from "react"
import { AccountName } from "~/domain"
import { i18nKeys } from "~/infrastructure/i18n"
import { ContainerContext } from "~/infrastructure/injector/context"

export const useOnboardingSettingWorkspacePageFormMessage = () => {
  const { i18n } = useContext(ContainerContext)
  return {
    word: {
      subdomain: i18n.translate(`${i18nKeys.word.subdomain}`)
    },
    action: {
      submit: i18n.translate(`${i18nKeys.action.submit}`)
    },
    form: {
      placeholder: {
        name: i18n.translate(`${i18nKeys.form.placeholder.input}`, {
          field: i18n.translate(`${i18nKeys.word.subdomain}`)
        })
      },
      validation: {
        subdomain: {
          required: i18n.translate(`${i18nKeys.form.required}`, {
            field: i18n.translate(`${i18nKeys.word.subdomain}`)
          }),
          max: i18n.translate(`${i18nKeys.form.max}`, {
            field: i18n.translate(`${i18nKeys.word.email}`),
            max: AccountName.max
          }),
          regex: i18n.translate(`${i18nKeys.form.regex}`, { field: i18n.translate(`${i18nKeys.word.subdomain}`) })
        }
      }
    }
  }
}
