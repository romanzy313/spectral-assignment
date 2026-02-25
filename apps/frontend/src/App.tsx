import { SensorTable } from "./component/SensorTable";
import { useTimeseries } from "./hook/useTimeseries";
import { ClientApi } from "./util/api";
import "./global.css";
import { Button } from "./ui/Button";

const api = new ClientApi(import.meta.env.PUBLIC_SPECTRAL_GRPC_CLIENT_ORIGIN);

function App() {
  const { data, loadMore, canLoadMore, clearError, error } = useTimeseries({
    api,
    limit: 1000,
  });

  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-2xl font-bold mb-2">Spectral assignment</h1>
      <div className="mb-4">
        <Button disabled={!canLoadMore} onClick={loadMore}>
          Load more
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
      <div>
        <p>Total sensor readings: {data.length}</p>
        {data.length > 0 ? <SensorTable data={data} /> : <p>No data</p>}
      </div>
    </div>
  );
}

export default App;
