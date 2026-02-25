import { useEffect, useState } from "react";
import type { ApiCaller } from "../utils/api";

export type SensorData = {
  t: number;
  v: number;
};

type SensorPageRequest = {
  cursor: number | null;
  limit: number;
};

type SensorPage = {
  nextCursor: number | null;
  data: SensorData[];
};

export function useTimeseries({
  api,
  limit,
}: {
  api: ApiCaller;
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
      const more = await api.apiCall<SensorPageRequest, SensorPage>(
        "GET",
        "/page",
        {
          cursor: cursor,
          limit,
        },
      );
      setData((data) => [...data, ...more.data]);
      setCursor(more.nextCursor);
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
