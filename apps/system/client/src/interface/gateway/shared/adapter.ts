import { Result } from "true-myth"
import { AccountId, AccountName, Email, User } from "~/domain"
import {
  Address,
  AddressBuilding,
  AddressCity,
  AddressCountry,
  AddressPrefecture,
  AddressStreet,
  AddressZipCode
} from "~/domain/shared/address"
import * as SharedPb from "~/generated/schema/api/shared/v1/shared_pb"
import { User as UserPb } from "~/generated/schema/api/user/v1/user_pb"
import { AdapterError } from "~/infrastructure/error"

export class SharedGatewayAdapter {
  adaptAddress(v: SharedPb.Address): Result<Address, Error> {
    if (v.zipCode === null || v.zipCode === undefined) {
      return Result.err(new AdapterError(SharedGatewayAdapter.name, this.adaptAddress.name, "v.zipCode is required"))
    }

    const zipCode = AddressZipCode.create(v.zipCode)
    if (zipCode.isErr) {
      return Result.err(zipCode.error)
    }

    if (v.country === null || v.country === undefined) {
      return Result.err(new AdapterError(SharedGatewayAdapter.name, this.adaptAddress.name, "v.country is required"))
    }

    const country = AddressCountry.create(v.country)
    if (country.isErr) {
      return Result.err(country.error)
    }

    if (v.prefecture === null || v.prefecture === undefined) {
      return Result.err(new AdapterError(SharedGatewayAdapter.name, this.adaptAddress.name, "v.prefecture is required"))
    }

    const prefecture = AddressPrefecture.create(v.prefecture)
    if (prefecture.isErr) {
      return Result.err(prefecture.error)
    }

    if (v.city === null || v.city === undefined) {
      return Result.err(new AdapterError(SharedGatewayAdapter.name, this.adaptAddress.name, "v.city is required"))
    }

    const city = AddressCity.create(v.city)
    if (city.isErr) {
      return Result.err(city.error)
    }

    if (v.street === null || v.street === undefined) {
      return Result.err(new AdapterError(SharedGatewayAdapter.name, this.adaptAddress.name, "v.street is required"))
    }

    const street = AddressStreet.create(v.street)
    if (street.isErr) {
      return Result.err(street.error)
    }

    if (v.building === null || v.building === undefined) {
      return Result.err(new AdapterError(SharedGatewayAdapter.name, this.adaptAddress.name, "v.building is required"))
    }

    const building = AddressBuilding.create(v.building)
    if (building.isErr) {
      return Result.err(building.error)
    }
    return Result.ok(
      Address.create({
        zipCode: zipCode.value,
        country: country.value,
        prefecture: prefecture.value,
        city: city.value,
        street: street.value,
        building: building.value
      })
    )
  }

  adaptUser(v: UserPb): Result<User, Error> {
    if (v.userId === null || v.userId === undefined) {
      return Result.err(new AdapterError(SharedGatewayAdapter.name, this.adaptUser.name, "v.userId is required"))
    }
    const id = AccountId.fromString(v.userId)
    if (id.isErr) {
      return Result.err(id.error)
    }

    if (v.firstName === null || v.firstName === undefined) {
      return Result.err(new AdapterError(SharedGatewayAdapter.name, this.adaptUser.name, "v.firstName is required"))
    }
    const firstName = AccountName.create(v.firstName)
    if (firstName.isErr) {
      return Result.err(firstName.error)
    }

    if (v.lastName === null || v.lastName === undefined) {
      return Result.err(new AdapterError(SharedGatewayAdapter.name, this.adaptUser.name, "v.lastName is required"))
    }
    const lastName = AccountName.create(v.lastName)
    if (lastName.isErr) {
      return Result.err(lastName.error)
    }

    if (v.email === null || v.email === undefined) {
      return Result.err(new AdapterError(SharedGatewayAdapter.name, this.adaptUser.name, "v.email is required"))
    }
    const email = Email.create(v.email)
    if (email.isErr) {
      return Result.err(email.error)
    }
    return Result.ok(
      User.create({
        id: id.value,
        firstName: firstName.value,
        lastName: lastName.value,
        email: email.value
      })
    )
  }
}
