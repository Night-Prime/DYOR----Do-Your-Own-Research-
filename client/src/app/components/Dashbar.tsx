"use client"
import { useAppSelector } from "../hooks/hook"

const Dashbar = () => {
  const user = useAppSelector((state) => state.auth.user);
  const name = user?.first_name ?? ''

  return (
    <section className=" fixed top-2 w-full h-24 border-b-1 border-white">
      <main className="w-full h-full flex flex-row justify-between p-6 text-lime-800">
        <div className='block'>
          <h1 className='text-3xl font-bold my-1'>
            News Feed
          </h1>
          <h3 className="text-md font-medium text-lime-900">
            Welcome back, {name}. Here’s what’s going on in the financial market today.
          </h3>

        </div>
      </main>
    </section>
  )
}

export default Dashbar
