import { useContext } from "react"
import { MemberRole } from "~/domain"
import { i18nKeys } from "../i18n"
import { ContainerContext } from "../injector/context"

export const useRole = () => {
  const { i18n } = useContext(ContainerContext)
  const translatedRoles = {
    owner: { name: i18n.translate(`${i18nKeys.word.owner}`) },
    admin: { name: i18n.translate(`${i18nKeys.word.admin}`) },
    member: { name: i18n.translate(`${i18nKeys.word.member}`) },
    guest: { name: i18n.translate(`${i18nKeys.word.guest}`) }
  }
  return {
    translateRole(role: MemberRole) {
      return translatedRoles[role].name
    }
  }
}
