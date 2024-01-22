import { useContext, useLayoutEffect, useRef, useState } from "react"
import { SubmitHandler } from "react-hook-form"
import { useNavigate } from "react-router-dom"
import { OtpFormValues, VerifyOTPPageForm } from "~/components/auth/otp/form"
import { i18nKeys } from "~/infrastructure/i18n"
import { ContainerContext } from "~/infrastructure/injector/context"

export const VerifyOtpPage = () => {
  const { store, controller, i18n, errorMessageProvider } = useContext(ContainerContext)
  const me = store.me((state) => state.me)
  const meRef = useRef(me)
  const [errorMessage, setErrorMessage] = useState("")

  const navigate = useNavigate()

  useLayoutEffect(() => {
    store.me.subscribe((state) => {
      meRef.current = state.me
    })
  }, [])

  const onSubmit: SubmitHandler<OtpFormValues> = async (d) => {
    console.log("onSubmit")
  }

  return (
    <div className="min-h-screen bg-gray-50 dark:bg-background/95">
      <section className="w-full max-w-md mx-auto py-12 px-4 sm:px-6 lg:px-8">
        <div className="space-y-6">
          <h2 className="text-center text-3xl font-extrabold text-gray-900">
            {i18n.translate(i18nKeys.action.enter, { field: i18n.translate(i18nKeys.word.otp) })}
          </h2>
          <p className="text-center text-sm text-muted-foreground">
            {i18n.translate(i18nKeys.page.verifyOtp.enterOtpMessage)}
          </p>
          <VerifyOTPPageForm onSubmit={onSubmit} errorMessage={errorMessage} />
        </div>
      </section>
    </div>
  )
}