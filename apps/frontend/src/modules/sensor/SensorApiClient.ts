import type { ApiCaller } from "../../util/api";
import type { SensorPage, SensorPageRequest } from "./types";

export class SensorApiClient {
  private apiCaller: ApiCaller;

  constructor(apiCaller: ApiCaller) {
    this.apiCaller = apiCaller;
  }

  async getPage(req: SensorPageRequest): Promise<SensorPage> {
    const res = await this.apiCaller<{
      nextCursor: number | null;
      data: { t: number; v: number }[];
    }>({
      method: "GET",
      path: "/api/v1/sensor/data",
      data: req,
    });

    return {
      nextCursor: res.nextCursor,
      data: res.data.map(({ t, v }) => ({
        timestamp: new Date(t * 1000),
        value: v,
      })),
    };
  }
}
