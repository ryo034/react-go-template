import { useContext } from "react"
import { i18nKeys } from "~/infrastructure/i18n"
import { ContainerContext } from "~/infrastructure/injector/context"

export const useInviteMembersFormMessage = () => {
  const { i18n } = useContext(ContainerContext)
  return {
    action: {
      invite: i18n.translate(`${i18nKeys.action.invite}`),
      inviteMember: i18n.translate(`${i18nKeys.action.inviteTarget}`, {
        target: i18n.translate(`${i18nKeys.word.member}`)
      }),
      add: i18n.translate(`${i18nKeys.action.add}`)
    },
    form: {
      email: {
        placeholder: i18n.translate(`${i18nKeys.word.email}`),
        required: i18n.translate(`${i18nKeys.form.required}`),
        regex: i18n.translate(`${i18nKeys.form.regex}`)
      },
      displayName: {
        placeholder: i18n.translate(`${i18nKeys.form.placeholder.optional}`, {
          value: i18n.translate(`${i18nKeys.word.displayName}`)
        })
      }
    }
  }
}
