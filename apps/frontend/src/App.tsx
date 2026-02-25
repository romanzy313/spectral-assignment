import { SensorTable } from "./component/SensorTable";
import { useTimeseries } from "./hook/useTimeseries";
import { ClientApi } from "./util/api";

const api = new ClientApi(import.meta.env.PUBLIC_SPECTRAL_GRPC_CLIENT_ORIGIN);

function App() {
  const { data, loadMore, canLoadMore, clearError, error } = useTimeseries({
    api,
    limit: 1000,
  });

  return (
    <>
      <h1>Spectral assignment</h1>
      <div>
        <button disabled={!canLoadMore} onClick={loadMore}>
          Load more
        </button>
        <p>Hello, Spectral</p>
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
    </>
  );
}

export default App;
