import { Slot } from "@radix-ui/react-slot"
import { type VariantProps, cva } from "class-variance-authority"
import { Loader2 } from "lucide-react"
import { ButtonHTMLAttributes, forwardRef } from "react"
import { cn } from "~/infrastructure/tailwindcss"

export const buttonBaseClass =
  "rounded-md text-sm font-medium ring-offset-background transition-colors disabled:pointer-events-none disabled:opacity-50 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 dark:focus:ring-blue-500"

const buttonVariants = cva(buttonBaseClass, {
  variants: {
    variant: {
      default: "bg-primary text-primary-foreground hover:bg-primary/90",
      destructive: "bg-destructive text-destructive-foreground hover:bg-destructive/90",
      outline: "border border-input bg-background hover:bg-accent hover:text-accent-foreground",
      secondary: "bg-secondary text-secondary-foreground hover:bg-secondary/80",
      ghost: "hover:bg-accent hover:text-accent-foreground",
      link: "text-primary underline-offset-4 hover:underline"
    },
    size: {
      default: "h-10 px-4 py-2",
      sm: "h-9 rounded-md px-3",
      lg: "h-11 rounded-md px-8",
      icon: "h-10 w-10"
    },
    fullWidth: {
      true: "w-full",
      false: ""
    },
    align: {
      flexLeft: "flex items-start justify-start",
      flexCenter: "flex items-center justify-center",
      flexRight: "flex items-end justify-end",
      flexBetween: "flex items-center justify-between",
      inlineLeft: "inline-flex items-start justify-start",
      inlineCenter: "inline-flex items-center justify-center",
      inlineRight: "inline-flex items-end justify-end",
      inlineBetween: "inline-flex items-center justify-between"
    }
  },
  defaultVariants: {
    variant: "default",
    size: "default",
    fullWidth: false,
    align: "flexCenter"
  }
})

export interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement>, VariantProps<typeof buttonVariants> {
  asChild?: boolean
}

export interface LoadingButtonProps extends ButtonProps {
  text?: string
  dataTestId?: string
}

const Button = forwardRef<HTMLButtonElement, ButtonProps>(
  ({ className, fullWidth, variant, size, asChild = false, ...props }, ref) => {
    const Comp = asChild ? Slot : "button"
    return <Comp className={cn(buttonVariants({ variant, size, className, fullWidth }))} ref={ref} {...props} />
  }
)
Button.displayName = "Button"

const LoadingButton = forwardRef<HTMLButtonElement, LoadingButtonProps>(
  ({ className, fullWidth, text = "Loading...", dataTestId = "loadingButton", ...props }, ref) => {
    return (
      <Button
        disabled
        ref={ref}
        {...props}
        className={cn(buttonVariants({ className, fullWidth }))}
        data-testid={dataTestId}
      >
        <Loader2 className="mr-2 h-4 w-4 animate-spin" />
        {text}
      </Button>
    )
  }
)

export { Button, LoadingButton, buttonVariants }
