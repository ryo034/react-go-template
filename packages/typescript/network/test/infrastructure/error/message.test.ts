import { describe, expect, it } from "vitest"
import {
  AlreadyExistError,
  AuthenticationError,
  BadRequestError,
  CannotConnectNetworkError,
  ErrorHandlingServiceMessageProvider,
  ForbiddenError,
  InternalServerError,
  NotFoundError,
  RequestTimeoutError,
  RequiredErrorHandlers
} from "../../../src"

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
  if (err.message === "MyCustomError") {
    return "this is my custom error message"
  }
  return null
}

const errorMessageProvider = new ErrorHandlingServiceMessageProvider(errorHandlers, customErrorMessageHandler)

describe("ErrorMessageProvider", () => {
  it("should return the correct message for CannotConnectNetworkError", () => {
    const error = new CannotConnectNetworkError("")
    expect(errorMessageProvider.resolve(error)).toBe("Cannot connect network.")
  })

  it("should return the correct message for RequestTimeoutError", () => {
    const error = new RequestTimeoutError(503, "")
    expect(errorMessageProvider.resolve(error)).toBe("Request timeout.")
  })

  it("should return the correct message for BadRequestError", () => {
    const error = new BadRequestError(400, "")
    expect(errorMessageProvider.resolve(error)).toBe("Bad request.")
  })

  it("should return the correct message for ForbiddenError", () => {
    const error = new ForbiddenError(403, "")
    expect(errorMessageProvider.resolve(error)).toBe("Forbidden.")
  })

  it("should return the correct message for AuthenticationError", () => {
    const error = new AuthenticationError(401, "")
    expect(errorMessageProvider.resolve(error)).toBe("Authentication error.")
  })

  it("should return the correct message for NotFoundError", () => {
    const error = new NotFoundError(404, "")
    expect(errorMessageProvider.resolve(error)).toBe("Not found.")
  })

  it("should return the correct message for AlreadyExistError", () => {
    const error = new AlreadyExistError(409, "")
    expect(errorMessageProvider.resolve(error)).toBe("Already exist.")
  })

  it("should return the correct message for InternalServerError", () => {
    const error = new InternalServerError(500, "")
    expect(errorMessageProvider.resolve(error)).toBe("Internal server error.")
  })

  it("should return custom error message for custom errors", () => {
    const error = new Error("MyCustomError")
    expect(errorMessageProvider.resolve(error)).toBe("this is my custom error message")
  })

  it('should return "Unknown Error" for unhandled errors', () => {
    const error = new Error()
    expect(errorMessageProvider.resolve(error)).toBe("Unknown Error")
  })
})
