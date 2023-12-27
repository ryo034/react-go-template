import { cva } from "class-variance-authority"
import { HTMLAttributes, ReactNode } from "react"
import { Link } from "react-router-dom"
import { cn } from "~/infrastructure/tailwindcss"

export interface NavItem {
  label: string
  icon: ReactNode | undefined
  link: string
}

export interface Props extends HTMLAttributes<HTMLAnchorElement> {
  menu: NavItem
  variant?: "default"
}

const variants = cva(
  "flex items-center rounded-lg px-3 py-2 text-slate-900 dark:text-white focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 dark:focus:ring-blue-500",
  {
    variants: {
      variant: {
        default: "hover:bg-sidebar-hover"
      }
    },
    defaultVariants: {
      variant: "default"
    }
  }
)

export const SidebarListItem = ({ menu, variant, className }: Props) => {
  return (
    <li className="px-3">
      <Link className={cn(variants({ variant, className }))} to={menu.link}>
        <span className="inline-flex items-center justify-center gap-1">
          {menu.icon} <span className="ml-3 text-lg flex-1 whitespace-nowrap">{menu.label}</span>
        </span>
      </Link>
    </li>
  )
}
