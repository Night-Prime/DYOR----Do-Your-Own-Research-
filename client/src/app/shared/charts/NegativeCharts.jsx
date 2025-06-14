import React, { useMemo } from 'react';
import { AreaChart, Area, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer } from 'recharts';

const NegativeAreaChart = ({tickers}) => {
    const top5 = useMemo(() => {
        const sorted = tickers
            .slice()
            .sort(
                (a, b) =>
                    parseFloat(a.change_percentage.replace('%', '')) -
                    parseFloat(b.change_percentage.replace('%', ''))
            )
            .slice(0, 5)
            .reverse() // Optional: reverse to show from 5 to 1
            .map((ticker) => ({
                name: ticker.ticker,
                value: parseFloat(ticker.change_percentage.replace('%', '')),
                amount: parseFloat(ticker.price),
            }));
            
    
        return sorted;
    }, [tickers]);
    

    // Custom tooltip formatter
    const renderTooltip = ({ active, payload, label }) => {
        if (active && payload && payload.length) {
            return (
                <div className="bg-white p-2 border border-gray-200 rounded shadow-sm">
                    <p className="font-medium text-gray-800">{label}</p>
                    <p className="text-red-600">
                        {payload[0].value.toFixed(2)}%
                    </p>
                </div>
            );
        }
        return null;
    };

    return (
        <div className="w-full h-80 p-4 bg-white rounded-lg shadow-sm">
            <ResponsiveContainer width="100%" height="90%">
                <AreaChart
                    data={top5}
                    margin={{
                        top: 10,
                        right: 30,
                        left: 10,
                        bottom: 0,
                    }}
                >
                    <CartesianGrid strokeDasharray="3 3" stroke="#f0f0f0" />
                    <XAxis
                        dataKey="name"
                        tick={{ fill: '#666', fontSize: 12 }}
                        axisLine={{ stroke: '#e0e0e0' }}
                    />
                    <YAxis
                        tick={{ fill: '#666', fontSize: 12 }}
                        axisLine={{ stroke: '#e0e0e0' }}
                        tickFormatter={(value) => `${Math.abs(value)}`} // Show absolute values
                        label={{ value: 'Loss ($)', angle: -90, position: 'insideLeft', style: { textAnchor: 'middle', fill: '#666' } }}
                    />
                    <Tooltip content={renderTooltip} />
                    <defs>
                        <linearGradient id="negativeGradient" x1="0" y1="0" x2="0" y2="1">
                            <stop offset="0%" stopColor="#ef4444" stopOpacity={0.8} />
                            <stop offset="100%" stopColor="#ef4444" stopOpacity={0.2} />
                        </linearGradient>
                    </defs>
                    <Area
                        type="monotone"
                        dataKey="value"
                        stroke="#ef4444"
                        fill="url(#negativeGradient)"
                        strokeWidth={2}
                        activeDot={{ r: 5, fill: '#ef4444', stroke: '#fff' }}
                    />
                </AreaChart>
            </ResponsiveContainer>
        </div>
    );
};

export default NegativeAreaChart;