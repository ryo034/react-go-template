// @generated by protoc-gen-es v1.6.0 with parameter "target=ts"
// @generated from file user/v1/user.proto (package user.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message user.v1.User
 */
export class User extends Message<User> {
  /**
   * @generated from field: string user_id = 1;
   */
  userId = "";

  /**
   * @generated from field: string email = 2;
   */
  email = "";

  /**
   * @generated from field: string first_name = 3;
   */
  firstName = "";

  /**
   * @generated from field: string last_name = 4;
   */
  lastName = "";

  /**
   * @generated from field: optional string phone = 5;
   */
  phone?: string;

  constructor(data?: PartialMessage<User>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "user.v1.User";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "email", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "first_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "last_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "phone", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): User {
    return new User().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): User {
    return new User().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): User {
    return new User().fromJsonString(jsonString, options);
  }

  static equals(a: User | PlainMessage<User> | undefined, b: User | PlainMessage<User> | undefined): boolean {
    return proto3.util.equals(User, a, b);
  }
}

