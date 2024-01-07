import { beforeEach, describe, expect, it, vi } from "vitest"
import { ApiErrorHandler, CannotConnectNetworkError, UnknownError } from "../../../src/infrastructure/error"

describe("ApiErrorHandler", () => {
  beforeEach(() => {
    vi.restoreAllMocks()
  })

  it("should return CannotConnectNetworkError when offline", () => {
    vi.spyOn(window.navigator, "onLine", "get").mockReturnValue(false)
    const handler = new ApiErrorHandler(() => null)
    const error = handler.adapt(new Error("network error"))
    expect(error).toBeInstanceOf(CannotConnectNetworkError)
  })

  it("should return custom error when provided", () => {
    const customError = new Error("custom error")
    const customErrorFunc = () => customError
    const handler = new ApiErrorHandler(customErrorFunc)
    const error = handler.adapt(new Error("some error"))
    expect(error).toBe(customError)
  })

  it("should return the same error object for error instances", () => {
    const originalError = new Error("error")
    const handler = new ApiErrorHandler(() => null)
    const error = handler.adapt(originalError)
    expect(error).toBe(originalError)
  })

  it("should return UnknownError for string errors", () => {
    const handler = new ApiErrorHandler(() => null)
    const error = handler.adapt("string error")
    expect(error).toBeInstanceOf(UnknownError)
    expect(error.message).toBe("string error")
  })

  it("should return UnknownError for non-string and non-error objects", () => {
    const handler = new ApiErrorHandler(() => null)
    const error = handler.adapt({ some: "object" })
    expect(error).toBeInstanceOf(UnknownError)
    expect(error.message).toBe("unknown error")
  })

  it("should return Error when customErrorChecker returns null", () => {
    const handler = new ApiErrorHandler(() => null)
    const err = new Error("error")
    const error = handler.adapt(err)
    expect(error).toBeInstanceOf(Error)
  })
})
