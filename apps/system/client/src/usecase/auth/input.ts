import type { Otp } from "~/domain/auth"
import type { Email } from "~/domain/shared"

export type StartWithEmailInput = {
  email: Email
}

export type VerifyOtpInput = {
  email: Email
  otp: Otp
}

export type FindInvitationByTokenInput = {
  token: string
}

export type ProceedInvitationByEmailInput = {
  token: string
  email: Email
}

export type ProceedInvitationByOAuthInput = {
  token: string
}
