"use client"
import { useHeaderData } from "../hooks/useHeaderData";

const Dashbar = () => {
  const metadata = useHeaderData();

  

  return (
    <section className=" fixed top-0 w-full h-24 border-b-1 border-white">
      <main className="w-full h-full flex flex-row justify-between p-6 text-lime-800">
        <div className='block'>
          <h1 className='text-2xl font-bold my-1'>
           {metadata?.header || 'Insights'}
          </h1>
          <h3 className="text-md font-medium text-lime-900">
             {metadata?.description}
          </h3>

        </div>
      </main>
    </section>
  )
}

export default Dashbar
