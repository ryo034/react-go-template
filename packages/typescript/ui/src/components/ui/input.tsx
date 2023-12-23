import { InputHTMLAttributes, forwardRef } from "react"
import { UseFormRegisterReturn } from "react-hook-form"
import { cn } from "~/infrastructure/tailwindcss"

export interface InputProps extends InputHTMLAttributes<HTMLInputElement> {
  id?: string
  fullWidth?: boolean
  reactHookForm?: UseFormRegisterReturn
}

export const inputBaseClass =
  "flex h-10 w-full rounded-md border border-input bg-transparent px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"

const Input = forwardRef<HTMLInputElement, InputProps>(
  ({ className, type, id, fullWidth, reactHookForm, ...props }, ref) => {
    return (
      <input
        id={id}
        data-testid={id}
        type={type}
        className={cn(inputBaseClass, className, fullWidth ? "w-full" : "")}
        ref={ref}
        {...props}
        {...reactHookForm}
      />
    )
  }
)
Input.displayName = "Input"

export { Input }
