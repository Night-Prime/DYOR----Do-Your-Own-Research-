"use client"

import Hero from "../components/Hero"
import Navbar from "../components/Navbar"
import SubHero from "../components/SubHero"

const page = () => {
  return (
      <div className="w-full h-full flex flex-col justify-center items-center gap-10">
          <main className="px-10 w-full h-full">
              <Navbar />
              <section className="w-full h-3/4 flex justify-center items-center">
                 <Hero />
              </section>
          </main>
          <SubHero />
      </div>
  )
}

export default page
