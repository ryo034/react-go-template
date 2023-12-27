import { useState } from "react"
import { CheckboxWithLabel, Input, InputProps, Label } from "shared-ui"
import { usePasswordInputComponentMessage } from "~/components/common/form/message"
import { FormErrorMessage } from "./errorMessage"

interface Props extends InputProps {
  title: string
  errorMessage: string
  fullWidth?: boolean
  showToggle?: boolean
  isCurrent?: boolean
}

export const FormPasswordInputSection = ({
  title,
  id,
  placeholder,
  reactHookForm,
  errorMessage,
  fullWidth,
  isCurrent = false,
  showToggle = true
}: Props) => {
  const message = usePasswordInputComponentMessage()
  const [isRevealPassword, setIsRevealPassword] = useState(false)
  return (
    <div className="space-y-2">
      <Label htmlFor={id} title={title}>
        {title}
      </Label>
      <Input
        fullWidth={fullWidth}
        id={id}
        type={isRevealPassword ? "text" : "password"}
        autoComplete={isCurrent ? "current-password" : "new-password"}
        placeholder={placeholder}
        reactHookForm={reactHookForm}
      />
      {showToggle && (
        <CheckboxWithLabel
          onClick={() => setIsRevealPassword(!isRevealPassword)}
          id={`${id}-togglePasswordVisibility`}
          data-testid={`${id}-togglePasswordVisibility`}
          label={message.action.showPassword}
        />
      )}
      <FormErrorMessage dataTestId={`${id}-errorMessage`} message={errorMessage} />
    </div>
  )
}
