// @generated by protoc-gen-connect-es v1.2.0 with parameter "target=ts"
// @generated from file me/v1/me.proto (package me.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { FindRequest, FindResponse, LoginRequest, LoginResponse, RegisterCompleteRequest, RegisterCompleteResponse, SignUpRequest, SignUpResponse, UpdateEmailRequest, UpdateEmailResponse, UpdateNameRequest, UpdateNameResponse, UpdatePhoneNumberRequest, UpdatePhoneNumberResponse } from "./me_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service me.v1.MeService
 */
export const MeService = {
  typeName: "me.v1.MeService",
  methods: {
    /**
     * @generated from rpc me.v1.MeService.Login
     */
    login: {
      name: "Login",
      I: LoginRequest,
      O: LoginResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc me.v1.MeService.SignUp
     */
    signUp: {
      name: "SignUp",
      I: SignUpRequest,
      O: SignUpResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc me.v1.MeService.RegisterComplete
     */
    registerComplete: {
      name: "RegisterComplete",
      I: RegisterCompleteRequest,
      O: RegisterCompleteResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc me.v1.MeService.Find
     */
    find: {
      name: "Find",
      I: FindRequest,
      O: FindResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc me.v1.MeService.UpdateName
     */
    updateName: {
      name: "UpdateName",
      I: UpdateNameRequest,
      O: UpdateNameResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc me.v1.MeService.UpdateEmail
     */
    updateEmail: {
      name: "UpdateEmail",
      I: UpdateEmailRequest,
      O: UpdateEmailResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc me.v1.MeService.UpdatePhoneNumber
     */
    updatePhoneNumber: {
      name: "UpdatePhoneNumber",
      I: UpdatePhoneNumberRequest,
      O: UpdatePhoneNumberResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
