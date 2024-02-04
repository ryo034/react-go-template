import { BeforeScenario, BeforeSuite } from "gauge-ts"
import { AfterSuite } from "gauge-ts"
import { firebaseConfig } from "./config"
import { Databases } from "./database"
import { Firebase } from "./firebase"

export default class SetupAndTearDown {
  // @BeforeSuite()
  // async beforeSuite() {
  // 	await Promise.all([clearFirebase(), clearDB()]);
  // 	await Promise.all([setupFirebase(), setupDB()]);
  // }

  @BeforeSuite()
  @AfterSuite()
  @BeforeScenario({ tags: ["stateful"] })
  async beforeScenarioStatefulAll() {
    const fb = new Firebase(firebaseConfig, { showConsole: false })
    const dbs = Databases.gen()
    await Promise.all([fb.clear(), dbs.clear()])
    await Promise.all([fb.setup(), dbs.setup()])
  }
}
