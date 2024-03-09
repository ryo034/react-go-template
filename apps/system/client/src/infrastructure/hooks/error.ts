import { useContext } from "react"
import { useNavigate } from "react-router-dom"
import { ErrorHandlingServiceMessageProvider, isAuthenticationError } from "shared-network"
import { MessageProvider } from "../error"
import { i18nKeys } from "../i18n"
import { ContainerContext } from "../injector/context"
import { routeMap } from "../route/path"

export const useErrorHandler = () => {
  const navigate = useNavigate()
  const handleError = (error: Error) => {
    if (isAuthenticationError(error)) {
      navigate(routeMap.auth)
      return
    }
  }
  return {
    handleError
  }
}

export const useErrorMessageHandler = () => {
  const { i18n } = useContext(ContainerContext)

  const systemErrorMessageProvider = new MessageProvider(i18n)

  // can not pass hooks directly, so pass the string from i18n to hooks
  const cannotConnectNetworkErrorMessages = i18n.translate(i18nKeys.network.cannotConnect)
  const requestTimeoutErrorMessages = i18n.translate(i18nKeys.network.requestTimeout)
  const badRequestErrorMessages = i18n.translate(i18nKeys.network.badRequest)
  const forbiddenErrorMessages = i18n.translate(i18nKeys.network.forbidden)
  const authenticationErrorMessages = i18n.translate(i18nKeys.network.authentication)
  const notFoundErrorMessages = i18n.translate(i18nKeys.network.notFound)
  const alreadyExistErrorMessages = i18n.translate(i18nKeys.network.alreadyExist)
  const internalServerErrorMessages = i18n.translate(i18nKeys.network.internalServer)

  const errorMessageProvider = new ErrorHandlingServiceMessageProvider(
    {
      CannotConnectNetworkError: (_err) => cannotConnectNetworkErrorMessages,
      RequestTimeoutError: (_err) => requestTimeoutErrorMessages,
      BadRequestError: (_err) => badRequestErrorMessages,
      ForbiddenError: (_err) => forbiddenErrorMessages,
      AuthenticationError: (_err) => authenticationErrorMessages,
      NotFoundError: (_err) => notFoundErrorMessages,
      AlreadyExistError: (_err) => alreadyExistErrorMessages,
      InternalServerError: (_err) => internalServerErrorMessages
    },
    systemErrorMessageProvider.translate
  )

  const handleErrorMessage = (error: Error) => {
    const message = errorMessageProvider.resolve(error)
    if (message) {
      return message
    }
    return i18n.translate(i18nKeys.network.internalServer)
  }

  return {
    handleErrorMessage
  }
}
