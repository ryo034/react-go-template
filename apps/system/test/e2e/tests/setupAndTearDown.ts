import { AfterSuite, BeforeScenario, BeforeSuite } from "gauge-ts"
import { firebaseConfig } from "./config"
import { MainDb } from "./database"
import { Firebase } from "./firebase"

export default class SetupAndTearDown {
  @AfterSuite()
  @BeforeSuite()
  @BeforeScenario({ tags: ["stateful"] })
  async beforeScenarioStatefulAll() {
    const fb = new Firebase(firebaseConfig, { showConsole: false })
    const db = new MainDb()
    await Promise.all([fb.clear(), db.clear()])
    await Promise.all([fb.setup(), db.setup()])
  }
}
