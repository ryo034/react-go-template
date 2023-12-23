import * as CheckboxPrimitive from "@radix-ui/react-checkbox"
import { Check } from "lucide-react"
import { ComponentPropsWithoutRef, ElementRef, forwardRef } from "react"

import { cn } from "~/infrastructure/tailwindcss"

const Checkbox = forwardRef<
  ElementRef<typeof CheckboxPrimitive.Root>,
  ComponentPropsWithoutRef<typeof CheckboxPrimitive.Root>
>(({ className, ...props }, ref) => (
  <CheckboxPrimitive.Root
    ref={ref}
    className={cn(
      "peer h-4 w-4 shrink-0 rounded-sm border border-primary ring-offset-background focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 data-[state=checked]:bg-primary data-[state=checked]:text-primary-foreground dark:focus:ring-blue-600 dark:ring-offset-gray-800",
      className
    )}
    {...props}
  >
    <CheckboxPrimitive.Indicator className={cn("flex items-center justify-center text-current")}>
      <Check className="h-4 w-4" />
    </CheckboxPrimitive.Indicator>
  </CheckboxPrimitive.Root>
))
Checkbox.displayName = CheckboxPrimitive.Root.displayName

interface CheckboxWithLabelProps extends ComponentPropsWithoutRef<typeof CheckboxPrimitive.Root> {
  label: string
}

const CheckboxWithLabel = forwardRef<ElementRef<typeof CheckboxPrimitive.Root>, CheckboxWithLabelProps>(
  ({ id, className, label, ...props }, ref) => (
    <div className="flex items-center space-x-2 mt-2">
      <Checkbox id={id} className={className} {...props} ref={ref} />
      <label
        htmlFor={id}
        className="text-sm font-medium text-gray-900 leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70 dark:text-gray-300"
      >
        {label}
      </label>
    </div>
  )
)
CheckboxWithLabel.displayName = "CheckboxWithLabel"

export { Checkbox, CheckboxWithLabel }
