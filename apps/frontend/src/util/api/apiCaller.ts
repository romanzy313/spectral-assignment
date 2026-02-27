import { ApiError } from "./ApiError";
import type { ApiCaller, AppCallArgs, GetData } from "./types";

export function makeApiCaller(
  baseUrl: string,
  fetchFn: typeof window.fetch = window.fetch.bind(window),
): ApiCaller {
  if (!baseUrl) throw new Error("baseUrl is required");

  return async function <T>({ method, path, data }: AppCallArgs): Promise<T> {
    const url = new URL(path, baseUrl);
    const options: RequestInit = { method, headers: new Headers() };

    if (data) {
      if (method === "GET") {
        for (const key in data) {
          const value = (data as GetData)[key];
          if (value === undefined) continue;

          if (value === null) {
            url.searchParams.append(key, "null");
          } else {
            url.searchParams.append(key, value.toString());
          }
        }
      } else {
        (options.headers as Headers).set("Content-Type", "application/json");
        options.body = JSON.stringify(data);
      }
    }

    try {
      const response = await fetchFn(url, options);
      const json = await response.json();

      if (!response.ok) {
        throw new Error("message" in json ? json.message : "unknown error");
      }

      return json;
    } catch (cause) {
      throw new ApiError(
        cause instanceof Error ? cause.message : `${cause}`,
        cause,
      );
    }
  };
}
