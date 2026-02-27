import { describe, it, expect, vi, beforeEach, type Mock } from "vitest";
import { makeApiCaller, type ApiCaller } from "../../util/api";

describe("SensorApiClient", () => {
  let apiCaller: ApiCaller;
  let fetchMock: Mock;

  beforeEach(() => {
    fetchMock = vi.fn();

    apiCaller = makeApiCaller("http://test", fetchMock);
  });

  it("creates correct GET request", async () => {
    fetchMock.mockImplementationOnce(() => ({
      ok: true,
      json: () => Promise.resolve({ testing: true }),
    }));

    await apiCaller({
      method: "GET",
      path: "/api",
      data: {
        hello: "world",
        number: 42,
        works: true,
        errors: null,
      },
    });
    expect(fetchMock.mock.calls).toMatchInlineSnapshot(`
      [
        [
          "http://test/api?hello=world&number=42&works=true&errors=null",
          {
            "headers": Headers {},
            "method": "GET",
          },
        ],
      ]
    `);
  });

  it("creates correct POST request", async () => {
    fetchMock.mockImplementationOnce(() => ({
      ok: true,
      json: () => Promise.resolve({ testing: true }),
    }));

    await apiCaller({
      method: "POST",
      path: "/api",
      data: {
        hello: "world",
        number: 42,
        works: true,
        errors: null,
      },
    });
    expect(fetchMock.mock.calls).toMatchInlineSnapshot(`
      [
        [
          "http://test/api",
          {
            "body": "{"hello":"world","number":42,"works":true,"errors":null}",
            "headers": Headers {},
            "method": "POST",
          },
        ],
      ]
    `);
  });

  it("handles errors", async () => {
    fetchMock.mockImplementationOnce(() => ({
      ok: false,
      json: () => Promise.resolve({ message: "test error" }),
    }));

    await expect(
      apiCaller({
        method: "POST",
        path: "/api",
      }),
    ).rejects.toThrowErrorMatchingInlineSnapshot(`[ApiError: test error]`);
  });

  it("returns response data", async () => {
    fetchMock.mockImplementationOnce(() => ({
      ok: true,
      json: () => Promise.resolve({ testing: true }),
    }));

    const res = await apiCaller({
      method: "POST",
      path: "/api",
    });

    expect(res).toMatchInlineSnapshot(`
      {
        "testing": true,
      }
    `);
  });
});
