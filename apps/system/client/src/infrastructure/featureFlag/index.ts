export interface FeatureFlagKeyValue {
    key: string
    default: boolean
}

export const featureFlagKeyValuesList = {
    forDevelopment: {
        key: 'forDevelopment',
        default: false,
    },
    test: {
        key: 'test',
        default: false,
    },
} as const

export type FeatureFlagKeyValues = typeof featureFlagKeyValuesList
