import type { SensorData } from "../modules/sensor";

type Props = {
  data: SensorData[];
};

export function SensorTable({ data }: Props) {
  return (
    <div className="rounded-lg border border-gray-200 max-h-160 overflow-y-auto">
      <table className="w-full text-sm text-left">
        <thead className="bg-gray-800 text-gray-100 text-xs tracking-wider sticky top-0">
          <tr>
            <th className="px-6 py-3">TIME</th>
            <th className="px-6 py-3">VALUE (kWh?)</th>
          </tr>
        </thead>
        <tbody className="divide-y divide-gray-200">
          {data.map((item, index) => (
            <tr
              key={index}
              className="bg-white odd:bg-gray-50 hover:bg-primary-500/30 transition-colors duration-150"
            >
              <td className="px-6 py-3 text-gray-700 font-mono">
                {item.timestamp.toISOString()}
              </td>
              <td className="px-6 py-3 text-gray-900 font-medium">
                {item.value}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
