import { Message } from "@bufbuild/protobuf"
import { Code, ConnectError } from "@connectrpc/connect"
import { ErrorCode } from "~/generated/schema/api/error/v1/error_pb"
import {
  AlreadyExistError,
  AuthenticationError,
  BadRequestError,
  EmailAlreadyInUseError,
  EmailNotVerifiedError,
  ForbiddenError,
  InternalServerError,
  InvalidAddressError,
  InvalidEmailUseError,
  NotFoundError
} from "~/infrastructure/error/network"
import { HttpStatusCode } from "~/infrastructure/error/statusCode"

interface CustomErrorDetail {
  reason: string
  metadata: {
    [key: string]: string
  }
}

export class ConnectNetworkError extends Error {
  static statusCode(err: ConnectError): number {
    switch (err.code) {
      case Code.InvalidArgument:
        return HttpStatusCode.BAD_REQUEST
      case Code.PermissionDenied:
        return HttpStatusCode.FORBIDDEN
      case Code.Unauthenticated:
        return HttpStatusCode.UNAUTHORIZED
      case Code.NotFound:
        return HttpStatusCode.NOT_FOUND
      case Code.AlreadyExists:
        return HttpStatusCode.CONFLICT
      case Code.DeadlineExceeded:
        return HttpStatusCode.REQUEST_TIMEOUT
      case Code.Internal:
        return HttpStatusCode.INTERNAL_SERVER_ERROR
      default:
        return HttpStatusCode.INTERNAL_SERVER_ERROR
    }
  }

  static create(err: ConnectError): Error {
    const statusCode = ConnectNetworkError.statusCode(err)

    const { details, code, message } = err
    if (details.length !== 0) {
      for (let idx = 0; idx < details.length; idx++) {
        const detail = details[idx]
        if (detail instanceof Message) {
          // connectErrorDetails(err, detail, [])
        } else {
          if (detail.debug !== undefined && detail.debug !== null) {
            const { metadata, reason } = detail.debug as unknown as CustomErrorDetail
            if (metadata === undefined || metadata === null) {
              continue
            }
            if (/^\d+$/.test(metadata.code)) {
              switch (Number(metadata.code)) {
                case ErrorCode.EMAIL_ALREADY_IN_USE:
                  return new EmailAlreadyInUseError(statusCode, reason)
                case ErrorCode.INVALID_EMAIL:
                  return new InvalidEmailUseError(statusCode, reason)
                // case ErrorCode.PHONE_NUMBER_ALREADY_IN_USE:
                //   return new PhoneNumberAlreadyInUseError(statusCode, reason)
                case ErrorCode.INVALID_ADDRESS:
                  return new InvalidAddressError(statusCode, reason)
                case ErrorCode.EMAIL_NOT_VERIFIED:
                  return new EmailNotVerifiedError(statusCode, reason)
              }
            }

            switch (metadata.code) {
              case "ERROR_CODE_EMAIL_ALREADY_IN_USE":
                return new EmailAlreadyInUseError(statusCode, reason)
              case "ERROR_CODE_INVALID_EMAIL":
                return new InvalidEmailUseError(statusCode, reason)
              // case "ERROR_CODE_PHONE_NUMBER_ALREADY_IN_USE":
              //   return new PhoneNumberAlreadyInUseError(statusCode, reason)
              case "ERROR_CODE_INVALID_ADDRESS":
                return new InvalidAddressError(statusCode, reason)
              case "ERROR_CODE_EMAIL_NOT_VERIFIED":
                return new EmailNotVerifiedError(statusCode, reason)
            }
          }
        }
      }
    }

    switch (code) {
      case Code.InvalidArgument:
        return new BadRequestError(statusCode, message)
      case Code.PermissionDenied:
        return new ForbiddenError(statusCode, message)
      case Code.Unauthenticated:
        return new AuthenticationError(statusCode, message)
      case Code.NotFound:
        return new NotFoundError(statusCode, message)
      case Code.AlreadyExists:
        return new AlreadyExistError(statusCode, message)
      case Code.DeadlineExceeded:
        return new InternalServerError(statusCode, message)
      case Code.Internal:
        return new InternalServerError(statusCode, message)
      default:
        return new InternalServerError(statusCode, message)
    }
  }
}
