import { Otp } from "~/domain/auth"
import { Email } from "~/domain/shared"

export type StartWithEmailInput = {
  email: Email
}

export type VerifyOtpInput = {
  email: Email
  otp: Otp
}
