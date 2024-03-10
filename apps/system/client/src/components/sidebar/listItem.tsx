import type { LucideIcon } from "lucide-react"
import type { HTMLAttributes } from "react"
import { Link } from "react-router-dom"
import { buttonVariants } from "shared-ui"
import { cn } from "~/infrastructure/tailwindcss"

export interface NavItem {
  title: string
  label?: string
  icon: LucideIcon
  variant: "default" | "ghost"
  to: string
}

export interface Props extends HTMLAttributes<HTMLAnchorElement> {
  menu: NavItem
  variant?: "default"
}

export const SidebarListItem = (link: NavItem) => {
  return (
    <Link
      to={link.to}
      className={cn(
        buttonVariants({ variant: link.variant, size: "sm" }),
        link.variant === "default" && "dark:bg-muted dark:text-white dark:hover:bg-muted dark:hover:text-white",
        "justify-start"
      )}
    >
      <link.icon className="mr-2 h-4 w-4" />
      {link.title}
      {link.label && (
        <span className={cn("ml-auto", link.variant === "default" && "text-background dark:text-white")}>
          {link.label}
        </span>
      )}
    </Link>
  )
}

export const SidebarListItemCollapsed = (link: NavItem) => {
  return (
    <Link
      to={link.to}
      className={cn(
        buttonVariants({ variant: link.variant, size: "icon" }),
        "h-9 w-9",
        link.variant === "default" &&
          "dark:bg-muted dark:text-muted-foreground dark:hover:bg-muted dark:hover:text-white"
      )}
    >
      <link.icon className="h-4 w-4" />
      <span className="sr-only">{link.title}</span>
    </Link>
  )
}
