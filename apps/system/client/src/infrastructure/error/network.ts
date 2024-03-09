import { FirebaseError } from "firebase/app"
import { NetworkBaseError, NetworkErrorInterpreter, convertToErrorByStatusCode } from "shared-network"
import {
  AlreadyAcceptedInvitationError,
  AlreadyExpiredInvitationError,
  AlreadyRevokeInvitationError
} from "~/domain/workspace/invitation/error"
import { FirebaseErrorAdapter } from "./firebase"

export class EmailAlreadyInUseError extends NetworkBaseError {}
export class InvalidEmailUseError extends NetworkBaseError {}
export class InvalidAddressError extends NetworkBaseError {}
export class EmailNotVerifiedError extends NetworkBaseError {}

export const openapiFetchErrorInterpreter = (res: unknown): Error | null => {
  if (res !== null && typeof res === "object" && "response" in res && (res as any).response instanceof Response) {
    const r = res as {
      data?: undefined
      error?: { code?: string; message?: string }
      response: Response
    }
    const err = customApplicationError(r.error?.message || "", r.error?.code || "")
    if (err !== null) {
      console.error("Network error", err)
      return err
    }
    return convertToErrorByStatusCode(r.response.status, r.error?.message)
  }
  return null
}

const customApplicationError = (message: string, customCode: string): Error | null => {
  if (customCode === "") {
    return null
  }
  switch (customCode) {
    case "410-001":
      return new AlreadyExpiredInvitationError(message)
    case "410-002":
      return new AlreadyRevokeInvitationError(message)
    case "410-003":
      return new AlreadyAcceptedInvitationError(message)
  }
  return null
}

export class SystemNetworkErrorInterpreter extends NetworkErrorInterpreter {
  constructor() {
    super()
    this.convertToSpecificError = this.convertToSpecificError.bind(this)
    this.isValidGenericError = this.isValidGenericError.bind(this)
  }

  convertToSpecificError(err: unknown): Error | null {
    if (this.isValidGenericError(err)) {
      return convertToErrorByStatusCode(err.statusCode, err.message)
    }

    if (err instanceof FirebaseError) {
      return FirebaseErrorAdapter.create(err)
    }

    return openapiFetchErrorInterpreter(err)
  }
}
