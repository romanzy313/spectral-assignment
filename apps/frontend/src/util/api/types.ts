export type GetData = Record<string, string | number | boolean | null>;
export type ApiCallMethods = "GET" | "POST" | "PUT" | "DELETE";

export type AppCallArgs =
  | {
      method: "GET";
      path: string;
      data?: Record<string, unknown>;
    }
  | {
      method: Exclude<ApiCallMethods, "GET">;
      path: string;
      data?: GetData;
    };

export type ApiCaller = <T>(args: AppCallArgs) => Promise<T>;
