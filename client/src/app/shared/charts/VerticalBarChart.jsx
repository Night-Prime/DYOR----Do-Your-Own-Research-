import React from 'react';
import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
  ResponsiveContainer
} from 'recharts';

const data = [
  {
    name: 'GOLD',
    price: 800,
  },
  {
    name: 'BTC',
    price: 967,
  },
  {
    name: 'AAPL',
    price: 1098,
  },
  {
    name: 'TSLA',
    price: 1200,
  },
  {
    name: 'ETH',
    price: 1108,
  },
  {
    name: 'SILVER',
    price: 680,
  },
];

const VerticalBarChart = () => {
  return (
    <div className="bg-gray-100 rounded-3xl w-full h-80 p-2">
      <ResponsiveContainer width="100%" height="100%">
        <BarChart
          layout="vertical"
          data={data}
          margin={{
            top: 20,
            right: 30,
            left: 20,
            bottom: 5,
          }}
        >
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis type="number" />
          <YAxis dataKey="name" type="category" scale="band" width={80} />
          <Tooltip 
            formatter={(value) => [`${value}`, 'Value']}
            labelStyle={{ color: '#333' }}
            contentStyle={{ 
              backgroundColor: 'white', 
              borderRadius: '6px',
              border: '1px solid #e0e0e0',
              boxShadow: '0 2px 4px rgba(0,0,0,0.1)'
            }}
          />
          <Legend />
          <Bar 
            dataKey="price" 
            fill="#006400" 
            radius={[0, 4, 4, 0]} // Rounded corners on right side
            barSize={20}
          />
        </BarChart>
      </ResponsiveContainer>
    </div>
  );
};

export default VerticalBarChart;