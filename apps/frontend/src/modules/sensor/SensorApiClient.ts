import type { ApiCaller } from "../../util/api";
import type { SensorPage, SensorPageRequest } from "./types";

export class SensorApiClient {
  private apiCaller: ApiCaller;

  constructor(apiCaller: ApiCaller) {
    this.apiCaller = apiCaller;
  }

  getPage(req: SensorPageRequest): Promise<SensorPage> {
    return this.apiCaller.rpc(
      "GET",
      "/api/v1/sensor/data",
      req,
      (res: {
        nextCursor: number | null;
        data: { t: number; v: number }[];
      }) => ({
        nextCursor: res.nextCursor,
        data: res.data.map(({ t, v }) => ({
          timestamp: new Date(t * 1000),
          value: v,
        })),
      }),
    );
  }
}
