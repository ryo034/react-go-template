import { PromiseResult } from "~/infrastructure/shared/result"
import { SearchAddress } from "./search"
import { AddressZipCode } from "./zipCode"

export interface AddressRepository {
  searchByZipCode(zipCode: AddressZipCode): PromiseResult<SearchAddress, Error>
}
