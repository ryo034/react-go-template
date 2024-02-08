export const domainKeys = {
  Otp: "Otp",
  Jwt: "Jwt",
  AppDateTime: "AppDateTime",
  WorkspaceSubdomain: "WorkspaceSubdomain",
  Photo: "Photo",
  PhotoPath: "PhotoPath",
  Password: "Password",
  Phone: "Phone",
  Email: "Email",
  AddressZipCode: "AddressZipCode",
  AddressStreet: "AddressStreet",
  AddressCity: "AddressCity",
  AddressCountry: "AddressCountry",
  AddressBuilding: "AddressBuilding",
  AddressPrefecture: "AddressPrefecture",
  BusinessEntityName: "BusinessEntityName",
  BusinessEntity: "BusinessEntity",
  MemberDisplayName: "MemberDisplayName",
  MemberIdNumber: "MemberIdNumber",
  AccountName: "AccountName",
  AccountFullName: "AccountFullName",
  StringId: "StringId",
  MFAAuthCode: "MFAAuthCode"
} as const

type DomainKeyTypes = keyof typeof domainKeys

interface DomainErrorProps<T> {
  domainKey: DomainKeyTypes
  value: T
  message?: string
}

export class DomainError<T> extends Error {
  domainKey: DomainKeyTypes
  value: T
  constructor(v: DomainErrorProps<T>) {
    const msg = v.message ?? `Invalid ${v.domainKey} value: ${v.value}`
    super(msg)
    this.name = "DomainError"
    this.domainKey = v.domainKey
    this.value = v.value
  }
}
