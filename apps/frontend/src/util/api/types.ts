export type ApiMethod = "GET" | "POST" | "PUT" | "DELETE";
type SimpleValue = string | number | boolean | null;

export interface ApiCaller {
  rpc<T, U>(
    method: Exclude<ApiMethod, "GET">,
    path: string,
    data: T,
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    responseConverter?: (data: any) => U,
  ): Promise<U>;

  rpc<T extends Record<string, SimpleValue>, U>(
    method: "GET",
    path: string,
    data: T,
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    responseConverter?: (data: any) => U,
  ): Promise<U>;
}

export class ApiError extends Error {
  constructor(message: string, cause: unknown) {
    super(message);
    this.name = "ApiError";
    this.cause = cause;
  }
}
