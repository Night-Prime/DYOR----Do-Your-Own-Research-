import React from 'react'
import { PowerIcon, SettingsIcon } from '../shared/icons'

const Sidebar = () => {
    return (
        <div className='w-full h-full'>
            <main className='py-6 h-full w-full flex flex-col justify-between items-center'>
                <div className='block cursor-pointer'>
                    <h3 className='text-lg text-lime-800 font-extrabold'>
                        DYOR
                    </h3>
                    <div className='w-full h-1 bg-lime-800 rounded-3xl'></div>
                </div>

                <div className='flex flex-col items-center gap-6'>
                    {/* <HomeOutlinedIcon className='cursor-pointer w-6 h-6' /> */}
                </div>

                <div className='flex flex-col items-end gap-6'>
                    <SettingsIcon className='cursor-pointer w-6 h-6' />
                    <PowerIcon className='cursor-pointer w-6 h-6 text-red-600' />
                </div>
            </main>
        </div>
    )
}

export default Sidebar
