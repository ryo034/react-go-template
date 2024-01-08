import { describe, expect, it } from "vitest"
import {
  AlreadyExistError,
  AuthenticationError,
  BadRequestError,
  ForbiddenError,
  HttpStatusCode,
  NetworkErrorInterpreter,
  NotFoundError,
  UnknownError,
  convertToErrorByStatusCode
} from "../../../src"

class TestableNetworkErrorInterpreter extends NetworkErrorInterpreter {
  convertToSpecificError(): Error | null {
    throw new Error("Method not implemented.")
  }

  testIsValidGenericError(error: unknown): boolean {
    return this.isValidGenericError(error)
  }
}

describe("TestableNetworkErrorInterpreter", () => {
  const handler = new TestableNetworkErrorInterpreter()

  it("correctly identifies generic errors", () => {
    const genericError = { statusCode: HttpStatusCode.BAD_REQUEST, message: "Error", code: "some-code" }
    expect(handler.testIsValidGenericError(genericError)).toBeTruthy()

    const nonGenericError = new Error("Regular error")
    expect(handler.testIsValidGenericError(nonGenericError)).toBeFalsy()
  })
})

describe("convertToErrorByStatusCode", () => {
  it("returns BadRequestError for status code 400", () => {
    const error = convertToErrorByStatusCode(HttpStatusCode.BAD_REQUEST)
    expect(error).toBeInstanceOf(BadRequestError)
    expect(error.message).toBe("Bad request")
  })

  it("returns AuthenticationError for status code 401", () => {
    const error = convertToErrorByStatusCode(HttpStatusCode.UNAUTHORIZED)
    expect(error).toBeInstanceOf(AuthenticationError)
    expect(error.message).toBe("Unauthorized")
  })

  it("returns ForbiddenError for status code 403", () => {
    const error = convertToErrorByStatusCode(HttpStatusCode.FORBIDDEN)
    expect(error).toBeInstanceOf(ForbiddenError)
    expect(error.message).toBe("Forbidden")
  })

  it("returns NotFoundError for status code 404", () => {
    const error = convertToErrorByStatusCode(HttpStatusCode.NOT_FOUND)
    expect(error).toBeInstanceOf(NotFoundError)
    expect(error.message).toBe("Not found")
  })

  it("returns AlreadyExistError for status code 409", () => {
    const error = convertToErrorByStatusCode(HttpStatusCode.CONFLICT)
    expect(error).toBeInstanceOf(AlreadyExistError)
    expect(error.message).toBe("Conflict")
  })

  it("returns UnknownError for an unknown status code", () => {
    const error = convertToErrorByStatusCode(999)
    expect(error).toBeInstanceOf(UnknownError)
    expect(error.message).toBe("Unknown error")
  })

  it("overrides the default message if provided", () => {
    const customMessage = "Custom error message"
    const error = convertToErrorByStatusCode(HttpStatusCode.BAD_REQUEST, customMessage)
    expect(error).toBeInstanceOf(BadRequestError)
    expect(error.message).toBe(customMessage)
  })
})
