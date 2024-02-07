import { Loader2 } from "lucide-react"

export const Loading = () => {
  return (
    <div className="fixed top-0 left-0 right-0 bottom-0 w-full h-screen z-50 overflow-hidden bg-gray-700 opacity-10 flex flex-col items-center justify-center">
      <Loader2 className="animate-spin" size={54} color="#fff" />
    </div>
  )
}
