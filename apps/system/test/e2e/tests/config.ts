export const isHeadless = false
export const env = process.env.ENV || "localhost"
export const isLocal = env === "localhost"
export const homeURL = process.env.HOME_URL || "http://localhost:5173"
export const initializeData = process.env.DO_INITIALIZE_DATA || "false"
export const rootPath = initializeData === "true" && env !== "localhost" ? "./e2e" : "."
export const dbHost = process.env.DB_HOST || "localhost"
export const dbUser = process.env.DB_USER || "root"
export const dbPassword = process.env.DB_PASSWORD || "password"
export const dbDatabase = process.env.DB_DATABASE || "main"
export const dbPort = process.env.DB_PORT || "15432"

export const firebaseApiKey = process.env.FIREBASE_API_KEY || "test"
export const firebaseAuthDomain = process.env.FIREBASE_AUTH_DOMAIN || "localhost"
export const firebaseProjectId = process.env.FIREBASE_PROJECT_ID || "test"
export const firebaseStorageBucket = process.env.FIREBASE_STORAGE_BUCKET || "test"
export const firebaseMessagingSenderId = process.env.FIREBASE_MESSAGING_SENDER_ID || "test"
export const firebaseAppId = process.env.FIREBASE_APP_ID || "test"

export const firebaseConfig = {
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
}

export const firebaseClientConfig = {
  apiKey: process.env.FIREBASE_API_KEY || "test",
  authDomain: process.env.FIREBASE_AUTH_DOMAIN || "localhost",
  projectId: process.env.FIREBASE_PROJECT_ID || "test",
  storageBucket: process.env.FIREBASE_STORAGE_BUCKET || "test",
  messagingSenderId: process.env.FIREBASE_MESSAGING_SENDER_ID || "test",
  appId: process.env.FIREBASE_APP_ID || "test"
}

export const dbConfig = {
  host: dbHost,
  port: parseInt(dbPort),
  user: dbUser,
  password: dbPassword,
  database: dbDatabase
}
