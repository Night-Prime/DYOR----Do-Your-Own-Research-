"use client"
import React from 'react'
import { HomeOutlinedIcon, PortfolioIcon, PowerIcon, SettingsIcon } from '../shared/icons'
import { logOut } from '../utils/auth'
import { useRouter } from 'next/navigation'
import Link from 'next/link'

const Sidebar = () => {
    const router = useRouter();
    const handleClick = async() => {
        const response = await logOut();
        if(response.success) {
            router.push('/')
        }
    }
    return (
        <div className='w-full h-full'>
            <main className='py-6 h-full w-full flex flex-col justify-between items-center'>
                <div className='block cursor-pointer'>
                    <h3 className='text-md text-lime-800 font-extrabold'>
                        Insights
                    </h3>
                    <div className='w-full h-1 bg-lime-800 rounded-3xl'></div>
                </div>

                <div className='flex flex-col items-center gap-6'>
                    <div className="flex flex-col items-center justify-center">
                        <Link href={'/dashboard'} >
                            <HomeOutlinedIcon className='cursor-pointer w-8 h-8' />
                        </Link>
                        <p className='text-xs font-semibold'>Home</p>
                    </div> 

                    <div className="flex flex-col items-center justify-center">
                        <Link href={'/dashboard/portfolio'} >
                            <PortfolioIcon className='cursor-pointer w-8 h-8' />
                        </Link>
                        <p className='text-xs font-semibold'>Portfolio</p>
                    </div> 
                    
                </div>

                <div className='flex flex-col items-end gap-6'>
                    <SettingsIcon className='cursor-pointer w-6 h-6' />
                    <span onClick={() => handleClick()}>
                        <PowerIcon className='cursor-pointer w-6 h-6 text-red-600' />
                    </span>
                </div>
            </main>
        </div>
    )
}

export default Sidebar
