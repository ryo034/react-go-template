import { useVariableValue } from "@devcycle/react-client-sdk"

export const useFeatureFlag = (key: string, defaultValue: boolean) => {
  return useVariableValue(key, defaultValue)
}
