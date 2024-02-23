import { useContext } from "react"
import { i18nKeys } from "~/infrastructure/i18n"
import { ContainerContext } from "~/infrastructure/injector/context"

export const useReceivedInvitationsPageMessage = () => {
  const { i18n } = useContext(ContainerContext)
  return {
    title: i18n.translate(`${i18nKeys.page.receivedInvitations.title}`),
    action: {
      join: i18n.translate(`${i18nKeys.action.join}`)
    },
    word: {
      inviter: i18n.translate(`${i18nKeys.word.inviter}`)
    }
  }
}
