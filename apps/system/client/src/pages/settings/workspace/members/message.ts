import { useContext } from "react"
import { i18nKeys } from "~/infrastructure/i18n"
import { ContainerContext } from "~/infrastructure/injector/context"

export const useSettingsMembersPageMessage = () => {
  const { i18n } = useContext(ContainerContext)
  return {
    word: {
      ownerRole: i18n.translate(`${i18nKeys.word.owner}`),
      adminRole: i18n.translate(`${i18nKeys.word.admin}`),
      memberRole: i18n.translate(`${i18nKeys.word.member}`),
      guestRole: i18n.translate(`${i18nKeys.word.guest}`)
    }
  }
}
