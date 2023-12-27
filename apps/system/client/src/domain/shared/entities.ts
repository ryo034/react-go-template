export abstract class Entities<T> {
  protected props: T[]

  constructor(props: T[]) {
    this.props = props
  }

  get values(): T[] {
    return this.props
  }

  get size(): number {
    return this.props.length
  }

  get isEmpty(): boolean {
    return this.size === 0
  }

  get isNotEmpty(): boolean {
    return !this.isEmpty
  }

  get first(): T | null | undefined {
    return this.isEmpty ? null : this.values[0]
  }

  get last(): T | null | undefined {
    return this.isEmpty ? null : this.values[this.size - 1]
  }
}

export default Entities
