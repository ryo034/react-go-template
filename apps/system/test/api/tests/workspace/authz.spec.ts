import { expect, test } from "@playwright/test"
import { defaultPostHeaders } from "../../config/config"
import { genAPIClient } from "../../scripts"

const client = genAPIClient()

test.describe("Workspace Authz", () => {
  // // 他の組織の情報を参照した場合、404が返ること
  // test("", async () => {})
  // // メンバー権限の場合、組織の情報を参照できること
  // test("", async () => {})
  // // メンバー権限の場合、組織の情報を更新できない
  // test("", async () => {})
})
