import { NetworkBaseError, NetworkErrorInterpreter, convertToErrorByStatusCode } from "shared-network"

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
  convertToSpecificError(err: unknown): Error | null {
    console.log("convertToSpecificError", err)

    if (this.isValidGenericError(err)) {
      return convertToErrorByStatusCode(err.statusCode, err.message)
    }
    // Firebaseのエラーをここに追加する

    return openapiFetchErrorInterpreter(err)
  }
}
