export class ApiError extends Error {
  constructor(message: string, cause: unknown) {
    super(message);
    this.name = "ApiError";
    this.cause = cause;
  }
}
