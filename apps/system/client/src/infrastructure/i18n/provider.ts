import { useTranslation } from "react-i18next"

export interface I18nProvider {
  translate(key: string, variables?: { [k: string]: any } | undefined): string
}

export class ReactI18nextProvider implements I18nProvider {
  translate(key: string, variables?: { [k: string]: any } | undefined): string {
    const { t } = useTranslation()
    return t(key, variables)
  }
}
