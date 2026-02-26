import { SensorTable } from "./component/SensorTable";
import { useTimeseries } from "./hook/useTimeseries";
import { Button } from "./ui/Button";
import { sensorApiClient } from "./modules/_runtime";
import "./global.css";

function App() {
  const { data, loadAll, loadMore, canLoadMore, isLoading, clearError, error } =
    useTimeseries({
      sensorApiClient,
      limit: 1000,
    });

  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-2xl font-bold mb-2">Spectral assignment</h1>
      <div className="mb-4 flex gap-2">
        <Button disabled={!canLoadMore || isLoading} onClick={loadMore}>
          {canLoadMore ? "Load more" : "No more data"}
        </Button>
        <Button disabled={!canLoadMore || isLoading} onClick={loadAll}>
          {canLoadMore ? "Load All" : "No more data"}
        </Button>
      </div>
      <div>
        {error && (
          <>
            <p>{error}</p>
            <button onClick={clearError}>x</button>
          </>
        )}
      </div>
      <p className="text-gray-600 mb-1">
        Total sensor data points: {data.length}
      </p>
      <div className="max-w-lg">
        {data.length > 0 ? <SensorTable data={data} /> : <p></p>}
      </div>
    </div>
  );
}

export default App;
