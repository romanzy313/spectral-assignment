import type { ApiCaller } from "../../util/api";
import type {
  ISensorApiClient,
  SensorCount,
  SensorPage,
  SensorPageRequest,
} from "./types";

export class SensorApiClient implements ISensorApiClient {
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

  getCount(): Promise<SensorCount> {
    return this.apiCaller.rpc("GET", "/api/v1/sensor/count", {});
  }
}
