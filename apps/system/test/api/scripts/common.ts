import { readFileSync } from "fs"
import createClient from "openapi-fetch"
import { paths } from "../schema/openapi/systemApi"
import { MainDb } from "./database"
import { Firebase, FirebaseUser } from "./firebase"

export const statefulBeforeEach = async () => {
  const fb = Firebase.getInstance({
    apiKey: process.env.FIREBASE_API_KEY || "test",
    authDomain: process.env.FIREBASE_AUTH_DOMAIN || "localhost",
    projectId: process.env.FIREBASE_PROJECT_ID || "test",
    storageBucket: process.env.FIREBASE_STORAGE_BUCKET || "test",
    messagingSenderId: process.env.FIREBASE_MESSAGING_SENDER_ID || "test",
    appId: process.env.FIREBASE_APP_ID || "test",
    localConfig: {
      firebaseEmulatorHost: process.env.FIREBASE_EMULATOR_HOST || "localhost:9099",
      firestoreEmulatorHost: process.env.FIRESTORE_EMULATOR_HOST || "localhost:8080"
    }
  })
  const jsonData = JSON.parse(readFileSync("./setup/firebase/auth/users.json", "utf8").toString())
  const us = jsonData.users as FirebaseUser[]
  const db = new MainDb()
  await Promise.all([fb.clear(), db.clear()])
  await Promise.all([fb.setup(us), db.setup()])
}

export const genAPIClient = () => {
  return createClient<paths>({ baseUrl: "http://localhost:19004" })
}
