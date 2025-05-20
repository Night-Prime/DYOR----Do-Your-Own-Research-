import React from 'react'
import { PlusIcon, CheckIcon } from './icons'

interface ButtonProps {
    name: string;
    select: boolean;
}

const AssetBtn: React.FC<ButtonProps> = ({ name, select }) => {
    return (
        <section className={`w-max h-8 rounded-3xl ${
        select ? 'bg-red-600' : 'bg-lime-700'} hover:bg-lime-900 cursor-pointer transition-colors shadow-lg`}>
            <main className='w-full h-full flex flex-row items-center justify-center gap-1 text-white p-2'>
                <h3 className='text-xs font-medium'>{name}</h3>
                {select ? <CheckIcon /> : <PlusIcon/>}
            </main>
        </section>
    );
};


export default AssetBtn;
