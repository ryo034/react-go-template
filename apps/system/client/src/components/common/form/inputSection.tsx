import { ReactNode } from "react"
import { Input, InputProps, Label } from "shared-ui"
import { FormErrorMessage } from "~/components/common/form/errorMessage"

interface Props extends InputProps {
  showLabel?: boolean
  fullWidth?: boolean
  title: string
  errorMessage: string
  autoComplete?: string
  prefixElm?: ReactNode
  suffixElm?: ReactNode
  customClass?: string[]
  rootClass?: string[]
}

export const FormInputSection = ({
  showLabel = true,
  fullWidth = false,
  title,
  id,
  type,
  placeholder,
  autoComplete,
  prefixElm,
  suffixElm,
  reactHookForm,
  errorMessage,
  customClass,
  rootClass
}: Props) => {
  return (
    <div className={`${showLabel ? "space-y-2" : ""} ${rootClass ? rootClass.join(" ") : ""}`.trim() || undefined}>
      <Label htmlFor={id} title={title} className={showLabel ? "" : "sr-only"}>
        {title}
      </Label>
      {!(suffixElm && prefixElm) ? (
        <Input
          className={customClass ? customClass.join(" ") : ""}
          autoComplete={autoComplete}
          fullWidth={fullWidth}
          id={id}
          type={type}
          placeholder={placeholder}
          reactHookForm={reactHookForm}
        />
      ) : (
        <div className="flex space-x-4 items-center relative">
          {prefixElm}
          <Input
            className={customClass ? customClass.join(" ") : ""}
            fullWidth={fullWidth}
            autoComplete={autoComplete}
            id={id}
            type={type}
            placeholder={placeholder}
            reactHookForm={reactHookForm}
          />
          {suffixElm}
        </div>
      )}
      <FormErrorMessage dataTestId={`${id}-errorMessage`} message={errorMessage} />
    </div>
  )
}
