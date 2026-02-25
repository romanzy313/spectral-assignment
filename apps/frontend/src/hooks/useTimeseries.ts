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

export function useTimeseries({ api }: { api: ApiCaller }) {
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
      const fullData = await api.apiCall<SensorPageRequest, SensorPage>(
        "GET",
        "/page",
        {
          cursor: cursor,
          limit: 9999,
        },
      );
      setData(fullData.data);
      setCursor(fullData.nextCursor);
      setError("");
    } catch (error) {
      console.error("Error fetching timeseries data:", error);

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
