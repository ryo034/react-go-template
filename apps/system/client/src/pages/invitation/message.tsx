import { useContext } from "react"
import { i18nKeys } from "~/infrastructure/i18n"
import { ContainerContext } from "~/infrastructure/injector/context"

export const useStartInvitationPageMessage = () => {
  const { i18n } = useContext(ContainerContext)
  return {
    title: `${i18nKeys.page.startInvitation.title}`,
    description: (email: string) => i18n.translate(`${i18nKeys.page.startInvitation.description}`, { email }),
    action: {
      start: i18n.translate(`${i18nKeys.action.start}`)
    },
    word: {
      goBack: i18n.translate(`${i18nKeys.word.goBack}`)
    },
    error: {
      invitationAlreadyAcceptedTitle: i18n.translate(
        `${i18nKeys.page.startInvitation.error.invitationAlreadyAcceptedTitle}`
      ),
      invitationInvalidTokenTitle: i18n.translate(`${i18nKeys.page.startInvitation.error.invitationInvalidTokenTitle}`),
      invitationAlreadyRevokedTitle: i18n.translate(
        `${i18nKeys.page.startInvitation.error.invitationAlreadyRevokedTitle}`
      ),
      unknownTitle: i18n.translate(`${i18nKeys.page.startInvitation.error.unknownTitle}`),
      invitationAlreadyAcceptedDescription: i18n.translate(
        `${i18nKeys.page.startInvitation.error.invitationAlreadyAcceptedDescription}`
      ),
      invitationInvalidTokenDescription: i18n.translate(
        `${i18nKeys.page.startInvitation.error.invitationInvalidTokenDescription}`
      ),
      invitationAlreadyRevokedDescription: i18n.translate(
        `${i18nKeys.page.startInvitation.error.invitationAlreadyRevokedDescription}`
      ),
      unknownDescription: i18n.translate(`${i18nKeys.page.startInvitation.error.unknownDescription}`)
    }
  }
}
