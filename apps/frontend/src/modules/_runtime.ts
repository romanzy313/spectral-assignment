import { makeApiCaller } from "../util/api";
import { SensorApiClient } from "./sensor";

const apiCaller = makeApiCaller(
  import.meta.env.PUBLIC_SPECTRAL_GRPC_CLIENT_ORIGIN,
);

export const sensorApiClient = new SensorApiClient(apiCaller);
