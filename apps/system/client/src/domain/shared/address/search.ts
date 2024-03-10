import { Entity } from "~/domain/shared"
import type {
  AddressCity,
  AddressCountry,
  AddressPrefecture,
  AddressStreet,
  AddressZipCode
} from "~/domain/shared/address"

interface AddressProps {
  zipCode: AddressZipCode
  country: AddressCountry
  prefecture: AddressPrefecture
  city: AddressCity
  street: AddressStreet | null
}

export class SearchAddress extends Entity<AddressProps> {
  static create(v: AddressProps): SearchAddress {
    return new SearchAddress(v)
  }

  get zipCode(): AddressZipCode {
    return this.value.zipCode
  }

  get country(): AddressCountry {
    return this.value.country
  }

  get prefecture(): AddressPrefecture {
    return this.value.prefecture
  }

  get city(): AddressCity {
    return this.value.city
  }

  get street(): AddressStreet | null {
    return this.value.street
  }
}
