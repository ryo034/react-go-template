export const env = process.env.ENV || "localhost"
export const isLocal = env === "localhost"
export const homeURL = process.env.HOME_URL || "http://localhost:5173"
export const initializeData = process.env.DO_INITIALIZE_DATA || "false"
export const rootPath = initializeData === "true" && env !== "localhost" ? "./e2e" : "."
export const isSilent = process.env.SILENT === "true" || false

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
  host: process.env.DB_HOST || "localhost",
  port: Number.parseInt(process.env.DB_PORT || "15432"),
  user: process.env.DB_USER || "root",
  password: process.env.DB_PASSWORD || "password",
  database: process.env.DB_DATABASE || "main"
}
