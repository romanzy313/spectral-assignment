export interface ApiCaller {
  apiCall<T, U>(
    method: "POST" | "PUT" | "DELETE",
    path: string,
    data: T,
  ): Promise<U>;
  apiCall<T extends Record<string, string | number | boolean | null>, U>(
    method: "GET",
    path: string,
    data: T,
  ): Promise<U>;
}

export class ClientApi implements ApiCaller {
  constructor(private baseUrl: string) {}

  async apiCall<T, U>(
    method: "GET" | "POST" | "PUT" | "DELETE",
    path: string,
    data: T,
  ): Promise<U> {
    const url = new URL(path, this.baseUrl);
    const options: RequestInit = {
      method,
      headers: {
        "Content-Type": "application/json",
      },
    };

    if (method === "GET") {
      for (const key in data) {
        url.searchParams.append(key, JSON.stringify(data[key]));
      }
    } else {
      options.body = JSON.stringify(data);
    }

    try {
      const response = await fetch(url, options);
      return await response.json();
    } catch (cause) {
      throw new ApiError("Failed to fetch", cause);
    }
  }
}

export class ApiError extends Error {
  constructor(message: string, cause: unknown) {
    super(message);
    this.name = "ApiError";
    this.cause = cause;
  }
}
