/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */


export interface paths {
  "/ping": {
    /** Checks if the server is running */
    get: operations["Ping"];
  };
  "/api/v1/auth/otp": {
    /**
     * Send OTP
     * @description One Time Password (OTP) to user.
     */
    post: operations["APIV1AuthByOtp"];
  };
  "/api/v1/auth/otp/verify": {
    /**
     * Verify OTP
     * @description Verify OTP sent by user.
     */
    post: operations["APIV1VerifyOTP"];
  };
  "/api/v1/auth/invitations/process/email": {
    /** Process an invitation by verifying token and email */
    post: operations["APIV1ProcessInvitationEmail"];
  };
  "/api/v1/auth/invitations/process/oauth": {
    /** Process an invitation by verifying token and OAuth, and register or add user to workspace. */
    post: operations["APIV1ProcessInvitationOAuth"];
  };
  "/api/v1/auth/invitations": {
    /** Get Invitation by token */
    get: operations["APIV1GetInvitationByToken"];
  };
  "/api/v1/auth/oauth": {
    /**
     * Auth by OAuth
     * @description Auth by OAuth.
     */
    post: operations["APIV1AuthByOAuth"];
  };
  "/api/v1/me": {
    /**
     * Get Admin User
     * @description Returns the admin user
     */
    get: operations["APIV1GetMe"];
  };
  "/api/v1/me/profile": {
    /**
     * Update Profile
     * @description Updates the user profile
     */
    put: operations["APIV1UpdateProfile"];
  };
  "/api/v1/me/profile/photo": {
    /**
     * Update Profile Photo
     * @description Updates the user profile photo
     */
    put: operations["APIV1UpdateProfilePhoto"];
    /**
     * Delete Profile Photo
     * @description Deletes the user profile photo
     */
    delete: operations["APIV1RemoveProfilePhoto"];
  };
  "/api/v1/me/member/profile": {
    /**
     * Update Me Member Profile
     * @description Updates Me the member profile
     */
    put: operations["APIV1UpdateMeMemberProfile"];
  };
  "/api/v1/me/workspace/leave": {
    /**
     * Leave Workspace
     * @description Leaves the workspace
     */
    post: operations["APIV1LeaveWorkspace"];
  };
  "/api/v1/workspaces": {
    /**
     * Get Joined Workspaces
     * @description Returns the workspaces the user is a member of
     */
    get: operations["APIV1GetWorkspaces"];
    /**
     * Create Workspace
     * @description Creates a new workspace
     */
    post: operations["APIV1CreateWorkspace"];
  };
  "/api/v1/workspaces/{workspaceId}": {
    /**
     * Update Workspace
     * @description Updates the workspace
     */
    put: operations["APIV1UpdateWorkspace"];
  };
  "/api/v1/members": {
    /**
     * Get Members
     * @description Returns the members of the workspace
     */
    get: operations["APIV1GetMembers"];
  };
  "/api/v1/members/{memberId}": {
    /**
     * Remove Member
     * @description Removes a member from the workspace
     */
    delete: operations["APIV1RemoveMember"];
  };
  "/api/v1/members/{memberId}/role": {
    /**
     * Update Member Role
     * @description Updates the role of a member
     */
    put: operations["APIV1UpdateMemberRole"];
  };
  "/api/v1/invitations": {
    /**
     * Get pending invitations
     * @description Returns the pending invitations (not used yet)
     */
    get: operations["APIV1GetInvitations"];
  };
  "/api/v1/members/invitations/bulk": {
    /** Invite multiple users to the workspace by email */
    post: operations["APIV1InviteMultipleUsers"];
  };
  "/api/v1/members/invitations/{invitationId}/accept": {
    /** Accept an invitation to join a workspace */
    post: operations["APIV1AcceptInvitation"];
  };
  "/api/v1/members/invitations/{invitationId}/revoke": {
    /** Revoke invitation */
    post: operations["APIV1RevokeInvitation"];
  };
  "/api/v1/members/invitations/{invitationId}/resend": {
    /** Resend invitation */
    post: operations["APIV1ResendInvitation"];
  };
}

export type webhooks = Record<string, never>;

export interface components {
  schemas: {
    JwtToken: {
      /** @description JWT token */
      token: string;
    };
    /**
     * @description Authentication provider
     * @enum {string}
     */
    AuthProvider: "email" | "google";
    User: {
      /** Format: uuid */
      userId: string;
      /** Format: email */
      email: string;
      name?: string;
      phoneNumber?: string;
      /** Format: uri */
      photo?: string;
    };
    MultiFactor: {
      factorId: string;
      phoneNumber: string;
    };
    Workspace: {
      /**
       * Format: uuid
       * @description workspace id in the format of UUID v7.
       * @example 123e4567-e89b-12d3-a456-426614174000
       */
      workspaceId: string;
      /** @description workspace name */
      name: string;
      /** @description workspace subdomain (e.x. example-test) */
      subdomain: string;
    };
    Workspaces: components["schemas"]["Workspace"][];
    Inviter: {
      member: components["schemas"]["Member"];
      workspace: components["schemas"]["Workspace"];
    };
    Member: {
      /** Format: uuid */
      id: string;
      profile: components["schemas"]["MemberProfile"];
      user: components["schemas"]["User"];
      role: components["schemas"]["MemberRole"];
      /** @enum {string} */
      membershipStatus: "ACTIVE" | "LEFT";
    };
    /**
     * @description Role of the member
     * @enum {string}
     */
    MemberRole: "OWNER" | "ADMIN" | "MEMBER" | "GUEST";
    MemberProfile: {
      displayName: string;
      idNumber?: string;
      /** @description User's bio */
      bio?: string;
    };
    Members: components["schemas"]["Member"][];
    Invitation: {
      /**
       * Format: uuid
       * @description Invitation ID
       */
      id: string;
      accepted: boolean;
      /** Format: date-time */
      expiredAt: string;
      /**
       * Format: email
       * @description Email of the invitee
       */
      inviteeEmail: string;
      /** @description Display name of the invitee */
      displayName: string;
      inviter: components["schemas"]["Member"];
    };
    Invitations: components["schemas"]["Invitation"][];
    Invitee: {
      /** @description Display name of the invitee */
      name: string;
      /**
       * Format: email
       * @description Email of the invitee
       */
      email: string;
    };
    Invitees: components["schemas"]["Invitee"][];
    MembershipPeriod: {
      /** Format: date-time */
      start: string;
      /** Format: date-time */
      end: string;
    };
    Me: {
      self: components["schemas"]["User"];
      member?: components["schemas"]["Member"];
      currentWorkspace?: components["schemas"]["Workspace"];
      joinedWorkspaces: components["schemas"]["Workspace"][];
      receivedInvitations?: components["schemas"]["ReceivedInvitation"][];
      providers?: components["schemas"]["AuthProvider"][];
    };
    ReceivedInvitation: {
      invitation: components["schemas"]["Invitation"];
      inviter: components["schemas"]["Inviter"];
    };
  };
  responses: {
    /** @description General Error */
    GeneralError: {
      content: {
        "application/json": {
          /**
           * @description The HTTP status code generated for this occurrence of the problem.
           * @example 400
           */
          status?: number;
          /** @description A short, human-readable summary of the problem type */
          title?: string;
          /** @description A human-readable explanation specific to this occurrence of the problem. */
          detail?: string;
          /**
           * @description A custom code identifying the specific error.
           * @example 400-001
           */
          code?: string;
        };
      };
    };
    /** @description Bad request */
    BadRequestError: {
      content: {
        "application/json": {
          /**
           * @description The HTTP status code generated for this occurrence of the problem.
           * @example 400
           */
          status?: number;
          /** @description A short, human-readable summary of the problem type */
          title?: string;
          /** @description A human-readable explanation specific to this occurrence of the problem. */
          detail?: string;
          /**
           * @description A custom code identifying the specific error.
           * @example 400-001
           */
          code?: string;
        };
      };
    };
    /** @description Forbidden */
    ForbiddenError: {
      content: {
        "application/json": {
          /**
           * @description The HTTP status code generated for this occurrence of the problem.
           * @example 400
           */
          status?: number;
          /** @description A short, human-readable summary of the problem type */
          title?: string;
          /** @description A human-readable explanation specific to this occurrence of the problem. */
          detail?: string;
          /**
           * @description A custom code identifying the specific error.
           * @example 400-001
           */
          code?: string;
        };
      };
    };
    /** @description Entity not found. */
    NotFoundError: {
      content: {
        "application/json": {
          /**
           * @description The HTTP status code generated for this occurrence of the problem.
           * @example 400
           */
          status?: number;
          /** @description A short, human-readable summary of the problem type */
          title?: string;
          /** @description A human-readable explanation specific to this occurrence of the problem. */
          detail?: string;
          /**
           * @description A custom code identifying the specific error.
           * @example 400-001
           */
          code?: string;
        };
      };
    };
    /** @description Unauthorized */
    UnauthorizedError: {
      content: {
        "application/json": {
          /**
           * @description The HTTP status code generated for this occurrence of the problem.
           * @example 400
           */
          status?: number;
          /** @description A short, human-readable summary of the problem type */
          title?: string;
          /** @description A human-readable explanation specific to this occurrence of the problem. */
          detail?: string;
          /**
           * @description A custom code identifying the specific error.
           * @example 400-001
           */
          code?: string;
        };
      };
    };
    /** @description Conflict */
    ConflictError: {
      content: {
        "application/json": {
          /**
           * @description The HTTP status code generated for this occurrence of the problem.
           * @example 400
           */
          status?: number;
          /** @description A short, human-readable summary of the problem type */
          title?: string;
          /** @description A human-readable explanation specific to this occurrence of the problem. */
          detail?: string;
          /**
           * @description A custom code identifying the specific error.
           * @example 400-001
           */
          code?: string;
        };
      };
    };
    /** @description Gone */
    GoneError: {
      content: {
        "application/json": {
          /**
           * @description The HTTP status code generated for this occurrence of the problem.
           * @example 400
           */
          status?: number;
          /** @description A short, human-readable summary of the problem type */
          title?: string;
          /** @description A human-readable explanation specific to this occurrence of the problem. */
          detail?: string;
          /**
           * @description A custom code identifying the specific error.
           * @example 400-001
           */
          code?: string;
        };
      };
    };
    /** @description Too many requests */
    TooManyRequestsError: {
      content: {
        "application/json": {
          /**
           * @description The HTTP status code generated for this occurrence of the problem.
           * @example 400
           */
          status?: number;
          /** @description A short, human-readable summary of the problem type */
          title?: string;
          /** @description A human-readable explanation specific to this occurrence of the problem. */
          detail?: string;
          /**
           * @description A custom code identifying the specific error.
           * @example 400-001
           */
          code?: string;
        };
      };
    };
    /** @description Internal server error */
    InternalServerError: {
      content: {
        "application/json": {
          /**
           * @description The HTTP status code generated for this occurrence of the problem.
           * @example 400
           */
          status?: number;
          /** @description A short, human-readable summary of the problem type */
          title?: string;
          /** @description A human-readable explanation specific to this occurrence of the problem. */
          detail?: string;
          /**
           * @description A custom code identifying the specific error.
           * @example 400-001
           */
          code?: string;
        };
      };
    };
    /** @description Get invitation by token */
    GetInvitationByTokenResponse: {
      content: {
        "application/json": {
          receivedInvitation: components["schemas"]["ReceivedInvitation"];
        };
      };
    };
    /** @description Successfully sent invitations */
    InvitationsBulkResponse: {
      content: {
        "application/json": {
          /** @description Total number of invitations */
          total: number;
          successfulInvitations: components["schemas"]["Invitations"];
          failedInvitations: components["schemas"]["Invitations"];
          registeredInvitations: components["schemas"]["Invitations"];
        };
      };
    };
    /** @description Invitations without the revoked one */
    RevokeInvitationResponse: {
      content: {
        "application/json": components["schemas"]["Invitations"];
      };
    };
    /** @description Resend invitation response */
    ResendInvitationResponse: {
      content: {
        "application/json": components["schemas"]["Invitation"];
      };
    };
    /** @description Get joined workspaces */
    WorkspacesResponse: {
      content: {
        "application/json": {
          workspaces: components["schemas"]["Workspace"][];
        };
      };
    };
    /** @description Get workspace invitations */
    InvitationsResponse: {
      content: {
        "application/json": {
          invitations: components["schemas"]["Invitation"][];
        };
      };
    };
    /** @description Get workspace members */
    MembersResponse: {
      content: {
        "application/json": {
          members: components["schemas"]["Member"][];
        };
      };
    };
    /** @description Me */
    MeResponse: {
      content: {
        "application/json": {
          me: components["schemas"]["Me"];
        };
      };
    };
    /** @description Update profile */
    UpdateProfileResponse: {
      content: {
        "application/json": {
          me: components["schemas"]["Me"];
        };
      };
    };
    /** @description Update profile photo */
    UpdateProfilePhotoResponse: {
      content: {
        "application/json": {
          me: components["schemas"]["Me"];
        };
      };
    };
    /** @description Remove profile photo */
    RemoveProfilePhotoResponse: {
      content: {
        "application/json": {
          me: components["schemas"]["Me"];
        };
      };
    };
    /** @description Update Me Member Profile */
    UpdateMeMemberProfileResponse: {
      content: {
        "application/json": {
          me: components["schemas"]["Me"];
        };
      };
    };
    /** @description Invitation accepted and user added to the workspace */
    InvitationsAcceptResponse: {
      content: {
        "application/json": {
          me: components["schemas"]["Me"];
        };
      };
    };
    /** @description Workspace created */
    CreateWorkspaceResponse: {
      content: {
        "application/json": {
          workspace: components["schemas"]["Workspace"];
        };
      };
    };
    /** @description Login response */
    LoginResponse: {
      content: {
        "application/json": components["schemas"]["Me"];
      };
    };
    /** @description Successfully verified OTP. The user is now authenticated. */
    AuthVerifyOTPResponse: {
      content: {
        "application/json": components["schemas"]["JwtToken"];
      };
    };
    /** @description Auth by OAuth response */
    AuthByOAuthResponse: {
      content: {
        "application/json": components["schemas"]["Me"];
      };
    };
    /** @description Invitation processed by OAuth */
    InvitationProcessOAuthResponse: {
      content: {
        "application/json": components["schemas"]["Me"];
      };
    };
    /** @description Update member role response */
    UpdateMemberRoleResponse: {
      content: {
        "application/json": {
          member: components["schemas"]["Member"];
        };
      };
    };
    /** @description Workspace updated */
    UpdateWorkspaceResponse: {
      content: {
        "application/json": {
          workspace: components["schemas"]["Workspace"];
        };
      };
    };
  };
  parameters: never;
  requestBodies: {
    /** @description Authenticate user by OTP. */
    AuthByOtpPost: {
      content: {
        "application/json": {
          /** Format: email */
          email: string;
        };
      };
    };
    /** @description Verify OTP sent by user. */
    AuthVerifyPost: {
      content: {
        "application/json": {
          /** Format: email */
          email: string;
          otp: string;
        };
      };
    };
    /** @description Creates a new workspace */
    CreateWorkspace: {
      content: {
        "application/json": {
          /** @description workspace subdomain (e.x. example-test) */
          subdomain: string;
        };
      };
    };
    /** @description Update Me Profile */
    UpdateMeProfile: {
      content: {
        "application/json": {
          profile: {
            name?: string;
          };
        };
      };
    };
    /** @description Update Me Profile Photo */
    UpdateMeProfilePhoto: {
      content: {
        "multipart/form-data": {
          /** Format: binary */
          photo: string;
        };
      };
    };
    /** @description Process an invitation by verifying token and email, and register or add user to workspace. */
    InvitationProcessEmail: {
      content: {
        "application/json": {
          /**
           * Format: uuid
           * @description The invitation token.
           */
          token: string;
          /**
           * Format: email
           * @description The user's email address.
           */
          email: string;
        };
      };
    };
    /** @description Process an invitation by verifying token and OAuth, and register or add user to workspace. */
    InvitationProcessOAuth: {
      content: {
        "application/json": {
          /**
           * Format: uuid
           * @description The invitation token.
           */
          token: string;
        };
      };
    };
    /** @description Bulk invite users to workspace */
    InvitationsBulk: {
      content: {
        "application/json": {
          invitees?: components["schemas"]["Invitee"][];
        };
      };
    };
    /** @description Update Me Member Profile */
    UpdateMeMemberProfile: {
      content: {
        "application/json": {
          memberProfile: components["schemas"]["MemberProfile"];
        };
      };
    };
    /** @description Update member role */
    UpdateMemberRole: {
      content: {
        "application/json": {
          /** @enum {string} */
          role: "admin" | "member" | "guest";
        };
      };
    };
    /** @description Update workspace */
    UpdateWorkspace: {
      content: {
        "application/json": {
          /** @description workspace name */
          name?: string;
          /** @description workspace subdomain */
          subdomain?: string;
        };
      };
    };
  };
  headers: never;
  pathItems: never;
}

export type $defs = Record<string, never>;

export type external = Record<string, never>;

export interface operations {

  /** Checks if the server is running */
  Ping: {
    responses: {
      /** @description Ping response */
      200: {
        content: {
          "application/json": {
            /**
             * @description Ping response message
             * @example pong
             */
            message?: string;
          };
        };
      };
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Send OTP
   * @description One Time Password (OTP) to user.
   */
  APIV1AuthByOtp: {
    requestBody: components["requestBodies"]["AuthByOtpPost"];
    responses: {
      /** @description OTP has been sent successfully. */
      200: {
        content: never;
      };
      400: components["responses"]["BadRequestError"];
      429: components["responses"]["TooManyRequestsError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Verify OTP
   * @description Verify OTP sent by user.
   */
  APIV1VerifyOTP: {
    requestBody: components["requestBodies"]["AuthVerifyPost"];
    responses: {
      200: components["responses"]["AuthVerifyOTPResponse"];
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorizedError"];
      429: components["responses"]["TooManyRequestsError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /** Process an invitation by verifying token and email */
  APIV1ProcessInvitationEmail: {
    requestBody: components["requestBodies"]["InvitationProcessEmail"];
    responses: {
      /** @description OTP has been sent successfully. */
      200: {
        content: never;
      };
      400: components["responses"]["BadRequestError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /** Process an invitation by verifying token and OAuth, and register or add user to workspace. */
  APIV1ProcessInvitationOAuth: {
    requestBody: components["requestBodies"]["InvitationProcessOAuth"];
    responses: {
      200: components["responses"]["InvitationProcessOAuthResponse"];
      400: components["responses"]["BadRequestError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /** Get Invitation by token */
  APIV1GetInvitationByToken: {
    parameters: {
      query: {
        /** @description Invitation token */
        token: string;
      };
    };
    responses: {
      200: components["responses"]["GetInvitationByTokenResponse"];
      400: components["responses"]["BadRequestError"];
      409: components["responses"]["ConflictError"];
      410: components["responses"]["GoneError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Auth by OAuth
   * @description Auth by OAuth.
   */
  APIV1AuthByOAuth: {
    responses: {
      200: components["responses"]["AuthByOAuthResponse"];
      401: components["responses"]["UnauthorizedError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Get Admin User
   * @description Returns the admin user
   */
  APIV1GetMe: {
    responses: {
      200: components["responses"]["MeResponse"];
      401: components["responses"]["UnauthorizedError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Update Profile
   * @description Updates the user profile
   */
  APIV1UpdateProfile: {
    requestBody: components["requestBodies"]["UpdateMeProfile"];
    responses: {
      200: components["responses"]["UpdateProfileResponse"];
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorizedError"];
      403: components["responses"]["ForbiddenError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Update Profile Photo
   * @description Updates the user profile photo
   */
  APIV1UpdateProfilePhoto: {
    requestBody: components["requestBodies"]["UpdateMeProfilePhoto"];
    responses: {
      200: components["responses"]["UpdateProfilePhotoResponse"];
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorizedError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Delete Profile Photo
   * @description Deletes the user profile photo
   */
  APIV1RemoveProfilePhoto: {
    responses: {
      200: components["responses"]["RemoveProfilePhotoResponse"];
      401: components["responses"]["UnauthorizedError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Update Me Member Profile
   * @description Updates Me the member profile
   */
  APIV1UpdateMeMemberProfile: {
    requestBody: components["requestBodies"]["UpdateMeMemberProfile"];
    responses: {
      200: components["responses"]["UpdateMeMemberProfileResponse"];
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorizedError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Leave Workspace
   * @description Leaves the workspace
   */
  APIV1LeaveWorkspace: {
    responses: {
      /** @description Workspace left */
      204: {
        content: never;
      };
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorizedError"];
      403: components["responses"]["ForbiddenError"];
      404: components["responses"]["NotFoundError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Get Joined Workspaces
   * @description Returns the workspaces the user is a member of
   */
  APIV1GetWorkspaces: {
    responses: {
      200: components["responses"]["WorkspacesResponse"];
      401: components["responses"]["UnauthorizedError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Create Workspace
   * @description Creates a new workspace
   */
  APIV1CreateWorkspace: {
    requestBody: components["requestBodies"]["CreateWorkspace"];
    responses: {
      201: components["responses"]["CreateWorkspaceResponse"];
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorizedError"];
      409: components["responses"]["ConflictError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Update Workspace
   * @description Updates the workspace
   */
  APIV1UpdateWorkspace: {
    parameters: {
      path: {
        /** @description Workspace id */
        workspaceId: string;
      };
    };
    requestBody: components["requestBodies"]["UpdateWorkspace"];
    responses: {
      200: components["responses"]["UpdateWorkspaceResponse"];
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorizedError"];
      403: components["responses"]["ForbiddenError"];
      404: components["responses"]["NotFoundError"];
      409: components["responses"]["ConflictError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Get Members
   * @description Returns the members of the workspace
   */
  APIV1GetMembers: {
    responses: {
      200: components["responses"]["MembersResponse"];
      401: components["responses"]["UnauthorizedError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Remove Member
   * @description Removes a member from the workspace
   */
  APIV1RemoveMember: {
    parameters: {
      path: {
        /** @description Member id */
        memberId: string;
      };
    };
    responses: {
      /** @description Member removed */
      204: {
        content: never;
      };
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorizedError"];
      403: components["responses"]["ForbiddenError"];
      404: components["responses"]["NotFoundError"];
      409: components["responses"]["ConflictError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Update Member Role
   * @description Updates the role of a member
   */
  APIV1UpdateMemberRole: {
    parameters: {
      path: {
        /** @description Member id */
        memberId: string;
      };
    };
    requestBody: components["requestBodies"]["UpdateMemberRole"];
    responses: {
      200: components["responses"]["UpdateMemberRoleResponse"];
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorizedError"];
      403: components["responses"]["ForbiddenError"];
      404: components["responses"]["NotFoundError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Get pending invitations
   * @description Returns the pending invitations (not used yet)
   */
  APIV1GetInvitations: {
    parameters: {
      query?: {
        /** @description Invitation status */
        status?: "accepted";
      };
    };
    responses: {
      200: components["responses"]["InvitationsResponse"];
      401: components["responses"]["UnauthorizedError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /** Invite multiple users to the workspace by email */
  APIV1InviteMultipleUsers: {
    requestBody: components["requestBodies"]["InvitationsBulk"];
    responses: {
      200: components["responses"]["InvitationsBulkResponse"];
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorizedError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /** Accept an invitation to join a workspace */
  APIV1AcceptInvitation: {
    parameters: {
      path: {
        /** @description Invitation token */
        invitationId: string;
      };
    };
    responses: {
      200: components["responses"]["InvitationsAcceptResponse"];
      401: components["responses"]["UnauthorizedError"];
      409: components["responses"]["ConflictError"];
      410: components["responses"]["GoneError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /** Revoke invitation */
  APIV1RevokeInvitation: {
    parameters: {
      path: {
        /** @description Invitation id */
        invitationId: string;
      };
    };
    responses: {
      200: components["responses"]["RevokeInvitationResponse"];
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorizedError"];
      403: components["responses"]["ForbiddenError"];
      404: components["responses"]["NotFoundError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /** Resend invitation */
  APIV1ResendInvitation: {
    parameters: {
      path: {
        /** @description Invitation id */
        invitationId: string;
      };
    };
    responses: {
      200: components["responses"]["ResendInvitationResponse"];
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorizedError"];
      403: components["responses"]["ForbiddenError"];
      404: components["responses"]["NotFoundError"];
      500: components["responses"]["InternalServerError"];
    };
  };
}
