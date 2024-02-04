import { BadRequestError, NetworkBaseError } from "shared-network"
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
import { ReactI18nextProvider, i18nKeys } from "~/infrastructure/i18n"

const adaptNetworkError = (err: Error): string => {
  let msg = "不明なエラーが発生しました"
  if (err instanceof BadRequestError) {
    msg = "不正なリクエストです"
  } else if (err instanceof EmailAlreadyInUseError) {
    msg = "すでにそのメールアドレスは使用されています"
  } else if (err instanceof InvalidEmailUseError) {
    msg = "不正なメールアドレスです"
  } else if (err instanceof InvalidAddressError) {
    msg = "住所が正しくありません。正しい住所を入力して下さい"
  } else if (err instanceof EmailNotVerifiedError) {
    msg = "メールアドレスが認証されていません"
  }
  return msg
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
  translate(err: Error): string {
    if (err instanceof NetworkBaseError) {
      return adaptNetworkError(err)
    }
    if (err instanceof AuthProviderCustomError) {
      return adaptAuthProviderError(err)
    }
    return this.i18n.translate(`${i18nKeys.word.error.unknown}`)
  }
}
