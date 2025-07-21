"use client"
import { useAppSelector } from "../hooks/hook"
import { useHeaderData } from "../hooks/useHeaderData";

const Dashbar = () => {
  const user = useAppSelector((state) => state.auth.user);
  const name = user?.first_name ?? '';

  const metadata = useHeaderData();
  console.log("Header Data: ", metadata);

  

  return (
    <section className=" fixed top-0 w-full h-24 border-b-1 border-white">
      <main className="w-full h-full flex flex-row justify-between p-6 text-lime-800">
        <div className='block'>
          <h1 className='text-2xl font-bold my-1'>
           {metadata?.header || 'Insights'}
          </h1>
          <h3 className="text-md font-medium text-lime-900">
            Welcome back, {name} ,
             {metadata?.description}
          </h3>

        </div>
      </main>
    </section>
  )
}

export default Dashbar
