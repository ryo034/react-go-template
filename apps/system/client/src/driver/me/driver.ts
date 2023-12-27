import { PromiseClient } from "@connectrpc/connect"
import { Result } from "true-myth"
import { Email, Phone } from "~/domain"
import { AccountName } from "~/domain/account"
import { MeService } from "~/generated/schema/api/me/v1/me_connect"
import {
  FindRequest,
  FindResponse,
  LoginRequest,
  LoginResponse,
  RegisterCompleteRequest,
  RegisterCompleteResponse,
  SignUpRequest,
  SignUpResponse,
  UpdateEmailRequest,
  UpdateEmailResponse,
  UpdateNameRequest,
  UpdateNameResponse,
  UpdatePhoneNumberRequest,
  UpdatePhoneNumberResponse
} from "~/generated/schema/api/me/v1/me_pb"
import { ErrorHandler } from "~/infrastructure/error/handler"
import { PromiseResult } from "~/infrastructure/shared/result"

export class MeDriver {
  constructor(private readonly client: PromiseClient<typeof MeService>) {}

  async login(): PromiseResult<LoginResponse, Error> {
    try {
      const res = await this.client.login(new LoginRequest())
      return Result.ok(res)
    } catch (e) {
      return Result.err(ErrorHandler.adapt(e))
    }
  }

  async signUp(firstName: AccountName, lastName: AccountName): PromiseResult<SignUpResponse, Error> {
    try {
      const res = await this.client.signUp(
        new SignUpRequest({
          lastName: lastName.value,
          firstName: firstName.value
        })
      )
      return Result.ok(res)
    } catch (e) {
      return Result.err(ErrorHandler.adapt(e))
    }
  }

  async verifyEmail(): PromiseResult<RegisterCompleteResponse, Error> {
    try {
      const res = await this.client.registerComplete(new RegisterCompleteRequest())
      return Result.ok(res)
    } catch (e) {
      return Result.err(ErrorHandler.adapt(e))
    }
  }

  async find(): PromiseResult<FindResponse, Error> {
    try {
      const res = await this.client.find(new FindRequest())
      return Result.ok(res)
    } catch (e) {
      return Result.err(ErrorHandler.adapt(e))
    }
  }

  async updateName(firstName: AccountName, lastName: AccountName): PromiseResult<UpdateNameResponse, Error> {
    try {
      const res = await this.client.updateName(
        new UpdateNameRequest({ firstName: firstName.value, lastName: lastName.value })
      )
      return Result.ok(res)
    } catch (e) {
      return Result.err(ErrorHandler.adapt(e))
    }
  }

  async updateEmail(email: Email): PromiseResult<UpdateEmailResponse, Error> {
    try {
      const res = await this.client.updateEmail(new UpdateEmailRequest({ email: email.value }))
      return Result.ok(res)
    } catch (e) {
      return Result.err(ErrorHandler.adapt(e))
    }
  }

  async updatePhoneNumber(phoneNumber: Phone): PromiseResult<UpdatePhoneNumberResponse, Error> {
    try {
      const res = await this.client.updatePhoneNumber(new UpdatePhoneNumberRequest({ phoneNumber: phoneNumber.value }))
      return Result.ok(res)
    } catch (e) {
      return Result.err(ErrorHandler.adapt(e))
    }
  }
}
