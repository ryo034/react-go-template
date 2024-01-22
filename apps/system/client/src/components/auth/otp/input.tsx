import { forwardRef } from "react"
import { Input, InputProps } from "shared-ui"

type OptDigitInputProps = InputProps

const OptDigitInput = forwardRef<HTMLInputElement, OptDigitInputProps>(
  ({ className, type, id, fullWidth, reactHookForm, ...props }, ref) => {
    return (
      <Input
        className="font-bold w-12 h-12 text-center text-2xl rounded-md border-gray-300 focus:border-indigo-500 focus:ring-indigo-500 [appearance:textfield] [&::-webkit-outer-spin-button]:appearance-none [&::-webkit-inner-spin-button]:appearance-none"
        maxLength={1}
        type="number"
        reactHookForm={reactHookForm}
        ref={ref}
        {...props}
      />
    )
  }
)

export { OptDigitInput }
