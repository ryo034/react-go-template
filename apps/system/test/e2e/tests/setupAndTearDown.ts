import { BeforeScenario, BeforeSuite } from "gauge-ts"
import { AfterSuite } from "gauge-ts"
import { firebaseConfig } from "./config"
import { MainDb } from "./database"
import { Firebase } from "./firebase"

export default class SetupAndTearDown {
  // @BeforeSuite()
  // async beforeSuite() {
  // 	await Promise.all([clearFirebase(), clearDB()]);
  // 	await Promise.all([setupFirebase(), setupDB()]);
  // }

  // @BeforeScenario({ tags: ["stateful"] })
  // @AfterSuite()
  @BeforeSuite()
  async beforeScenarioStatefulAll() {
    const fb = new Firebase(firebaseConfig, { showConsole: false })
    const db = new MainDb()
    await Promise.all([fb.clear(), db.clear()])
    await Promise.all([fb.setup(), db.setup()])
  }
}
