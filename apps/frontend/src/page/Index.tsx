import { SensorTable } from "../component/SensorTable";
import { useTimeseries } from "../modules/sensor";
import { sensorApiClient } from "../modules/_runtime";
import { Button } from "../ui/button";

export default function IndexPage() {
  const { data, loadAll, loadMore, canLoadMore, isLoading, clearError, error } =
    useTimeseries({
      sensorApiClient,
      limit: 1000,
    });

  return (
    <>
      <h1 className="text-2xl font-bold mb-2">Spectral assignment</h1>
      <div className="mb-4 flex gap-2">
        <Button
          disabled={!canLoadMore || isLoading}
          onClick={loadMore}
          data-testid="load-more-button"
        >
          {canLoadMore ? "Load more" : "No more data"}
        </Button>
        <Button
          disabled={!canLoadMore || isLoading}
          onClick={loadAll}
          data-testid="load-all-button"
        >
          {canLoadMore ? "Load All" : "No more data"}
        </Button>
      </div>
      {error && (
        <div className="flex justify-between text-red-500 border border-red-500 rounded px-4 py-2 mb-2">
          <p data-testid="error-message">{error}</p>
          <button onClick={clearError}>x</button>
        </div>
      )}
      <p className="text-gray-600 mb-1">
        Total sensor data points:{" "}
        <span data-testid="total-data-points">{data.length}</span>
      </p>
      <div>{data.length > 0 ? <SensorTable data={data} /> : <p></p>}</div>
    </>
  );
}
