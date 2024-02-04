import { FirebaseError } from "firebase/app"
import { NetworkBaseError, NetworkErrorInterpreter, convertToErrorByStatusCode } from "shared-network"
import { FirebaseErrorAdapter } from "./firebase"

export class EmailAlreadyInUseError extends NetworkBaseError {}
export class InvalidEmailUseError extends NetworkBaseError {}
export class InvalidAddressError extends NetworkBaseError {}
export class EmailNotVerifiedError extends NetworkBaseError {}

export const openapiFetchErrorInterpreter = (res: unknown): Error | null => {
  if (res !== null && typeof res === "object" && "response" in res && (res as any).response instanceof Response) {
    const r = res as {
      data?: undefined
      error?: { code?: number; message?: string }
      response: Response
    }
    return convertToErrorByStatusCode(r.response.status, r.error?.message)
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
