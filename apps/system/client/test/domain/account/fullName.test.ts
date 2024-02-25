import { describe, expect, it } from "vitest"
import { AccountFullName } from "~/domain/account"
import { DomainError } from "~/domain/shared"

describe("AccountFullName", () => {
  describe("create", () => {
    it("should fail to create an instance of AccountFullName with invalid input %s", () => {
      const actual = "a".repeat(AccountFullName.max + 1)
      const result = AccountFullName.create(actual)
      expect(result.isErr).toBe(true)
      result.mapErr((e) => expect(e).instanceOf(DomainError))
    })

    it.each(["一郎", "いちろう", "鈴木 一郎", "鈴木　一郎", "イチロウ", "ichiroh", "John Due"])(
      "should create an instance of AccountFullName with valid input %s",
      (actual) => {
        const result = AccountFullName.create(actual)
        expect(result.isOk).toBe(true)
      }
    )

    it.each(["", " ", "  ", "メールアドレス", "&lt;&copy;&amp;", "㌶Ⅲ⑳㏾㈱髙﨑", "ヲンヰヱヴーヾ・", "ｧｰｭｿﾏﾞﾟ"])(
      "should fail to create an instance of AccountFullName with empty input %s",
      (actual) => {
        const result = AccountFullName.create(actual)
        expect(result.isErr).toBe(true)
        result.mapErr((e) => expect(e).instanceOf(DomainError))
      }
    )
  })
})
