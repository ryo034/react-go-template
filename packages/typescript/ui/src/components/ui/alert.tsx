import { type VariantProps, cva } from "class-variance-authority"
import { type HTMLAttributes, type ReactNode, forwardRef } from "react"
import { HiExclamationCircle, HiExclamationTriangle, HiInformationCircle } from "react-icons/hi2"
import { cn } from "~/infrastructure/tailwindcss"

// export const baseAlertClass = "relative w-full rounded-lg border p-4 [&>svg]:absolute [&>svg]:text-foreground [&>svg]:left-4 [&>svg]:top-4 [&>svg+div]:translate-y-[-3px] [&:has(svg)]:pl-11"
export const baseAlertClass =
  "relative w-full rounded-lg border p-4 [&>svg]:absolute [&>svg]:left-4 [&>svg]:top-4 [&>svg+div]:translate-y-[-3px] [&:has(svg)]:pl-11"

const variants = cva(baseAlertClass, {
  variants: {
    variant: {
      default:
        "border-gray-300 text-gray-800 bg-gray-50 [&>svg]:text-gray-800 dark:text-white dark:border-gray-400 dark:bg-transparent dark:[&>svg]:text-white",
      destructive:
        "border-red-300 text-red-800 bg-red-50 [&>svg]:text-red-800 dark:text-white dark:border-red-800 dark:bg-transparent dark:[&>svg]:text-white",
      warning:
        "border-yellow-300 text-yellow-800 bg-yellow-50 [&>svg]:text-yellow-800 dark:text-white dark:border-yellow-800 dark:bg-transparent dark:[&>svg]:text-white"
    }
  },
  defaultVariants: {
    variant: "default"
  }
})

const iconSize = 20

export const AlertCategoryIconMap: Record<string, ReactNode> = {
  default: <HiInformationCircle size={iconSize} />,
  destructive: <HiExclamationCircle size={iconSize} />,
  warning: <HiExclamationTriangle size={iconSize} />
} as const

const Alert = forwardRef<HTMLDivElement, HTMLAttributes<HTMLDivElement> & VariantProps<typeof variants>>(
  ({ className, variant, ...props }, ref) => (
    <div ref={ref} role="alert" className={cn(variants({ variant }), className)} {...props} />
  )
)
Alert.displayName = "Alert"

const AlertWithIcon = forwardRef<HTMLDivElement, HTMLAttributes<HTMLDivElement> & VariantProps<typeof variants>>(
  ({ className, variant, children, ...props }, ref) => (
    <Alert ref={ref} className={cn(variants({ variant }), className)} {...props}>
      {variant && AlertCategoryIconMap[variant]}
      {children}
    </Alert>
  )
)
Alert.displayName = "AlertWithIcon"

const AlertTitle = forwardRef<HTMLParagraphElement, HTMLAttributes<HTMLHeadingElement>>(
  ({ className, ...props }, ref) => (
    <p ref={ref} className={cn("mb-1 font-medium leading-none tracking-tight", className)} {...props} />
  )
)
AlertTitle.displayName = "AlertTitle"

const AlertDescription = forwardRef<HTMLParagraphElement, HTMLAttributes<HTMLParagraphElement>>(
  ({ className, ...props }, ref) => (
    <div ref={ref} className={cn("text-sm [&_p]:leading-relaxed", className)} {...props} />
  )
)
AlertDescription.displayName = "AlertDescription"

export { Alert, AlertWithIcon, AlertTitle, AlertDescription }
