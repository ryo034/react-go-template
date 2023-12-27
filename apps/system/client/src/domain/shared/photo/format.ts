export const PhotoFormatList = {
  png: "png",
  jpeg: "jpeg",
  gif: "gif",
  webp: "webp"
} as const

export type PhotoFormat = keyof typeof PhotoFormatList
