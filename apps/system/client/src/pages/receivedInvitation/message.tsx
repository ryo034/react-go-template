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
      unknownTitle: i18n.translate(`${i18nKeys.page.startInvitation.error.unknownTitle}`),
      unknownDescription: i18n.translate(`${i18nKeys.page.startInvitation.error.unknownDescription}`),
      invitationInvalidTokenTitle: i18n.translate(`${i18nKeys.page.startInvitation.error.invitationInvalidTokenTitle}`),
      invitationInvalidTokenDescription: i18n.translate(
        `${i18nKeys.page.startInvitation.error.invitationInvalidTokenDescription}`
      ),
      invitationAlreadyExpiredTitle: i18n.translate(
        `${i18nKeys.page.startInvitation.error.invitationAlreadyExpiredTitle}`
      ),
      invitationAlreadyExpiredDescription: i18n.translate(
        `${i18nKeys.page.startInvitation.error.invitationAlreadyExpiredDescription}`
      ),
      invitationAlreadyRevokedTitle: i18n.translate(
        `${i18nKeys.page.startInvitation.error.invitationAlreadyRevokedTitle}`
      ),
      invitationAlreadyRevokedDescription: i18n.translate(
        `${i18nKeys.page.startInvitation.error.invitationAlreadyRevokedDescription}`
      ),
      invitationAlreadyAcceptedTitle: i18n.translate(
        `${i18nKeys.page.startInvitation.error.invitationAlreadyAcceptedTitle}`
      ),
      invitationAlreadyAcceptedDescription: i18n.translate(
        `${i18nKeys.page.startInvitation.error.invitationAlreadyAcceptedDescription}`
      )
    }
  }
}
