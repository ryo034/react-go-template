import { APIRequestContext } from "@playwright/test"
import { getApp, getApps, initializeApp } from "firebase-admin/app"
import { MultiFactorCreateSettings, UserProvider, getAuth } from "firebase-admin/auth"

const firebaseApiHost = process.env.FIREBASE_API_HOST || "http://localhost:9099/identitytoolkit.googleapis.com"
const firebaseApiKey = process.env.FIREBASE_API_KEY || "test"

interface FirebaseResult {
  kind: string
  registered: boolean
  localId: string
  email: string
  idToken: string
  refreshToken: string
  expiresIn: string
}

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

interface FirebaseConfig {
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

export const getFirebaseApiKey = async (
  request: APIRequestContext,
  email: string,
  password: string
): Promise<FirebaseResult> => {
  try {
    const firebaseResult = await request.post(
      `${firebaseApiHost}/v1/accounts:signInWithPassword?key=${firebaseApiKey}`,
      {
        headers: { "Content-Type": "application/json" },
        data: { email, password, returnSecureToken: true }
      }
    )
    const res = (await firebaseResult.json()) as FirebaseResult
    if (!res.idToken) {
      throw new Error("firebaseResult.idToken is undefined")
    }
    return res
  } catch (e) {
    console.error("failed to get firebase api key", e)
  }
}

export class Firebase {
  private auth: ReturnType<typeof getAuth>
  private constructor(config: FirebaseConfig) {
    if (config.localConfig) {
      console.log("Setting up Firebase Emulator...")
      process.env.FIREBASE_AUTH_EMULATOR_HOST = config.localConfig.firebaseEmulatorHost
      process.env.FIRESTORE_EMULATOR_HOST = config.localConfig.firestoreEmulatorHost
    }
    const firebase = getApps().length === 0 ? initializeApp(config) : getApp()
    this.auth = getAuth(firebase)
  }

  static getInstance(config: FirebaseConfig) {
    return new Firebase(config)
  }

  async clear() {
    console.log("Clearing firebase...")
    const users = await this.auth.listUsers()
    for (const u of users.users) {
      await this.auth.revokeRefreshTokens(u.uid)
    }
    await this.auth.deleteUsers(users.users.map((user) => user.uid))
  }

  async setup(users: FirebaseUser[]) {
    console.log("Setting up firebase...")
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
      const password = users[idx].passwordHash.split(":")[2].split("=")[1]
      try {
        await this.auth.createUser({
          uid: users[idx].localId,
          email: users[idx].email,
          emailVerified: users[idx].emailVerified,
          phoneNumber: users[idx].phoneNumber,
          password,
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
