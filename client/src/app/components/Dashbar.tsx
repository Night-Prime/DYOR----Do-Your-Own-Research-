const Dashbar = () => {

  return (
    <section className=" fixed top-2 w-full h-24 border-b-1 border-white">
      <main className="w-full h-full flex flex-row justify-between p-6 text-lime-800">
        <div className='block'>
          <h1 className='text-4xl'>
            Your Insight Feed
          </h1>
          <h3 className="text-md font-bold">
            Your financial world at a glance. Smart insights, tailored for you.
          </h3>
        </div>
      </main>
    </section>
  )
}

export default Dashbar
