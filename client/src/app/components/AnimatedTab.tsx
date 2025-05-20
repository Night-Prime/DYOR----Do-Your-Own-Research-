/* eslint-disable @typescript-eslint/no-explicit-any */
import React, { useState } from 'react';


interface AnimatedTabSwitchProps {
    tab1Label: string;
    tab2Label: string;
    onTabChange?: (tabIndex: number) => void;
    chartComponent1?: React.ReactNode;
    chartComponent2?: React.ReactNode;
}

const AnimatedTab: React.FC<AnimatedTabSwitchProps> = ({
    tab1Label,
    tab2Label,
    onTabChange,
    chartComponent1,
    chartComponent2
}) => {
    const [activeTab, setActiveTab] = useState(0);

    const handleTabChange = (tabIndex: number) => {
        setActiveTab(tabIndex);
        if (onTabChange) {
            onTabChange(tabIndex);
        }
    };

    return (
        <div className="w-full bg-gray-200 max-w-xl mx-auto p-4 rounded-3xl my-2">
            <div className='rounded-3xl shadow-sm relative'>
                <div
                    className='absolute h-full w-1/2 bg-lime-950 rounded-3xl shadow-sm transition-transform duration-300 ease-in-out'
                    style={{ transform: `translateX(${activeTab * 100}%)` }}
                ></div>
                <div className='flex relative z-10'>
                    <button
                        className={`flex-1 py-2 px-4 rounded-3xl text-sm font-bold transition-colors duration-300 cursor-pointer ${activeTab === 0 ? 'text-white' : 'text-lime-800'}`}
                        onClick={() => handleTabChange(0)}
                    >
                        {tab1Label}
                    </button>
                    <button
                        className={`flex-1 py-2 px-4 rounded-3xl text-sm font-bold transition-colors duration-300 cursor-pointer ${activeTab === 1 ? 'text-white' : 'text-lime-800'}`}
                        onClick={() => handleTabChange(1)}
                    >
                        {tab2Label}
                    </button>
                </div>
            </div>
            <div>
                {activeTab === 0 && (
                    <div className="p-2">
                        {chartComponent1}
                    </div>
                )}
                {activeTab === 1 && (
                    <div className="p-2">
                        {chartComponent2}
                    </div>
                )}
            </div>
        </div>
    );
};

export default AnimatedTab;