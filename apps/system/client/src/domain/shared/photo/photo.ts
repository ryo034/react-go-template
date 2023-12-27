import { Result } from "true-myth"
import { DomainError, Entity, PhotoFormat, domainKeys } from "~/domain/shared"

interface Props {
  url: string
  width: number
  height: number
  format: PhotoFormat
}

export class Photo extends Entity<Props> {
  static create(v: Props): Result<Photo, Error> {
    if (!v.url.startsWith("http")) {
      return Result.err(
        new DomainError({
          domainKey: domainKeys.Photo,
          value: v,
          message: "Photo url must start with http"
        })
      )
    }
    if (v.width <= 0 || v.height <= 0) {
      return Result.err(
        new DomainError({
          domainKey: domainKeys.Photo,
          value: v,
          message: "Photo width and height must be greater than 0"
        })
      )
    }
    return Result.ok(new Photo(v))
  }

  get url(): string {
    return this.value.url
  }

  get width(): number {
    return this.value.width
  }

  get height(): number {
    return this.value.height
  }

  get format(): PhotoFormat {
    return this.value.format
  }
}
