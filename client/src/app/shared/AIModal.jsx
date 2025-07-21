import React from 'react';
import { ClearRounded } from '@mui/icons-material';

const AIModal = ({ close, summary, asset }) => {

    const handleClose = () => {
        if (close) close()
    }
    return (
        <div className="fixed inset-0 z-[9999] bg-black/50 backdrop-blur-xs flex items-center justify-center p-4">
            <div className="w-full max-w-2xl bg-white rounded-lg shadow-xl flex flex-col max-h-[90vh]">
                {/* Header with close button */}
                <div className="flex justify-end p-4">
                    <button
                        className="text-gray-500 hover:text-gray-700 focus:outline-none"
                        onClick={handleClose}
                    >
                        <ClearRounded />
                    </button>
                </div>

                {/* Scrollable content area */}
                <div className="px-6 pb-6 flex-1 overflow-y-auto">
                    {/* Asset symbol */}
                    <div className="mb-4">
                        <h1 className="text-6xl font-bold">
                            {asset ? `$${asset.symbol}` : ""}
                        </h1>
                    </div>

                    {/* Asset name and type */}
                    <div className="mb-6">
                        <h3 className="text-xl text-gray-600">
                            {asset ? `${asset.name} (${asset.type})` : ""}
                        </h3>
                    </div>

                    {/* Scrollable text content */}
                    <div className="bg-gray-50 rounded-lg p-4">
                        <p className="text-sm leading-relaxed overflow-y-auto max-h-[50vh]">
                            {summary}
                        </p>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default AIModal
