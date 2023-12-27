const shallowCompare = (obj1: any, obj2: any) =>
  Object.keys(obj1).length === Object.keys(obj2).length && Object.keys(obj1).every((key) => obj1[key] === obj2[key])

export abstract class AbstractEntity<T> {
  protected constructor(protected readonly value: T) {
    this.value = Object.freeze(value)
  }

  equals(vo?: AbstractEntity<T>): boolean {
    if (vo == null) {
      return false
    }
    return shallowCompare(this.value, vo.value)
  }
}

export default AbstractEntity
