import React from 'react';
import { AreaChart, Area, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer, Legend } from 'recharts';

const TopPerformingChart = () => {
    // Data showing positive growth trend for top performers
    const data = [
        {
            name: 'BTC',
            value: 2800,
        },
        {
            name: 'GOLD',
            value: 1200,
        },
        {
            name: 'AAPL',
            value: 2100,
        },
        {
            name: 'FTSE',
            value: 3100,
        },
        {
            name: 'USD',
            value: 3800,
        },
    ];

    // Custom tooltip formatter
    const renderTooltip = ({ active, payload, label }) => {
        if (active && payload && payload.length) {
            return (
                <div className="bg-white p-2 border border-gray-200 rounded shadow-sm">
                    <p className="font-medium text-gray-800">{label}</p>
                    <p className="text-emerald-600">
                        ${payload[0].value.toLocaleString()}
                    </p>
                </div>
            );
        }
        return null;
    };

    return (
        <div className="w-full h-80 p-4 bg-white rounded-lg shadow-sm">
            <div className="flex justify-between items-center mb-4">
                <h3 className="text-lg font-semibold text-gray-800">Top Performing Assets</h3>
                <div className="text-sm font-medium text-emerald-600">
                    +189% YTD
                </div>
            </div>

            <ResponsiveContainer width="100%" height="85%">
                <AreaChart
                    data={data}
                    margin={{
                        top: 5,
                        right: 30,
                        left: 10,
                        bottom: 5,
                    }}
                >
                    <CartesianGrid strokeDasharray="3 3" stroke="#f0f0f0" vertical={false} />
                    <XAxis
                        dataKey="name"
                        tick={{ fill: '#666', fontSize: 12 }}
                        axisLine={{ stroke: '#e0e0e0' }}
                        tickLine={false}
                    />
                    <YAxis
                        tick={{ fill: '#666', fontSize: 12 }}
                        axisLine={false}
                        tickLine={false}
                        tickFormatter={(value) => `$${value / 1000}k`} // Format as $Xk
                        domain={[0, 'dataMax + 1000']} // Add some space above the highest value
                    />
                    <Tooltip content={renderTooltip} />
                    <defs>
                        <linearGradient id="positiveGradient" x1="0" y1="0" x2="0" y2="1">
                            <stop offset="0%" stopColor="#10b981" stopOpacity={0.8} />
                            <stop offset="95%" stopColor="#10b981" stopOpacity={0.1} />
                        </linearGradient>
                    </defs>
                    <Area
                        type="monotone"
                        dataKey="value"
                        stroke="#10b981"
                        strokeWidth={2}
                        fill="url(#positiveGradient)"
                        activeDot={{
                            r: 6,
                            fill: '#10b981',
                            stroke: '#fff',
                            strokeWidth: 2
                        }}
                    />
                </AreaChart>
            </ResponsiveContainer>
        </div>
    );
};

export default TopPerformingChart;