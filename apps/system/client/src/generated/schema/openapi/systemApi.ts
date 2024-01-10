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
      userId: string;
      email: string;
      name: string;
      phoneNumber?: string;
    };
    MultiFactor: {
      factorId: string;
      phoneNumber: string;
    };
    Me: {
      /** @default false */
      emailVerified: boolean;
      multiFactor?: components["schemas"]["MultiFactor"];
      user: components["schemas"]["User"];
    };
  };
  responses: {
    /** @description General Error */
    GeneralError: {
      content: {
        "application/json": {
          /** Format: int32 */
          code?: number;
          message?: string;
        };
      };
    };
    /** @description Unauthorized */
    Unauthorized: {
      content: {
        "application/json": {
          /** Format: int32 */
          code?: number;
          message?: string;
        };
      };
    };
    /** @description Entity not found. */
    NotFound: {
      content: {
        "application/json": {
          /** Format: int32 */
          code?: number;
          message?: string;
        };
      };
    };
    /** @description Bad request */
    BadRequest: {
      content: {
        "application/json": {
          /** Format: int32 */
          code?: number;
          message?: string;
        };
      };
    };
    /** @description Internal server error */
    InternalServerError: {
      content: {
        "application/json": {
          /** Format: int32 */
          code?: number;
          message?: string;
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
    requestBody?: {
      content: {
        "application/json": {
          /**
           * @description First Name
           * @example John
           */
          first_name?: string;
          /**
           * @description Last Name
           * @example Doe
           */
          last_name?: string;
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
      400: components["responses"]["BadRequest"];
      401: components["responses"]["Unauthorized"];
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
