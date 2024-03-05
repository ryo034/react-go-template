import { readFileSync } from "fs"
import { getApp, getApps, initializeApp } from "firebase-admin/app"
import { MultiFactorCreateSettings, UserProvider, UserProviderRequest, getAuth } from "firebase-admin/auth"
import * as fba from "firebase/app"
import * as fb from "firebase/auth"
import { firebaseClientConfig } from "./config"

export interface FirebaseUser {
  localId: string
  email: string
  emailVerified: boolean
  mfaInfo: { phoneInfo: string }[]
  phoneNumber: string
  passwordHash: Buffer
  displayName: string
  photoUrl: string
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

      const providerData: UserProviderRequest[] = []

      if (users[idx].providerUserInfo && users[idx].providerUserInfo.length > 0) {
        for (const provider of users[idx].providerUserInfo) {
          if (provider.providerId === "google.com") {
            providerData.push({
              uid: users[idx].localId,
              displayName: provider.displayName,
              email: provider.email,
              phoneNumber: users[idx].phoneNumber,
              photoURL: users[idx].photoUrl,
              providerId: provider.providerId
            })
          }
        }
      }

      try {
        await this.firebaseAdminAuth.importUsers([
          {
            uid: users[idx].localId,
            email: users[idx].email,
            emailVerified: users[idx].emailVerified,
            phoneNumber: users[idx].phoneNumber,
            passwordHash: users[idx].passwordHash,
            displayName: users[idx].displayName,
            photoURL: users[idx].photoUrl,
            disabled: users[idx].disabled,
            providerData,
            multiFactor: mfa
          }
        ])
      } catch (e) {
        console.error("failed to create user", users[idx], e)
      }
    }
  }
}
