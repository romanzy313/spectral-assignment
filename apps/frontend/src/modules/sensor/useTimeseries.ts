import { useRef, useState } from "react";
import type { SensorApiClient, SensorData } from ".";

export function useTimeseries({
  sensorApiClient,
  limit,
}: {
  sensorApiClient: SensorApiClient;
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

  const loadMore = async (): Promise<boolean> => {
    if (cursorRef.current === null) {
      return false;
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

      return true;
    } catch (error) {
      // console.error("Error getting timeseries data:", error);

      setError(error instanceof Error ? error.message : `${error}`);
      return false;
    } finally {
      setIsLoading(false);
    }
  };

  const loadAll = async () => {
    while (cursorRef.current !== null) {
      const ok = await loadMore();
      if (!ok) {
        return;
      }
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
