import type { PromiseResult } from "~/infrastructure/shared/result"
import type { SearchAddress } from "./search"
import type { AddressZipCode } from "./zipCode"

export interface AddressRepository {
  searchByZipCode(zipCode: AddressZipCode): PromiseResult<SearchAddress, Error>
}
