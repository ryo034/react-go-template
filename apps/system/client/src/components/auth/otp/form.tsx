import { useEffect } from "react"
import { SubmitHandler, useForm } from "react-hook-form"
import { Button, FormResultErrorMessage } from "shared-ui"
import { useVerifyOtpPageFormMessage } from "~/components/auth/otp/message"
import { OptDigitInput } from "./input"

export type OtpFormValues = {
  otpInput1: string
  otpInput2: string
  otpInput3: string
  otpInput4: string
  otpInput5: string
  otpInput6: string
}

interface Props {
  onSubmit: SubmitHandler<OtpFormValues>
  errorMessage: string
}

const verifyOtpFormId = "verifyOtpForm"

export const VerifyOTPPageForm = ({ onSubmit, errorMessage }: Props) => {
  const message = useVerifyOtpPageFormMessage()

  const {
    register,
    handleSubmit,
    setFocus,
    setValue,
    watch,
    formState: { isSubmitSuccessful, isSubmitting }
  } = useForm<OtpFormValues>({
    defaultValues: {
      otpInput1: "",
      otpInput2: "",
      otpInput3: "",
      otpInput4: "",
      otpInput5: "",
      otpInput6: ""
    }
  })
  const { otpInput1, otpInput2, otpInput3, otpInput4, otpInput5, otpInput6 } = watch()

  const otpInput1Field = register("otpInput1", {
    maxLength: 1,
    onChange(e) {
      setValue("otpInput1", e.target.value.slice(-1))
      if (e.target.value.length !== 0) {
        setFocus("otpInput2")
      }
    }
  })
  const otpInput2Field = register("otpInput2", {
    maxLength: 1,
    onChange(e) {
      setValue("otpInput2", e.target.value.slice(-1))
      if (e.target.value.length !== 0) {
        setFocus("otpInput3")
      }
    }
  })
  const otpInput3Field = register("otpInput3", {
    maxLength: 1,
    onChange(e) {
      setValue("otpInput3", e.target.value.slice(-1))
      if (e.target.value.length !== 0) {
        setFocus("otpInput4")
      }
    }
  })
  const otpInput4Field = register("otpInput4", {
    maxLength: 1,
    onChange(e) {
      setValue("otpInput4", e.target.value.slice(-1))
      if (e.target.value.length !== 0) {
        setFocus("otpInput5")
      }
    }
  })
  const otpInput5Field = register("otpInput5", {
    maxLength: 1,
    onChange(e) {
      setValue("otpInput5", e.target.value.slice(-1))
      if (e.target.value.length !== 0) {
        setFocus("otpInput6")
      }
    }
  })
  const otpInput6Field = register("otpInput6", {
    maxLength: 1,
    onChange(e) {
      console.log("onChange")
      setValue("otpInput6", e.target.value.slice(-1))
    }
  })

  useEffect(() => {
    setFocus("otpInput1")
  }, [])

  useEffect(() => {
    if (isSubmitSuccessful || isSubmitting) return

    const allFilled =
      otpInput1 !== undefined &&
      otpInput1.length === 1 &&
      otpInput2 !== undefined &&
      otpInput2.length === 1 &&
      otpInput3 !== undefined &&
      otpInput3.length === 1 &&
      otpInput4 !== undefined &&
      otpInput4.length === 1 &&
      otpInput5 !== undefined &&
      otpInput5.length === 1 &&
      otpInput6 !== undefined &&
      otpInput6.length === 1

    if (allFilled) {
      handleSubmit(onSubmit)()
    }
  }, [
    otpInput1,
    otpInput2,
    otpInput3,
    otpInput4,
    otpInput5,
    otpInput6,
    isSubmitSuccessful,
    isSubmitting,
    handleSubmit,
    onSubmit
  ])

  return (
    <form className="mt-8 space-y-6" id={verifyOtpFormId} onSubmit={handleSubmit(onSubmit)}>
      <div className="flex justify-between">
        <OptDigitInput reactHookForm={otpInput1Field} data-testid="otpInput1" />
        <OptDigitInput reactHookForm={otpInput2Field} data-testid="otpInput2" />
        <OptDigitInput reactHookForm={otpInput3Field} data-testid="otpInput3" />
        <OptDigitInput reactHookForm={otpInput4Field} data-testid="otpInput4" />
        <OptDigitInput reactHookForm={otpInput5Field} data-testid="otpInput5" />
        <OptDigitInput reactHookForm={otpInput6Field} data-testid="otpInput6" />
      </div>
      <FormResultErrorMessage message={errorMessage} />
      <div>
        <Button type="submit" form={verifyOtpFormId} data-testid="verifyOtpButton" fullWidth>
          {message.word.submit}
        </Button>
      </div>
    </form>
  )
}
