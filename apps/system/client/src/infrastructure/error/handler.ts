import { FirebaseError } from "firebase/app"
import { DomainError } from "~/domain"
import { AdapterError, FirebaseErrorAdapter } from "~/infrastructure/error"

export class UnknownError extends Error {}

export class ErrorHandler extends Error {
  static adapt(err: unknown): Error {
    // ここでログが出力されてない場合はadapterのエラーの可能性が高い
    console.error(err)

    let error: Error
    if (err instanceof FirebaseError) {
      error = FirebaseErrorAdapter.create(err)
    } else if (err instanceof DomainError) {
      error = err
    } else if (err instanceof Error) {
      error = err
    } else if (err instanceof AdapterError) {
      error = err
    } else if (typeof err === "string") {
      error = new UnknownError(err)
    } else {
      error = new UnknownError("unknown error")
    }
    return error
  }
}
