import { NotFoundError } from "shared-network"
import { describe, expect, it } from "vitest"
import { openapiFetchErrorInterpreter } from "~/infrastructure/error"

describe("openapiFetchErrorInterpreter", () => {
  it("returns NotFoundError for openapi-fetch error status 404", () => {
    const mockResponse = { response: new Response(null, { status: 404 }) }
    expect(openapiFetchErrorInterpreter(mockResponse)).toBeInstanceOf(NotFoundError)
  })

  it("returns null for invalid response object", () => {
    const invalidResponse = { someInvalidProperty: 123 }
    expect(openapiFetchErrorInterpreter(invalidResponse)).toBeNull()
  })
})
