import React from 'react';
import {
    RadarChart,
    Radar,
    PolarGrid,
    PolarAngleAxis,
    PolarRadiusAxis,
    ResponsiveContainer,
    Tooltip,
    Legend
} from 'recharts';

const RiskRadarChart = () => {
    const data = [
        {
            subject: 'Liquidity',
            Portfolio: 120,
            Benchmark: 110,
            fullMark: 150,
        },
        {
            subject: 'Volatility',
            Portfolio: 98,
            Benchmark: 130,
            fullMark: 150,
        },
        {
            subject: 'Correlation',
            Portfolio: 99,
            Benchmark: 100,
            fullMark: 150,
        },
        {
            subject: 'Beta',
            Portfolio: 86,
            Benchmark: 130,
            fullMark: 150,
        },
        {
            subject: 'Drawdown',
            Portfolio: 65,
            Benchmark: 85,
            fullMark: 150,
        },
    ];

    // Custom tooltip formatter
    const renderTooltipContent = ({ active, payload }) => {
        if (active && payload && payload.length) {
            return (
                <div className="bg-white p-2 border border-gray-200 rounded shadow-sm">
                    <p className="font-medium text-gray-800">{payload[0].payload.subject}</p>
                    {payload.map((entry, index) => (
                        <p key={`item-${index}`} style={{ color: entry.color }}>
                            {`${entry.name}: ${entry.value}`}
                        </p>
                    ))}
                </div>
            );
        }
        return null;
    };

    return (
        <div className="w-full h-80 p-4 bg-white rounded-lg shadow-sm">
            <ResponsiveContainer width="100%" height="90%">
                <RadarChart cx="50%" cy="50%" outerRadius="80%" data={data}>
                    <PolarGrid stroke="#e0e0e0" />
                    <PolarAngleAxis
                        dataKey="subject"
                        tick={{ fill: '#666', fontSize: 12 }}
                    />
                    <PolarRadiusAxis angle={30} domain={[0, 150]} />

                    <Radar
                        name="Portfolio"
                        dataKey="Portfolio"
                        stroke="#8884d8"
                        fill="#8884d8"
                        fillOpacity={0.6}
                    />

                    <Radar
                        name="Benchmark"
                        dataKey="Benchmark"
                        stroke="#82ca9d"
                        fill="#82ca9d"
                        fillOpacity={0.6}
                    />

                    <Tooltip content={renderTooltipContent} />
                    <Legend
                        wrapperStyle={{
                            paddingTop: '10px',
                            fontSize: '12px'
                        }}
                    />
                </RadarChart>
            </ResponsiveContainer>
        </div>
    );
};

export default RiskRadarChart;