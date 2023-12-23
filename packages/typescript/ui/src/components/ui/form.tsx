import * as LabelPrimitive from "@radix-ui/react-label"
import { Slot } from "@radix-ui/react-slot"
import { ComponentPropsWithoutRef, ElementRef, HTMLAttributes, createContext, forwardRef, useContext } from "react"
import { Controller, ControllerProps, FieldPath, FieldValues, FormProvider, useFormContext } from "react-hook-form"
import { Label } from "~/components/ui/label"
import { cn } from "~/infrastructure/tailwindcss"
import { AlertTitle, AlertWithIcon } from "./alert"

const Form = FormProvider

type FormFieldContextValue<
  TFieldValues extends FieldValues = FieldValues,
  TName extends FieldPath<TFieldValues> = FieldPath<TFieldValues>
> = {
  name: TName
}

const FormFieldContext = createContext<FormFieldContextValue>({} as FormFieldContextValue)

const FormField = <
  TFieldValues extends FieldValues = FieldValues,
  TName extends FieldPath<TFieldValues> = FieldPath<TFieldValues>
>({
  ...props
}: ControllerProps<TFieldValues, TName>) => {
  return (
    <FormFieldContext.Provider value={{ name: props.name }}>
      <Controller {...props} />
    </FormFieldContext.Provider>
  )
}

const useFormField = () => {
  const fieldContext = useContext(FormFieldContext)
  const itemContext = useContext(FormItemContext)
  const { getFieldState, formState } = useFormContext()

  const fieldState = getFieldState(fieldContext.name, formState)

  if (!fieldContext) {
    throw new Error("useFormField should be used within <FormField>")
  }

  const id = itemContext.id
  return {
    id,
    name: fieldContext.name,
    formItemId: id,
    formDescriptionId: `${id}-description`,
    formMessageId: `${id}-message`,
    ...fieldState
  }
}

type FormItemContextValue = {
  id?: string
}

const FormItemContext = createContext<FormItemContextValue>({} as FormItemContextValue)

interface FormItemProps extends HTMLAttributes<HTMLDivElement> {
  id?: string
}

const FormItem = forwardRef<HTMLDivElement, FormItemProps>(({ className, id, ...props }, ref) => {
  return (
    <FormItemContext.Provider value={{ id }}>
      <div ref={ref} className={cn("space-y-2", className)} {...props} />
    </FormItemContext.Provider>
  )
})
FormItem.displayName = "FormItem"

const FormLabel = forwardRef<
  ElementRef<typeof LabelPrimitive.Root>,
  ComponentPropsWithoutRef<typeof LabelPrimitive.Root>
>(({ className, ...props }, ref) => {
  const { error, formItemId } = useFormField()
  return <Label ref={ref} className={cn(error && "text-destructive", className)} htmlFor={formItemId} {...props} />
})
FormLabel.displayName = "FormLabel"

const FormControl = forwardRef<ElementRef<typeof Slot>, ComponentPropsWithoutRef<typeof Slot>>(({ ...props }, ref) => {
  const { error, formItemId, formDescriptionId, formMessageId } = useFormField()

  return (
    <Slot
      ref={ref}
      id={formItemId}
      aria-describedby={!error ? `${formDescriptionId}` : `${formDescriptionId} ${formMessageId}`}
      aria-invalid={!!error}
      {...props}
    />
  )
})
FormControl.displayName = "FormControl"

const FormDescription = forwardRef<HTMLParagraphElement, HTMLAttributes<HTMLParagraphElement>>(
  ({ className, ...props }, ref) => {
    const { formDescriptionId } = useFormField()

    return <p ref={ref} id={formDescriptionId} className={cn("text-sm text-muted-foreground", className)} {...props} />
  }
)
FormDescription.displayName = "FormDescription"

interface FormMessageProps extends HTMLAttributes<HTMLParagraphElement> {
  field?: string
}
const FormMessage = forwardRef<HTMLParagraphElement, FormMessageProps>(
  ({ className, children, field, ...props }, ref) => {
    const { error, formMessageId } = useFormField()
    const body = error ? String(error?.message) : children
    if (!body) {
      return null
    }
    const dataTestId = field ? `${field}-errorMessage` : ""
    return (
      <p
        ref={ref}
        id={formMessageId}
        data-testid={dataTestId}
        className={cn("text-sm font-medium text-destructive", className)}
        {...props}
      >
        {body}
      </p>
    )
  }
)
FormMessage.displayName = "FormMessage"

interface FormResultErrorMessageProps extends HTMLAttributes<HTMLDivElement> {
  message: string
}

const FormResultErrorMessage = forwardRef<HTMLDivElement, FormResultErrorMessageProps>(({ message, ...props }, ref) => {
  if (message === "") {
    return null
  }
  return (
    <AlertWithIcon data-testid="resultError" variant="destructive" ref={ref} {...props}>
      <AlertTitle>{message}</AlertTitle>
    </AlertWithIcon>
  )
})
FormMessage.displayName = "FormResultErrorMessage"

export {
  useFormField,
  Form,
  FormItem,
  FormLabel,
  FormControl,
  FormDescription,
  FormMessage,
  FormResultErrorMessage,
  FormField
}
