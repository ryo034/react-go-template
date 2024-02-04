import { readFileSync } from "fs"
import { getApp, getApps, initializeApp } from "firebase-admin/app"
import { MultiFactorCreateSettings, UserProvider, getAuth } from "firebase-admin/auth"
import * as fba from "firebase/app"
import * as fb from "firebase/auth"
import { firebaseClientConfig } from "./config"

export interface FirebaseUser {
  localId: string
  email: string
  emailVerified: boolean
  mfaInfo: { phoneInfo: string }[]
  phoneNumber: string
  passwordHash: string
  displayName: string
  photoURL: string
  disabled: boolean
  providerUserInfo: UserProvider[]
  multiFactor: MultiFactorCreateSettings | null
}

export interface FirebaseConfig {
  apiKey: string
  authDomain: string
  projectId: string
  storageBucket: string
  messagingSenderId: string
  appId: string
  localConfig?: {
    firebaseEmulatorHost: string
    firestoreEmulatorHost: string
  }
}

export interface FirebaseTestConfig {
  showConsole: boolean
}

const firebaseAuthEmulatorHost = "http://localhost:9099"

export class Firebase {
  private firebaseAdminAuth: ReturnType<typeof getAuth>
  private firebaseClientAuth: ReturnType<typeof fb.getAuth>

  constructor(
    private readonly config: FirebaseConfig,
    private readonly testConfig: FirebaseTestConfig = { showConsole: false }
  ) {
    if (this.config.localConfig) {
      if (testConfig.showConsole) console.log("Setting up Firebase Emulator...")
      process.env.FIREBASE_AUTH_EMULATOR_HOST = this.config.localConfig.firebaseEmulatorHost
      process.env.FIRESTORE_EMULATOR_HOST = this.config.localConfig.firestoreEmulatorHost
    }
    const firebase = getApps().length === 0 ? initializeApp(config) : getApp()
    this.firebaseAdminAuth = getAuth(firebase)

    // Firebase Client
    const fbc = fba.initializeApp(firebaseClientConfig)
    const firebaseAuth = fb.getAuth(fbc)
    if (this.config.localConfig) {
      fb.connectAuthEmulator(firebaseAuth, firebaseAuthEmulatorHost, { disableWarnings: true })
    }
    this.firebaseClientAuth = firebaseAuth
  }

  async clear() {
    if (this.testConfig.showConsole) console.log("Clearing firebase...")
    const users = await this.firebaseAdminAuth.listUsers()
    for (const u of users.users) {
      await this.firebaseAdminAuth.revokeRefreshTokens(u.uid)
    }
    await this.firebaseAdminAuth.deleteUsers(users.users.map((u: any) => u.uid))
  }

  async signInWithCustomToken(token: string) {
    const res = await fb.signInWithCustomToken(this.firebaseClientAuth, token)
    return await res.user.getIdTokenResult()
  }

  get authInstance() {
    return this.firebaseAdminAuth
  }

  async setup() {
    const jsonData = JSON.parse(readFileSync("./setup/firebase/auth/users.json", "utf8").toString())
    const users = jsonData.users as FirebaseUser[]
    if (this.testConfig.showConsole) console.log("Setting up firebase...")
    for (let idx = 0; idx < users.length; idx++) {
      let mfa: MultiFactorCreateSettings | undefined
      if (users[idx].mfaInfo) {
        mfa = {
          enrolledFactors: []
        }
        for (const mfaInfo of users[idx].mfaInfo) {
          mfa.enrolledFactors.push({
            displayName: "",
            factorId: "phone",
            phoneNumber: mfaInfo.phoneInfo
          })
        }
      }
      // const password = users[idx].passwordHash.split(":")[2].split("=")[1]
      try {
        await this.firebaseAdminAuth.createUser({
          uid: users[idx].localId,
          email: users[idx].email,
          emailVerified: users[idx].emailVerified,
          phoneNumber: users[idx].phoneNumber,
          // password,
          displayName: users[idx].displayName,
          photoURL: users[idx].photoURL,
          disabled: users[idx].disabled,
          multiFactor: mfa
        })
      } catch (e) {
        console.error("failed to create user", users[idx], e)
      }
    }
  }
}