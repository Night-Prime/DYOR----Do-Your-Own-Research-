import React from 'react';
import { PieChart, Pie, Cell, Legend, Surface, Curve } from 'recharts';

const RADIAN = Math.PI / 180;
const data = [
  { name: 'Bearish', value: 40, color: '#D2042D' }, // Softer red
  { name: 'Stable', value: 20, color: '#FEBE10' },  // Softer yellow
  { name: 'Bullish', value: 40, color: '#006400' }, // Softer green
];

const PortfolioHealthGauge = () => {
  const width = 250;
  const height = 180;
  
  // Center points
  const cx = width / 2;
  const cy = height - 60;
  
  // Radius values
  const iR = 50;
  const oR = 100;
  
  // Needle value (position)
  const value = 50;
  const total = data.reduce((sum, entry) => sum + entry.value, 0);
  
  // Custom shape to create rounded edges on pie segments
  const CustomSector = (props) => {
    const { cx, cy, innerRadius, outerRadius, startAngle, endAngle, fill } = props;
    
    // Calculate the coordinates for the outer arc path
    const sin = Math.sin(-RADIAN * startAngle);
    const cos = Math.cos(-RADIAN * startAngle);
    const sin2 = Math.sin(-RADIAN * endAngle);
    const cos2 = Math.cos(-RADIAN * endAngle);
    
    // Start and end points of outer arc
    const outerStartX = cx + outerRadius * cos;
    const outerStartY = cy + outerRadius * sin;
    const outerEndX = cx + outerRadius * cos2;
    const outerEndY = cy + outerRadius * sin2;
    
    // Start and end points of inner arc
    const innerStartX = cx + innerRadius * cos;
    const innerStartY = cy + innerRadius * sin;
    const innerEndX = cx + innerRadius * cos2;
    const innerEndY = cy + innerRadius * sin2;
    
    // Create the path for the sector with rounded edges
    // We're adding a small radius for the corners
    const cornerRadius = 5;
    
    // This is a simplified approach - for a complete solution with perfect
    // rounded corners, more complex path calculations would be needed
    return (
      <g>
        <path
          d={`
            M ${innerStartX} ${innerStartY}
            L ${outerStartX} ${outerStartY}
            A ${outerRadius} ${outerRadius} 0 
              ${endAngle - startAngle > 180 ? 1 : 0} 
              1 
              ${outerEndX} ${outerEndY}
            L ${innerEndX} ${innerEndY}
            A ${innerRadius} ${innerRadius} 0 
              ${endAngle - startAngle > 180 ? 1 : 0} 
              0 
              ${innerStartX} ${innerStartY}
            Z
          `}
          fill={fill}
          stroke="white"
          strokeWidth={1}
        />
      </g>
    );
  };
  
  // Custom legend renderer
  const renderCustomizedLegend = (props) => {
    const { payload } = props;
    
    return (
      <div className="flex justify-center mt-1 pt-2">
        {payload.map((entry, index) => (
          <div key={`legend-item-${index}`} className="flex items-center mx-3">
            <div 
              className="w-3 h-3 mr-1 rounded-full" 
              style={{ backgroundColor: entry.color }}
            />
            <span className="text-xs text-gray-700">{entry.value}</span>
          </div>
        ))}
      </div>
    );
  };
  
  // Needle rendering function
  const renderNeedle = () => {
    const ang = 180.0 * (1 - value / total);
    const length = oR * 0.8;
    const sin = Math.sin(-RADIAN * ang);
    const cos = Math.cos(-RADIAN * ang);
    const r = 5;
    const x0 = cx;
    const y0 = cy;
    const xba = x0 + r * sin;
    const yba = y0 - r * cos;
    const xbb = x0 - r * sin;
    const ybb = y0 + r * cos;
    const xp = x0 + length * cos;
    const yp = y0 + length * sin;

    return [
      <circle key="needle-circle" cx={x0} cy={y0} r={r} fill="#555555" />, // Softer color
      <path
        key="needle-path"
        d={`M${xba} ${yba}L${xbb} ${ybb} L${xp} ${yp} L${xba} ${yba}`}
        fill="#555555" // Softer color
      />,
    ];
  };

  return (
    <div className="w-full flex flex-col items-center bg-gray-50 rounded-lg p-4">
      <PieChart width={width} height={height}>
        <defs>
          {/* Add subtle gradient to make gauge look better */}
          <linearGradient id="bearishGradient" x1="0" y1="0" x2="0" y2="1">
            <stop offset="0%" stopColor="#e57373" />
            <stop offset="100%" stopColor="#ef9a9a" />
          </linearGradient>
          <linearGradient id="normalGradient" x1="0" y1="0" x2="0" y2="1">
            <stop offset="0%" stopColor="#fff59d" />
            <stop offset="100%" stopColor="#fff9c4" />
          </linearGradient>
          <linearGradient id="bullishGradient" x1="0" y1="0" x2="0" y2="1">
            <stop offset="0%" stopColor="#81c784" />
            <stop offset="100%" stopColor="#a5d6a7" />
          </linearGradient>
        </defs>
        
        <Pie
          dataKey="value"
          startAngle={180}
          endAngle={0}
          data={data}
          cx={cx}
          cy={cy}
          innerRadius={iR}
          outerRadius={oR}
          fill="#8884d8"
          stroke="white"
          strokeWidth={2}
          activeShape={CustomSector}
          shape={<CustomSector />}
        >
          {data.map((entry, index) => {
            // Use gradient IDs instead of flat colors
            let fillId;
            if (index === 0) fillId = "url(#bearishGradient)";
            else if (index === 1) fillId = "url(#normalGradient)";
            else fillId = "url(#bullishGradient)";
            
            return <Cell key={`cell-${index}`} fill={entry.color} />;
          })}
        </Pie>
        {renderNeedle()}
        <Legend 
          content={renderCustomizedLegend}
          verticalAlign="bottom" 
          height={36}
        />
      </PieChart>
    </div>
  );
};

export default PortfolioHealthGauge;