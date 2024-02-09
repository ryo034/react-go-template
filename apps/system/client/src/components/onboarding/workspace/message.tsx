import { useContext } from "react"
import { WorkspaceSubdomain } from "~/domain"
import { i18nKeys } from "~/infrastructure/i18n"
import { ContainerContext } from "~/infrastructure/injector/context"

export type OnboardingSettingWorkspacePageFormMessage = {
  word: {
    subdomain: string
  }
  action: {
    submit: string
  }
  form: {
    placeholder: {
      name: string
    }
    validation: {
      subdomain: {
        required: string
        max: string
        regex: string
      }
    }
  }
}

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
            max: WorkspaceSubdomain.max
          }),
          regex: i18n.translate(`${i18nKeys.form.regex}`, { field: i18n.translate(`${i18nKeys.word.subdomain}`) })
        }
      }
    }
  }
}
