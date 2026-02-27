/* eslint-disable @typescript-eslint/no-explicit-any */
import { beforeEach, describe, expect, it, vi, type Mock } from "vitest";

import { renderHook, act } from "@testing-library/react";

import { useTimeseries } from "./useTimeseries";
import { SensorApiClient } from ".";

describe("useTimeseries", () => {
  let getPage: Mock<SensorApiClient["getPage"]>;
  let client: SensorApiClient;

  beforeEach(() => {
    getPage = vi.fn();
    client = {
      getPage,
    } as any;
  });

  it("initializes correctly", () => {
    const { result } = renderHook(() =>
      useTimeseries({
        sensorApiClient: client,
        limit: 2,
      }),
    );
    expect(result.current.data).toEqual([]);
    expect(result.current.canLoadMore).toEqual(true);
    expect(result.current.error).toBe("");
  });

  it("loads data correctly via loadMore()", async () => {
    const { result } = renderHook(() =>
      useTimeseries({
        sensorApiClient: client,
        limit: 2,
      }),
    );

    getPage
      .mockImplementationOnce(async () => {
        return { nextCursor: 2, data: [{ timestamp: new Date(1), value: 1 }] };
      })
      .mockImplementationOnce(async () => {
        return {
          nextCursor: null,
          data: [{ timestamp: new Date(2), value: 2 }],
        };
      });

    await act(async () => {
      await result.current.loadMore();
      await result.current.loadMore();
    });

    expect(result.current.isLoading).toBe(false);
    expect(result.current.canLoadMore).toBe(false);
    expect(getPage.mock.calls).toMatchInlineSnapshot(`
      [
        [
          {
            "cursor": 0,
            "limit": 2,
          },
        ],
        [
          {
            "cursor": 2,
            "limit": 2,
          },
        ],
      ]
    `);
    expect(result.current.data).toMatchInlineSnapshot(`
      [
        {
          "timestamp": 1970-01-01T00:00:00.001Z,
          "value": 1,
        },
        {
          "timestamp": 1970-01-01T00:00:00.002Z,
          "value": 2,
        },
      ]
    `);
  });

  it("loads data correctly via loadAll()", async () => {
    const { result } = renderHook(() =>
      useTimeseries({
        sensorApiClient: client,
        limit: 2,
      }),
    );

    getPage
      .mockImplementationOnce(async () => {
        return { nextCursor: 2, data: [{ timestamp: new Date(1), value: 1 }] };
      })
      .mockImplementationOnce(async () => {
        return {
          nextCursor: null,
          data: [{ timestamp: new Date(2), value: 2 }],
        };
      });

    await act(async () => {
      await result.current.loadAll();
    });

    expect(result.current.isLoading).toBe(false);
    expect(result.current.canLoadMore).toBe(false);
    expect(getPage.mock.calls).toMatchInlineSnapshot(`
      [
        [
          {
            "cursor": 0,
            "limit": 2,
          },
        ],
        [
          {
            "cursor": 2,
            "limit": 2,
          },
        ],
      ]
    `);
    expect(result.current.data).toMatchInlineSnapshot(`
      [
        {
          "timestamp": 1970-01-01T00:00:00.001Z,
          "value": 1,
        },
        {
          "timestamp": 1970-01-01T00:00:00.002Z,
          "value": 2,
        },
      ]
    `);
  });

  it("handles errors", async () => {
    const { result } = renderHook(() =>
      useTimeseries({
        sensorApiClient: client,
        limit: 2,
      }),
    );

    getPage.mockImplementationOnce(async () => {
      throw new Error("failed");
    });

    await act(async () => {
      await result.current.loadAll();
    });

    expect(result.current.isLoading).toBe(false);
    expect(result.current.canLoadMore).toBe(true); // loading state didnt change!
    expect(result.current.error).toBe("failed");

    act(() => {
      result.current.clearError();
    });

    expect(result.current.error).toBe("");
  });
});
