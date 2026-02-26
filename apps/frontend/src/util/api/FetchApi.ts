import { type ApiCaller, ApiError } from "./types";

export class FetchApi implements ApiCaller {
  private baseUrl: string;
  private fetch: typeof window.fetch;

  constructor(baseUrl: string, fetch?: typeof window.fetch) {
    if (!baseUrl) {
      throw new Error("baseUrl is required");
    }
    this.baseUrl = baseUrl;

    this.fetch = fetch || window.fetch.bind(window);
  }

  async rpc<T, U>(
    method: "GET" | "POST" | "PUT" | "DELETE",
    path: string,
    data: T,
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    responseConverter?: (data: any) => U,
  ): Promise<U> {
    const url = new URL(path, this.baseUrl);

    const options: RequestInit = {
      method,
      headers: new Headers(),
    };

    if (method === "GET") {
      for (const key in data) {
        const value = data[key];
        if (value === undefined) {
          continue;
        }

        if (typeof value === "object" || Array.isArray(value)) {
          throw new Error(
            `Invalid value for key ${key}: ${value}. Object or array is not supported`,
          );
        }

        url.searchParams.append(key, JSON.stringify(value));
      }
    } else {
      (options.headers as Headers).set("Content-Type", "application/json");
      options.body = JSON.stringify(data);
    }

    try {
      const response = await this.fetch(url, options);
      const data = await response.json();
      if (responseConverter) {
        return responseConverter(data);
      }
      return data;
    } catch (cause) {
      throw new ApiError("Failed to fetch", cause);
    }
  }
}
