import ReactDOM from "react-dom/client"
import "../../../../packages/typescript/ui/src/styles/index.css"
import { FeatureFlagProvider } from "./app"
import { initI18n } from "./infrastructure/i18n/i18n"

initI18n()

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(<FeatureFlagProvider />)
