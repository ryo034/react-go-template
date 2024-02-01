import { BeforeSuite } from "gauge-ts"
import { AfterSuite, BeforeScenario } from "gauge-ts"
import { firebaseConfig } from "./config"
import { Databases, MainDb } from "./database"
import { Firebase } from "./firebase"

export default class SetupAndTearDown {
  // @BeforeSuite()
  // async beforeSuite() {
  // 	await Promise.all([clearFirebase(), clearDB()]);
  // 	await Promise.all([setupFirebase(), setupDB()]);
  // }
  @BeforeScenario({ tags: ["initializeDb"] })
  async beforeScenario() {
    await new MainDb().clear()
  }

  @BeforeScenario({ tags: ["initializeFirebase"] })
  async beforeScenarioFirebase() {
    const fb = new Firebase(firebaseConfig, { showConsole: false })
    await fb.clear()
    await fb.setup()
  }

  @BeforeSuite()
  @AfterSuite()
  async beforeScenarioStatefulAll() {
    const fb = new Firebase(firebaseConfig, { showConsole: false })
    const dbs = Databases.gen()
    await Promise.all([fb.clear(), dbs.clear()])
    await Promise.all([fb.setup(), dbs.setup()])
  }
}
