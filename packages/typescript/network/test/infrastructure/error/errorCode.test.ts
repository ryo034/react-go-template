import { describe, expect, it } from "vitest"
import { NetworkErrorInterpreter } from "../../../src"

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
    const genericError = { statusCode: 400, message: "Error", code: "some-code" }
    expect(handler.testIsValidGenericError(genericError)).toBeTruthy()

    const nonGenericError = new Error("Regular error")
    expect(handler.testIsValidGenericError(nonGenericError)).toBeFalsy()
  })
})
