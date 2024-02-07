import i18n from "i18next"
import { initReactI18next } from "react-i18next"

import translationActionEn from "~/infrastructure/i18n/locales/en/action.json"
import translationFormEn from "~/infrastructure/i18n/locales/en/form.json"
import translationPageEn from "~/infrastructure/i18n/locales/en/page.json"
import translationRouterEn from "~/infrastructure/i18n/locales/en/router.json"
import translationWordEn from "~/infrastructure/i18n/locales/en/word.json"
import translationActionJa from "~/infrastructure/i18n/locales/ja/action.json"
import translationFormJa from "~/infrastructure/i18n/locales/ja/form.json"
import translationPageJa from "~/infrastructure/i18n/locales/ja/page.json"
import translationRouterJa from "~/infrastructure/i18n/locales/ja/router.json"
import translationWordJa from "~/infrastructure/i18n/locales/ja/word.json"

const resources = {
  en: {
    translation: {
      ...translationWordEn,
      ...translationFormEn,
      ...translationPageEn,
      ...translationActionEn,
      ...translationRouterEn
    }
  },
  ja: {
    translation: {
      ...translationWordJa,
      ...translationFormJa,
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
  page: "page"
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
    subdomain: `${i18nRootKeys.word}.subdomain`
  },
  action: {
    submit: `${i18nRootKeys.action}.submit`,
    back: `${i18nRootKeys.action}.back`,
    cancel: `${i18nRootKeys.action}.cancel`,
    login: `${i18nRootKeys.action}.login`,
    sendOneTimeCode: `${i18nRootKeys.action}.sendOneTimeCode`,
    startWithEmail: `${i18nRootKeys.action}.startWithEmail`,
    showItem: `${i18nRootKeys.action}.showItem`,
    logout: `${i18nRootKeys.action}.logout`,
    inputField: `${i18nRootKeys.action}.inputField`,
    doAction: `${i18nRootKeys.action}.doAction`,
    enter: `${i18nRootKeys.action}.enter`
  },
  form: {
    required: `${i18nRootKeys.form}.required`,
    regex: `${i18nRootKeys.form}.regex`,
    max: `${i18nRootKeys.form}.max`,
    min: `${i18nRootKeys.form}.min`,
    passwordRegex: `${i18nRootKeys.form}.passwordRegex`,
    placeholder: {
      input: `${i18nRootKeys.form}.placeholder.input`
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
    }
  }
}
