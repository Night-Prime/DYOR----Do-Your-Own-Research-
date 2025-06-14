import React from 'react'
import Link from 'next/link';
import { formatTimestamp } from '../utils/helper';
import Image from 'next/image';

interface InsightFeedProps {
    title: string;
    summary: string;
    url: string;
    time: string;
    imgLink: string;
}

const InsightFeed: React.FC<InsightFeedProps> = ({ title, summary, url, time, imgLink }) => {
    return (
        <section className='mx-auto w-full h-auto bg-gray-200 rounded-3xl shadow-sm p-6 mb-5'>
            <main className='w-full h-full text-lime-900 flex flex-col gap-3'>

                <div className='w-full flex items-center gap-4'>
                    {imgLink && (
                        <Image
                            src={imgLink}
                            alt='Article image'
                            width={60}
                            height={60}
                            className='rounded-xl object-cover'
                        />
                    )}
                    <h1 className='font-bold text-xl leading-tight flex-1'>{title}</h1>
                </div>

                <p className='font-medium text-sm text-lime-800'>
                    {summary}
                </p>

                <div className='w-full flex justify-between items-center mt-2'>
                    <small className='font-bold text-xs text-lime-700'>
                        {formatTimestamp(time)}
                    </small>
                    {url && (
                        <Link
                            href={url}
                            className='text-lime-700 text-xs underline hover:text-red-900 font-bold'
                            target='_blank'
                            rel='noopener noreferrer'
                        >
                            Read More
                        </Link>
                    )}
                </div>

            </main>
        </section>

    )
}

export default InsightFeed
