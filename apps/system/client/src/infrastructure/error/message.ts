import { AuthenticationError, NetworkBaseError } from "shared-network"
import {
  AuthProviderCustomError,
  AuthProviderEmailAlreadyInUseError,
  AuthProviderIdTokenExpiredError,
  AuthProviderIdTokenRevokedError,
  AuthProviderInternalError,
  AuthProviderInvalidPasswordError,
  AuthProviderMissingEmail,
  AuthProviderUnverifiedEmailError,
  AuthProviderUserNotFoundError,
  AuthProviderWrongPasswordError
} from "~/infrastructure/error/authProvider"
import {
  EmailAlreadyInUseError,
  EmailNotVerifiedError,
  InvalidAddressError,
  InvalidEmailUseError
} from "~/infrastructure/error/network"
import { type ReactI18nextProvider, i18nKeys } from "~/infrastructure/i18n"

const adaptNetworkError = (err: Error): string | null => {
  if (err instanceof EmailAlreadyInUseError) {
    return "すでにそのメールアドレスは使用されています"
  }
  if (err instanceof InvalidEmailUseError) {
    return "不正なメールアドレスです"
  }
  if (err instanceof InvalidAddressError) {
    return "住所が正しくありません。正しい住所を入力して下さい"
  }
  if (err instanceof EmailNotVerifiedError) {
    return "メールアドレスが認証されていません"
  }
  return null
}

export const adaptAuthProviderError = (err: AuthProviderCustomError): string => {
  let msg = "不明なエラーが発生しました"
  if (err instanceof AuthProviderUserNotFoundError) {
    msg = "メールアドレスが見つかりません"
  } else if (err instanceof AuthProviderMissingEmail) {
    msg = "メールアドレスが見つかりません"
  } else if (err instanceof AuthProviderWrongPasswordError) {
    msg = "メールアドレスまたはパスワードが一致しません"
  } else if (err instanceof AuthProviderInvalidPasswordError) {
    msg = "メールアドレスまたはパスワードが一致しません"
  } else if (err instanceof AuthProviderEmailAlreadyInUseError) {
    msg = "すでにそのメールアドレスは使用されています"
  } else if (err instanceof AuthProviderIdTokenExpiredError) {
  } else if (err instanceof AuthProviderUnverifiedEmailError) {
    msg = "メールアドレスが認証されていません"
  } else if (err instanceof AuthProviderIdTokenRevokedError) {
  } else if (err instanceof AuthProviderInvalidPasswordError) {
    msg = "入力された情報が正しくありません"
  } else if (err instanceof AuthProviderInternalError) {
    msg = "サーバーでエラーが発生しました"
  } else {
  }
  return msg
}

export class MessageProvider {
  constructor(private readonly i18n: ReactI18nextProvider) {}
  translate(err: Error): string | null {
    if (err instanceof NetworkBaseError) {
      return adaptNetworkError(err)
    }
    if (err instanceof AuthProviderCustomError) {
      return adaptAuthProviderError(err)
    }
    return this.i18n.translate(`${i18nKeys.word.error.unknown}`)
  }
}
