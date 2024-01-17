import {
  AlreadyExistError,
  AuthenticationError,
  BadRequestError,
  CannotConnectNetworkError,
  ForbiddenError,
  InternalServerError,
  NotFoundError,
  RequestTimeoutError
} from "./network"

const SpecificErrorNameMap = {
  CannotConnectNetworkError: CannotConnectNetworkError,
  RequestTimeoutError: RequestTimeoutError,
  BadRequestError: BadRequestError,
  ForbiddenError: ForbiddenError,
  AuthenticationError: AuthenticationError,
  NotFoundError: NotFoundError,
  AlreadyExistError: AlreadyExistError,
  InternalServerError: InternalServerError
} as const

export type RequiredErrorHandlers = {
  [P in keyof typeof SpecificErrorNameMap]: (err: InstanceType<(typeof SpecificErrorNameMap)[P]>) => string
}

/**
 * Provides error messages corresponding to custom errors
 *
 * example:
 * ```ts
const errorHandlers: RequiredErrorHandlers = {
    CannotConnectNetworkError: (err) => "Cannot connect network.",
    RequestTimeoutError: (err) => "Request timeout.",
    BadRequestError: (err) => "Bad request.",
    ForbiddenError: (err) => "Forbidden.",
    AuthenticationError: (err) => "Authentication error.",
    NotFoundError: (err) => "Not found.",
    AlreadyExistError: (err) => "Already exist.",
    InternalServerError: (err) => "Internal server error."
}

const customErrorMessageHandler = (err: Error) => {
    if (err instanceof MyCustomError) {
        return 'this is my custom error message'
    }
    return null
}
const errorMessageProvider = new ErrorHandlingServiceMessageProvider(errorHandlers, customErrorMessageHandler)
 * ```
 */
export class ErrorHandlingServiceMessageProvider {
  constructor(
    readonly errorHandlers: RequiredErrorHandlers,
    readonly customErrorMessageHandler: (err: Error) => string | null
  ) {}

  resolve(err: Error): string {
    const customErrorMessage = this.customErrorMessageHandler(err)
    if (customErrorMessage) {
      return customErrorMessage
    }
    // If the error is not in SpecificErrorNameMap, return "Unknown Error" message.
    if (err instanceof Error && !(err.constructor.name in SpecificErrorNameMap)) {
      return "Unknown Error"
    }
    return this.errorHandlers[err.constructor.name as keyof typeof SpecificErrorNameMap](err as any)
  }
}
