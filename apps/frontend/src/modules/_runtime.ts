import { FetchApi } from "../util/api/FetchApi";
import { SensorApiClient } from "./sensor";

export const sensorApiClient = new SensorApiClient(
  new FetchApi(import.meta.env.PUBLIC_SPECTRAL_GRPC_CLIENT_ORIGIN),
);
