import { useRef, useState } from "react";
import type { ISensorApiClient, SensorData } from "../modules/sensor";

export function useTimeseries({
  sensorApiClient,
  limit,
}: {
  sensorApiClient: ISensorApiClient;
  limit: number;
}) {
  const [data, setData] = useState<SensorData[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState("");
  const [canLoadMore, setCanLoadMore] = useState(true);

  const cursorRef = useRef<number | null>(0);

  const updateCursor = (next: number | null) => {
    cursorRef.current = next;
    setCanLoadMore(next !== null);
  };

  const clearError = () => {
    setError("");
  };

  const loadMore = async () => {
    if (cursorRef.current === null) {
      return;
    }

    try {
      setIsLoading(true);
      const page = await sensorApiClient.getPage({
        cursor: cursorRef.current,
        limit,
      });
      setData((data) => [...data, ...page.data]);
      updateCursor(page.nextCursor);
      setError("");
    } catch (error) {
      console.error("Error getting timeseries data:", error);

      setError(error instanceof Error ? error.message : `${error}`);
    } finally {
      setIsLoading(false);
    }
  };

  const loadAll = async () => {
    while (cursorRef.current !== null) {
      await loadMore();
    }
  };

  return {
    data,
    loadAll,
    loadMore,
    canLoadMore,
    isLoading,
    error,
    clearError,
  };
}
