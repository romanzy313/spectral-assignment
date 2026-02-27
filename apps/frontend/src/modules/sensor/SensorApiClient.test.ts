import { describe, it, expect, vi, beforeEach, type Mock } from "vitest";
import { SensorApiClient } from "./SensorApiClient";
import { makeApiCaller } from "../../util/api";

describe("SensorApiClient", () => {
  let apiClient: SensorApiClient;
  let fetchMock: Mock;

  beforeEach(() => {
    fetchMock = vi.fn();

    apiClient = new SensorApiClient(
      makeApiCaller("http://example.com", fetchMock),
    );
  });

  it("should call successfully", async () => {
    fetchMock.mockImplementationOnce(() => ({
      ok: true,
      json: () => Promise.resolve({ nextCursor: 2, data: [{ t: 1, d: 42 }] }),
    }));

    const data = await apiClient.getPage({
      cursor: 1,
      limit: 100,
    });

    expect(fetchMock.mock.calls).toMatchInlineSnapshot(`
      [
        [
          "http://example.com/api/v1/sensor/data?cursor=1&limit=100",
          {
            "headers": Headers {},
            "method": "GET",
          },
        ],
      ]
    `);
    expect(data).toMatchInlineSnapshot(`
      {
        "data": [
          {
            "timestamp": 1970-01-01T00:00:01.000Z,
            "value": undefined,
          },
        ],
        "nextCursor": 2,
      }
    `);
  });
});
