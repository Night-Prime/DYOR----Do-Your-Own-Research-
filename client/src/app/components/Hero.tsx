"use client"

const Hero = () => {
  return (
    <div className="w-full h-full flex flex-row justify-start items-center gap-4">
      <div className="w-[40%] h-full flex flex-col justify-center items-start gap-2">
        <h1 className="text-xl font-extrabold text-lime-900">What&apos;s D.Y.O.R ?</h1>
        <p className="text-xs font-bold text-lime-800 w-[60%]">
            Real-time insights for smarter investments in stocks, bonds, and crypto.
        </p>
      </div>

      <div className="w-[60%] h-full flex flex-col justify-center items-start gap-2 mt-16">
        <h1 className="text-4xl font-semibold text-lime-700">
            <span className="text-lime-800">D.Y.O.R</span> combines real-time market data + social sentiment to give you actionable investment signals.
        </h1>
      </div>
    </div>
  )
}

export default Hero
