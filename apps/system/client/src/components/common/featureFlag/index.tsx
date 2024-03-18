import { featureFlagKeyValuesList } from "~/infrastructure/featureFlag"
import { useFeatureFlag } from "~/infrastructure/hooks/featureFlag"

interface Props {
  feature: keyof typeof featureFlagKeyValuesList
  on: React.ReactNode
  off: React.ReactNode
}

export const FeatureFlag = ({ feature, on, off }: Props) => {
  const featureEnabled = useFeatureFlag(
    featureFlagKeyValuesList[feature].key,
    featureFlagKeyValuesList[feature].default
  )
  return featureEnabled ? on : off
}
