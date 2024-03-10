import type { StringId } from "~/domain/shared"

const keys = {
  selectedStoreId: "selectedStoreId"
}
export class LocalStorageDriver {
  constructor(private readonly storage: Storage) {}

  private setItem(key: string, v: string): void {
    this.storage.setItem(key, v)
  }

  private getItem(key: string): string | null {
    return this.storage.getItem(key)
  }

  private removeItem(key: string): void {
    this.storage.removeItem(key)
  }

  findSelectedStoreId(): string {
    return this.getItem(keys.selectedStoreId) ?? ""
  }

  setSelectedStoreId(storeId: StringId): void {
    this.setItem(keys.selectedStoreId, storeId.asString)
  }
}
