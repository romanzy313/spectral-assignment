import type { SensorData } from "../hooks/useTimeseries";

type Props = {
  data: SensorData[];
};

export function SensorTable({ data }: Props) {
  return (
    <table>
      <thead>
        <tr>
          <th>Time</th>
          <th>Value</th>
        </tr>
      </thead>
      <tbody>
        {data.map((item, index) => (
          <tr key={index}>
            <td>{item.timestamp}</td>
            <td>{item.value}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}
