import { beforeEach, describe, expect, it, vi, type Mock } from "vitest";

import { renderHook, act, waitFor } from "@testing-library/react";

import type { ISensorApiClient } from "../modules/sensor";
import { useTimeseries } from "./useTimeseries";

describe("useTimeseries", () => {
  let client: ISensorApiClient;

  beforeEach(() => {
    client = {
      getPage: vi.fn(async (req) => {
        // console.log("called get page", { req });
        return {
          nextCursor: null,
          data: [],
        };
      }),

      getCount: vi.fn(),
    };
  });

  it("initializes correctly", () => {
    const { result } = renderHook(() =>
      useTimeseries({
        sensorApiClient: client,
        limit: 2,
      }),
    );
    expect(result.current.data).toStrictEqual([]);
    expect(result.current.canLoadMore).toEqual(true);
  });

  it("handles end of iteration", async () => {
    const { result } = renderHook(() =>
      useTimeseries({
        sensorApiClient: client,
        limit: 2,
      }),
    );

    act(() => {
      result.current.loadMore();
    });

    await waitFor(() => {
      expect(result.current.canLoadMore).toBe(false);
    });
  });
});
