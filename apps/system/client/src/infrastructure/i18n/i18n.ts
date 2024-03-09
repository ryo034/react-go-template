import i18n from "i18next"
import { initReactI18next } from "react-i18next"
import translationActionEn from "~/infrastructure/i18n/locales/en/action.json"
import translationFormEn from "~/infrastructure/i18n/locales/en/form.json"
import translationNetworkEn from "~/infrastructure/i18n/locales/en/network.json"
import translationPageEn from "~/infrastructure/i18n/locales/en/page.json"
import translationRouterEn from "~/infrastructure/i18n/locales/en/router.json"
import translationWordEn from "~/infrastructure/i18n/locales/en/word.json"
import translationActionJa from "~/infrastructure/i18n/locales/ja/action.json"
import translationFormJa from "~/infrastructure/i18n/locales/ja/form.json"
import translationNetworkJa from "~/infrastructure/i18n/locales/ja/network.json"
import translationPageJa from "~/infrastructure/i18n/locales/ja/page.json"
import translationRouterJa from "~/infrastructure/i18n/locales/ja/router.json"
import translationWordJa from "~/infrastructure/i18n/locales/ja/word.json"

const resources = {
  en: {
    translation: {
      ...translationWordEn,
      ...translationFormEn,
      ...translationNetworkEn,
      ...translationPageEn,
      ...translationActionEn,
      ...translationRouterEn
    }
  },
  ja: {
    translation: {
      ...translationWordJa,
      ...translationFormJa,
      ...translationNetworkJa,
      ...translationPageJa,
      ...translationActionJa,
      ...translationRouterJa
    }
  }
}

const DEFAULT_LANGUAGE = "ja"

export function initI18n() {
  i18n.use(initReactI18next).init({
    lng: DEFAULT_LANGUAGE,
    fallbackLng: DEFAULT_LANGUAGE,
    debug: false,
    interpolation: {
      escapeValue: false
    },
    resources
  })
}

export const i18nRootKeys = {
  word: "word",
  action: "action",
  form: "form",
  page: "page",
  network: "network"
}

export const i18nKeys = {
  word: {
    email: `${i18nRootKeys.word}.email`,
    error: {
      unknown: `${i18nRootKeys.word}.error.unknown`
    },
    submit: `${i18nRootKeys.word}.submit`,
    otp: `${i18nRootKeys.word}.otp`,
    accountName: `${i18nRootKeys.word}.accountName`,
    subdomain: `${i18nRootKeys.word}.subdomain`,
    displayName: `${i18nRootKeys.word}.displayName`,
    goBack: `${i18nRootKeys.word}.goBack`,
    inviter: `${i18nRootKeys.word}.inviter`,
    google: `${i18nRootKeys.word}.google`,
    owner: `${i18nRootKeys.word}.owner`,
    admin: `${i18nRootKeys.word}.admin`,
    member: `${i18nRootKeys.word}.member`,
    guest: `${i18nRootKeys.word}.guest`
  },
  action: {
    submit: `${i18nRootKeys.action}.submit`,
    back: `${i18nRootKeys.action}.back`,
    cancel: `${i18nRootKeys.action}.cancel`,
    login: `${i18nRootKeys.action}.login`,
    start: `${i18nRootKeys.action}.start`,
    startWith: `${i18nRootKeys.action}.startWith`,
    sendOneTimeCode: `${i18nRootKeys.action}.sendOneTimeCode`,
    showItem: `${i18nRootKeys.action}.showItem`,
    logout: `${i18nRootKeys.action}.logout`,
    inputField: `${i18nRootKeys.action}.inputField`,
    doAction: `${i18nRootKeys.action}.doAction`,
    add: `${i18nRootKeys.action}.add`,
    enter: `${i18nRootKeys.action}.enter`,
    invite: `${i18nRootKeys.action}.invite`,
    inviting: `${i18nRootKeys.action}.inviting`,
    inviteTarget: `${i18nRootKeys.action}.inviteTarget`,
    successInvite: `${i18nRootKeys.action}.successInvite`,
    failedInvite: `${i18nRootKeys.action}.failedInvite`,
    join: `${i18nRootKeys.action}.join`
  },
  form: {
    required: `${i18nRootKeys.form}.required`,
    regex: `${i18nRootKeys.form}.regex`,
    max: `${i18nRootKeys.form}.max`,
    min: `${i18nRootKeys.form}.min`,
    passwordRegex: `${i18nRootKeys.form}.passwordRegex`,
    placeholder: {
      input: `${i18nRootKeys.form}.placeholder.input`,
      optional: `${i18nRootKeys.form}.placeholder.optional`
    }
  },
  page: {
    auth: {
      title: `${i18nRootKeys.page}.auth.title`
    },
    verifyOtp: {
      enterOtpMessage: `${i18nRootKeys.page}.verifyOtp.enterOtpMessage`
    },
    onboardingSettingName: {
      title: `${i18nRootKeys.page}.onboardingSettingName.title`,
      description: `${i18nRootKeys.page}.onboardingSettingName.description`
    },
    onboardingSettingWorkspace: {
      title: `${i18nRootKeys.page}.onboardingSettingWorkspace.title`,
      description: `${i18nRootKeys.page}.onboardingSettingWorkspace.description`,
      error: {
        alreadyExist: `${i18nRootKeys.page}.onboardingSettingWorkspace.error.alreadyExist`
      }
    },
    startInvitation: {
      title: `${i18nRootKeys.page}.startInvitation.title`,
      description: `${i18nRootKeys.page}.startInvitation.description`,
      error: {
        invitationAlreadyAcceptedTitle: `${i18nRootKeys.page}.startInvitation.error.invitationAlreadyAcceptedTitle`,
        invitationInvalidTokenTitle: `${i18nRootKeys.page}.startInvitation.error.invitationInvalidTokenTitle`,
        invitationAlreadyExpiredTitle: `${i18nRootKeys.page}.startInvitation.error.invitationAlreadyExpiredTitle`,
        invitationAlreadyRevokedTitle: `${i18nRootKeys.page}.startInvitation.error.invitationAlreadyRevokedTitle`,
        unknownTitle: `${i18nRootKeys.page}.startInvitation.error.unknownTitle`,
        invitationAlreadyExpiredDescription: `${i18nRootKeys.page}.startInvitation.error.invitationAlreadyExpiredDescription`,
        invitationAlreadyAcceptedDescription: `${i18nRootKeys.page}.startInvitation.error.invitationAlreadyAcceptedDescription`,
        invitationInvalidTokenDescription: `${i18nRootKeys.page}.startInvitation.error.invitationInvalidTokenDescription`,
        invitationAlreadyRevokedDescription: `${i18nRootKeys.page}.startInvitation.error.invitationAlreadyRevokedDescription`,
        unknownDescription: `${i18nRootKeys.page}.startInvitation.error.unknownDescription`
      }
    },
    receivedInvitations: {
      title: `${i18nRootKeys.page}.receivedInvitations.title`
    }
  },
  network: {
    cannotConnect: `${i18nRootKeys.network}.error.cannotConnect`,
    requestTimeout: `${i18nRootKeys.network}.error.requestTimeout`,
    badRequest: `${i18nRootKeys.network}.error.badRequest`,
    forbidden: `${i18nRootKeys.network}.error.forbidden`,
    authentication: `${i18nRootKeys.network}.error.authentication`,
    notFound: `${i18nRootKeys.network}.error.notFound`,
    alreadyExist: `${i18nRootKeys.network}.error.alreadyExist`,
    internalServer: `${i18nRootKeys.network}.error.internalServer`
  }
}
