import {
  AlreadyExistError,
  AuthenticationError,
  BadRequestError,
  ForbiddenError,
  HttpStatusCode,
  InternalServerError,
  NotFoundError,
  UnknownError
} from "."

interface GenericError {
  statusCode: number
  message: string
  code?: string
}

/**
 * Each service should implement this class to return custom error
 *
 * example:
 * ```ts
class MyNetworkErrorInterpreter extends NetworkErrorInterpreter {
  convertToSpecificError(error: unknown): Error | null {
    if (!this.isValidGenericError(error)) {
      return null
    }
    switch (error.code) {
      case "03-0001":
        return new MyCustomError(error.statusCode, error.message)
      default:
        return null
    }
  }
}
 * ```
 */
export abstract class NetworkErrorInterpreter {
  abstract convertToSpecificError(error: unknown): Error | null

  protected isValidGenericError(error: unknown): error is GenericError {
    const e = error as GenericError
    return e && typeof e.statusCode === "number" && typeof e.message === "string" && typeof e.code === "string"
  }
}

export const convertToErrorByStatusCode = (statusCode: number, message?: string): Error => {
  switch (statusCode) {
    case HttpStatusCode.BAD_REQUEST:
      return new BadRequestError(statusCode, message || "Bad request")
    case HttpStatusCode.UNAUTHORIZED:
      return new AuthenticationError(statusCode, message || "Unauthorized")
    case HttpStatusCode.FORBIDDEN:
      return new ForbiddenError(statusCode, message || "Forbidden")
    case HttpStatusCode.NOT_FOUND:
      return new NotFoundError(statusCode, message || "Not found")
    case HttpStatusCode.CONFLICT:
      return new AlreadyExistError(statusCode, message || "Conflict")
    case HttpStatusCode.INTERNAL_SERVER_ERROR:
      return new InternalServerError(statusCode, message || "Internal server error")
    default:
      return new UnknownError(message || "Unknown error")
  }
}
