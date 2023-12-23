import * as LabelPrimitive from "@radix-ui/react-label"
import { type VariantProps, cva } from "class-variance-authority"
import { ComponentPropsWithoutRef, ElementRef, forwardRef } from "react"
import { cn } from "~/infrastructure/tailwindcss"

const variants = cva(
  "text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70 dark:text-white"
)

const Label = forwardRef<
  ElementRef<typeof LabelPrimitive.Root>,
  ComponentPropsWithoutRef<typeof LabelPrimitive.Root> & VariantProps<typeof variants>
>(({ className, ...props }, ref) => <LabelPrimitive.Root ref={ref} className={cn(variants(), className)} {...props} />)
Label.displayName = LabelPrimitive.Root.displayName

export { Label }
