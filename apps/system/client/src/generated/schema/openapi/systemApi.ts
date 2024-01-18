/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */


export interface paths {
  "/ping": {
    /** Checks if the server is running */
    get: {
      responses: {
        /** @description Ping response */
        200: {
          content: never;
        };
        500: components["responses"]["InternalServerError"];
      };
    };
    /** Checks if the server is running */
    post: {
      requestBody?: {
        content: {
          "application/json": {
            /**
             * @description Name
             * @example John Doe
             */
            name: string;
          };
        };
      };
      responses: {
        /** @description Ping response */
        200: {
          content: never;
        };
        500: components["responses"]["InternalServerError"];
      };
    };
  };
  "/me": {
    /**
     * Get Admin User
     * @description Returns the admin user
     */
    get: {
      responses: {
        /** @description Admin User */
        200: {
          content: {
            "application/json": components["schemas"]["Me"];
          };
        };
        500: components["responses"]["InternalServerError"];
      };
    };
  };
  "/sign_up": {
    /**
     * Sign Up
     * @description Sign Up
     */
    post: operations["sign_up"];
  };
  "/login": {
    /**
     * Login
     * @description Login
     */
    post: operations["login"];
  };
}

export type webhooks = Record<string, never>;

export interface components {
  schemas: {
    User: {
      /** Format: uuid */
      userId: string;
      /** Format: email */
      email: string;
      name: string;
      phoneNumber?: string;
    };
    MultiFactor: {
      factorId: string;
      phoneNumber: string;
    };
    Workspace: {
      /** Format: uuid */
      workspaceId: string;
      name: string;
    };
    Member: {
      profile: components["schemas"]["MemberProfile"];
      user: components["schemas"]["User"];
    };
    MemberProfile: {
      display_name: string;
      idNumber?: string;
    };
    MembershipPeriod: {
      /** Format: date-time */
      start: string;
      /** Format: date-time */
      end: string;
    };
    Me: {
      /** @default false */
      emailVerified: boolean;
      multiFactor?: components["schemas"]["MultiFactor"];
      member: components["schemas"]["Member"];
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
          /**
           * @description error type
           * @example invalid_item_id
           */
          type?: string;
          /** @description A short, human-readable summary of the problem type */
          title?: string;
          /** @description A human-readable explanation specific to this occurrence of the problem. */
          detail?: string;
          /**
           * @description error code
           * @example invalid_item_id
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
          /**
           * @description error type
           * @example invalid_item_id
           */
          type?: string;
          /** @description A short, human-readable summary of the problem type */
          title?: string;
          /** @description A human-readable explanation specific to this occurrence of the problem. */
          detail?: string;
          /**
           * @description error code
           * @example invalid_item_id
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
          /**
           * @description error type
           * @example invalid_item_id
           */
          type?: string;
          /** @description A short, human-readable summary of the problem type */
          title?: string;
          /** @description A human-readable explanation specific to this occurrence of the problem. */
          detail?: string;
          /**
           * @description error code
           * @example invalid_item_id
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
          /**
           * @description error type
           * @example invalid_item_id
           */
          type?: string;
          /** @description A short, human-readable summary of the problem type */
          title?: string;
          /** @description A human-readable explanation specific to this occurrence of the problem. */
          detail?: string;
          /**
           * @description error code
           * @example invalid_item_id
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
          /**
           * @description error type
           * @example invalid_item_id
           */
          type?: string;
          /** @description A short, human-readable summary of the problem type */
          title?: string;
          /** @description A human-readable explanation specific to this occurrence of the problem. */
          detail?: string;
          /**
           * @description error code
           * @example invalid_item_id
           */
          code?: string;
        };
      };
    };
  };
  parameters: never;
  requestBodies: never;
  headers: never;
  pathItems: never;
}

export type $defs = Record<string, never>;

export type external = Record<string, never>;

export interface operations {

  /**
   * Sign Up
   * @description Sign Up
   */
  sign_up: {
    requestBody: {
      content: {
        "application/json": {
          /**
           * @description Name
           * @example John Doe
           */
          name: string;
        };
      };
    };
    responses: {
      /** @description Sign Up */
      200: {
        content: {
          "application/json": components["schemas"]["Me"];
        };
      };
      400: components["responses"]["BadRequestError"];
      401: components["responses"]["UnauthorizedError"];
      500: components["responses"]["InternalServerError"];
    };
  };
  /**
   * Login
   * @description Login
   */
  login: {
    responses: {
      /** @description Login response */
      200: {
        content: {
          "application/json": components["schemas"]["Me"];
        };
      };
      500: components["responses"]["InternalServerError"];
    };
  };
}
