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
