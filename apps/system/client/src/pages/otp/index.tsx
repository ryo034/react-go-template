import { useContext, useMemo, useState } from "react"
import { SubmitHandler } from "react-hook-form"
import { useNavigate, useSearchParams } from "react-router-dom"
import { OtpFormValues, VerifyOTPPageForm } from "~/components/auth/otp/form"
import { Email } from "~/domain"
import { i18nKeys } from "~/infrastructure/i18n"
import { ContainerContext } from "~/infrastructure/injector/context"
import { routeMap } from "~/infrastructure/route/path"

export const verifyOtpPageRoute = "/verify-otp"

export const VerifyOtpPage = () => {
  const { controller, i18n, errorMessageProvider } = useContext(ContainerContext)
  const [searchParams] = useSearchParams()
  const [errorMessage, setErrorMessage] = useState("")
  const navigate = useNavigate()
  const email = searchParams.get("email")
  const ema = Email.create(email ?? "")

  useMemo(() => {
    if (ema.isErr) {
      navigate(routeMap.home)
      return
    }
  }, [email])

  const onSubmit: SubmitHandler<OtpFormValues> = async (d) => {
    if (ema.isErr) {
      setErrorMessage(errorMessageProvider.resolve(ema.error))
      return
    }
    const opt = `${d.otpInput1}${d.otpInput2}${d.otpInput3}${d.otpInput4}${d.otpInput5}${d.otpInput6}`
    const res = await controller.auth.verifyOtp(ema.value, opt)
    if (res) {
      setErrorMessage(errorMessageProvider.resolve(res))
      return
    }
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
            <span>{email}</span>
          </p>
          <VerifyOTPPageForm onSubmit={onSubmit} errorMessage={errorMessage} />
        </div>
      </section>
    </div>
  )
}
