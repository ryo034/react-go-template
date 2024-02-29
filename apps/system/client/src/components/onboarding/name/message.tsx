import { useContext } from "react"
import { AccountFullName } from "~/domain"
import { i18nKeys } from "~/infrastructure/i18n"
import { ContainerContext } from "~/infrastructure/injector/context"

export type OnboardingSettingNamePageFormMessage = {
  word: {
    accountName: string
  }
  action: {
    submit: string
  }
  form: {
    placeholder: {
      name: string
    }
    validation: {
      name: {
        required: string
        max: string
        regex: string
      }
    }
  }
}

export const useOnboardingSettingNamePageFormMessage: () => OnboardingSettingNamePageFormMessage = () => {
  const { i18n } = useContext(ContainerContext)
  return {
    word: {
      accountName: i18n.translate(`${i18nKeys.word.accountName}`)
    },
    action: {
      submit: i18n.translate(`${i18nKeys.action.submit}`)
    },
    form: {
      placeholder: {
        name: i18n.translate(`${i18nKeys.form.placeholder.input}`, {
          field: i18n.translate(`${i18nKeys.word.accountName}`)
        })
      },
      validation: {
        name: {
          required: i18n.translate(`${i18nKeys.form.required}`, {
            field: i18n.translate(`${i18nKeys.word.accountName}`)
          }),
          max: i18n.translate(`${i18nKeys.form.max}`, {
            field: i18n.translate(`${i18nKeys.word.accountName}`),
            max: AccountFullName.max
          }),
          regex: i18n.translate(`${i18nKeys.form.regex}`, { field: i18n.translate(`${i18nKeys.word.accountName}`) })
        }
      }
    }
  }
}
