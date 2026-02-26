export type SensorData = {
  timestamp: Date;
  value: number;
};

export type SensorPageRequest = {
  cursor: number | null;
  limit: number;
};

export type SensorPage = {
  nextCursor: number | null;
  data: SensorData[];
};

export type SensorCount = {
  count: number;
};

export interface ISensorApiClient {
  getPage(req: SensorPageRequest): Promise<SensorPage>;
  getCount(): Promise<SensorCount>;
}
