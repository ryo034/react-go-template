import { Interceptor } from "@connectrpc/connect"
import { Auth } from "firebase/auth"

// can not use useContext in middleware
export class DriverAuthMiddleware {
  constructor(private readonly firebaseAuth: Auth) {}

  intercept: Interceptor = (next) => async (req) => {
    if (this.firebaseAuth.currentUser === null) {
      return await next(req)
    }
    const token = await this.firebaseAuth.currentUser.getIdToken()
    if (token) {
      req.header.append("Authorization", `Bearer ${token}`)
    }
    return await next(req)
  }
}
