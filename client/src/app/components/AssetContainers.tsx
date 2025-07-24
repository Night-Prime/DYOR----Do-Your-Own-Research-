import React from 'react';

interface AssetProps {
  name: string;
  symbol: string;
  type: 'stock' | 'crypto' | string | undefined;
  price: number | string | undefined;
}

const AssetContainers: React.FC<AssetProps> = ({ name, symbol, type, price }) => {
  const badgeColor = {
    stock: 'bg-blue-100 text-blue-800 border-blue-200',
    crypto: 'bg-yellow-100 text-yellow-800 border-yellow-200',
    default: 'bg-gray-100 text-gray-800 border-gray-200'
  };

  const typeClass = type === 'stock'
    ? badgeColor.stock
    : type === 'crypto'
      ? badgeColor.crypto
      : badgeColor.default;

  return (
      <div className='cursor-pointer w-full h-full flex flex-col justify-between p-4 bg-white rounded-lg shadow-sm hover:shadow-md transition-shadow duration-100'>
        <div className='w-full flex justify-between items-start'>
          <h3 className='text-sm font-semibold text-gray-900 truncate max-w-[70%]' title={name}>
            {name}
          </h3>
          <span className={`text-xs px-2 py-1 rounded-full border ${typeClass} font-medium`}>
            {type?.toUpperCase()}
          </span>
        </div>

        <div className='my-2'>
          <h1 className='text-2xl font-bold text-gray-900'>${symbol}</h1>
        </div>

        <div className='text-lg font-medium text-gray-700'>
          ${typeof price === 'number' ? price.toLocaleString() : price}
        </div>
      </div>
  );
};

export default AssetContainers
