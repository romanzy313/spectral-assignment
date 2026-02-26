import { useEffect, useState } from "react";
import type { SensorApiClient, SensorData } from "../modules/sensor";

export function useTimeseries({
  sensorApiClient,
  limit,
}: {
  sensorApiClient: SensorApiClient;
  limit: number;
}) {
  const [data, setData] = useState<SensorData[]>([]);
  const [cursor, setCursor] = useState<number | null>(0);
  const [error, setError] = useState("");

  const [canLoadMore, setCanLoadMore] = useState(true);

  useEffect(() => {
    setCanLoadMore(cursor !== null);
  }, [cursor]);

  const clearError = () => {
    setError("");
  };

  const loadMore = async () => {
    if (cursor === null) {
      return;
    }

    try {
      const page = await sensorApiClient.getPage({
        cursor,
        limit,
      });
      setData((data) => [...data, ...page.data]);
      setCursor(page.nextCursor);
      setError("");
    } catch (error) {
      console.error("Error getting timeseries data:", error);

      setError(error instanceof Error ? error.message : `${error}`);
    }
  };

  return {
    data,
    loadMore,
    canLoadMore,
    error,
    clearError,
  };
}
