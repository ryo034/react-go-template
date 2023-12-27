export abstract class ValueObject<T> {
  constructor(readonly value: T) {}

  public eq<Instance extends ValueObject<T>>(this: Instance, vo: Instance): boolean {
    return vo.value === this.value
  }
}
