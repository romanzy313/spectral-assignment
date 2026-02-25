import { SensorTable } from "./component/SensorTable";
import { useTimeseries } from "./hooks/useTimeseries";
import { ClientApi } from "./utils/api";

const api = new ClientApi("http://localhost:12001");

function App() {
  const { data, loadMore, canLoadMore, clearError, error } = useTimeseries({
    api,
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
        {data.length > 0 ? <SensorTable data={data} /> : <p>No data</p>}
      </div>
    </>
  );
}

export default App;
