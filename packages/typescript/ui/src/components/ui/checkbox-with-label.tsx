import type { Root } from "@radix-ui/react-checkbox"
import { type ComponentPropsWithoutRef, type ElementRef, forwardRef } from "react"
import { Checkbox } from "~/components/ui/checkbox"

interface CheckboxWithLabelProps extends ComponentPropsWithoutRef<typeof Root> {
  label: string
}

const CheckboxWithLabel = forwardRef<ElementRef<typeof Root>, CheckboxWithLabelProps>(
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

export { CheckboxWithLabel }
