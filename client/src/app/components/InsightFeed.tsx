import React from 'react'

interface InsightFeedProps {
    title : string;
    news : string;
}

const InsightFeed: React.FC<InsightFeedProps> = ({title, news}) => {
    return (
        <section className='mx-auto w-[90%] h-32 bg-gray-200 rounded-3xl shadow-sm p-6 my-5'>
            <main className='w-full h-full text-lime-900 flex flex-col items-start justify-evenly'>
                <h1 className='font-bold text-lg'>
                    {title}
                </h1>

                <p className='font-medium text-sm'>
                    {news}
                </p>
            </main>
        </section>
    )
}

export default InsightFeed
