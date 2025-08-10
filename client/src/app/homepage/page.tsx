"use client"

import Navbar from "../components/Navbar"

const page = () => {
    return (
        <div className="w-full h-auto bg-zinc-950">
            <div className="w-full h-full flex flex-col justify-center items-center gap-10">
                <Navbar />
                {/* Hero-section */}
                <div
                    className="w-screen h-screen flex justify-center items-center relative"
                    style={{
                        backgroundImage: 'url("/assets/img/Gemini_Generated_Image_m8cxeem8cxeem8cx.png")',
                        backgroundSize: 'cover',
                        backgroundPosition: 'center',
                    }}
                >
                    <div className="absolute inset-0 bg-black opacity-[40%]"></div>

                    <div className="h-[50%] w-[80%] m-auto text-center flex flex-col justify-center items-center gap-4 z-10">
                        <h1 className="text-7xl font-bold text-white">
                            Make Smarter Investment Decisions
                        </h1>
                        <p className="text-lg mt-6 font-semibold text-white/80">
                            Comprehensive market data, portfolio analytics, and cross-asset <span className="font-bold">Insights</span> in one powerful platform built for investors and finance professionals.
                        </p>

                        <div className="w-[60%] flex flex-row justify-center items-center gap-4 mt-6">
                            <button
                                className="text-white px-8 py-4 rounded-full bg-white/15 hover:bg-white/25 transition-all duration-300 font-medium backdrop-blur-sm border border-white/20 hover:border-white/30 shadow-lg text-lg"
                            >
                                Get Early Access
                            </button>
                            <button
                                className="text-white px-8 py-4 rounded-full bg-white/15 hover:bg-white/25 transition-all duration-300 font-medium backdrop-blur-sm border border-white/20 hover:border-white/30 shadow-lg text-lg"
                            >
                                View Demo
                            </button>
                        </div>

                    </div>

                    <div className="w-[40%] absolute bottom-20 left-1/2 transform -translate-x-1/2 text-white text-sm flex flex-row justify-between items-center gap-2">
                        <div>
                            <p>Portfolio Intelligence</p>
                        </div>

                        <div>
                            <p>Insight Engine</p>
                        </div>

                        <div>
                            <p>Cross Asset Analytics</p>
                        </div>
                    </div>

                </div>

                {/* Why - Section */}
                {/* <div className="w-screen h-screen">
                    <div className="h-full w-full flex flex-col justify-between items-start">
                        <div className="flex-[40%]">
                            <div className="w-full  h-full flex flex-col justify-evenly items-start gap-6 px-16 py-10">
                                <div>
                                    <h1 className="text-highlight-20 text-4xl text-center font-bold">
                                        Why Use <span className="font-bold">Insights</span> ?
                                    </h1>
                                </div>
                                <div className="w-full flex flex-col justify-center items-start">
                                    <h1 className="w-[90%] text-highlight-20 text-3xl font-bold text-justify">
                                        Over 90% of investors underperform the market due to information gaps, with the average investor missing out on 3-5% annual returns compared to institutional investors who have access to premium deal flow and market intelligence.
                                    </h1>
                                    <span className="px-6 py-3 border-2 border-white/30 rounded-full text-center self-end">
                                        <h1 className="text-2xl font-bold text-white">
                                            <span className="font-bold">Insights</span> intend to change that.
                                        </h1>
                                    </span>
                                </div>
                            </div>
                        </div>
                        <div className="flex-[60%] w-full h-full bg-highlight-50 flex flex-col items-center justify-center gap-16 py-5 p-16">
                            <div className="w-full flex flex-col justify-center items-end">
                                <h1 className="w-[90%] text-highlight-20 text-2xl font-bold text-right mb-8">
                                    The modern investment landscape is fragmented. Opportunities are scattered across stocks, bonds, crypto, REITs, and alternative assets each living in separate platforms with disconnected data. You&apos;re not just missing opportunities you&apos;re missing the connections between them.
                                </h1>
                                <span className="px-6 py-3 border-2 border-white/30 rounded-full text-center ml-[10%]">
                                    <h1 className="text-2xl font-bold text-white">
                                        <span className="font-bold">Insights</span> unifies it all. One platform for every asset type.
                                    </h1>
                                </span>
                            </div>

                            <div className="w-full flex flex-col justify-center items-start">
                                <h1 className="w-[90%] text-highlight-20 text-2xl font-bold text-left mb-8">
                                    Emerging markets move fast. High-potential assets don&apos;t wait for you to discover them.
                                    Without real-time intelligence, you&apos;re always one step behind the smart money.
                                    Every delayed decision is a missed opportunity.
                                </h1>
                                <span className="px-6 py-3 border-2 border-white/30 rounded-full text-center mr-[10%]">
                                    <h1 className="text-2xl font-bold text-white">
                                        <span className="font-bold">Insights</span> helps you discover hidden opportunities.
                                    </h1>
                                </span>
                            </div>
                        </div>
                    </div>
                </div> */}

                {/* Information Section */}
                <div className="w-screen h-screen flex justify-center items-center">
                    <div className="w-[90%] h-[90%] flex flex-col justify-evenly items-center ">
                        <div className="w-full h-full flex flex-row justify-between items-center gap-6">
                            {/* Card 1 */}
                            <div
                                className="w-[30%] h-[80%] rounded-2xl shadow-sm flex flex-col justify-start items-start p-4 gap-8 relative cursor-pointer"
                                style={{
                                    backgroundImage: 'url("/assets/img/Explore X Image (1).jpeg")',
                                    backgroundSize: 'cover',
                                    backgroundPosition: 'center',
                                }}
                            >
                                <div className="absolute inset-0 bg-yellow-800 opacity-50 rounded-2xl"></div>
                                <span className="m-2 text-white px-6 py-2 rounded-full bg-white/15 hover:bg-white/25 transition-all duration-300 font-medium backdrop-blur-sm border border-white/20 hover:border-white/30 shadow-lg glow-top-left glow-bottom-right z-10">
                                    <h2 className="text-xl font-bold">Smart Portfolio Management</h2>
                                </span>
                                <p className="w-[85%] mx-auto text-xl text-white font-semibold text-center z-10 my-[30%]">
                                    Insight enables you to monitor performance across stocks and crypto with intelligent analysis that spots weaknesses before they become losses. Get specific recommendations based on your risk profile and goals
                                </p>
                            </div>

                            {/* Card 2 */}
                            <div className="w-[30%] h-[80%] bg-highlight-50 rounded-2xl shadow-sm flex flex-col justify-start items-start p-4 gap-8 cursor-pointer">
                                <span className="m-2 text-white px-6 py-2 rounded-full bg-white/15 hover:bg-white/25 transition-all duration-300 font-medium backdrop-blur-sm border border-white/20 hover:border-white/30 shadow-lg glow-top-left glow-bottom-right">
                                    <h2 className="text-xl font-bold">Insight Engine</h2>
                                </span>
                                <p className="w-[85%] mx-auto text-xl text-white font-semibold text-center my-[30%] z-10">
                                    We are building an intelligent engine continuously recalibrates recommendations based on evolving market dynamics and your personal investment profile—ensuring you&apos;re always positioned for optimal outcomes, not yesterday&apos;s opportunities.
                                </p>
                            </div>

                            {/* Card 3 */}
                            <div
                                className="w-[30%] h-[80%] rounded-2xl shadow-sm flex flex-col justify-start items-start p-4 gap-8 relative overflow-hidden cursor-pointer"
                            >
                                <div
                                    className="absolute inset-0 bg-cover bg-center rounded-2xl"
                                    style={{
                                        backgroundImage: 'url("/assets/img/Explore X Image (1).jpeg")',
                                        transform: 'scaleX(-1)',
                                    }}
                                />

                                <div className="absolute inset-0 bg-yellow-800 opacity-50 rounded-2xl"></div>

                                <span className="m-2 text-white px-6 py-2 rounded-full bg-white/15 hover:bg-white/25 transition-all duration-300 font-medium backdrop-blur-sm border border-white/20 hover:border-white/30 shadow-lg glow-top-left glow-bottom-right z-10">
                                    <h2 className="text-xl font-bold">Cross Asset Analytics</h2>
                                </span>

                                <p className="w-[85%] mx-auto text-xl text-white font-semibold text-center my-[30%] z-10">
                                    Traditional platforms show you trees. We intend to show you the forest. Understand how macro trends flow between traditional equities and digital assets—spotting arbitrage opportunities, rotation signals, and correlation breakdowns that create alpha for informed investors willing to think beyond silos.
                                </p>
                            </div>


                        </div>
                    </div>
                </div>

                {/* What we do */}
                <div className="w-screen h-screen flex justify-center items-center">
                    <div className="w-full h-full flex flex-row justify-between items-end">
                        <div className="flex-2/3 w-full h-full flex flex-col justify-center items-start gap-6 p-16">
                            <h1 className="text-highlight-20 text-6xl text-left font-semibold">
                                Bridging the Investment Intelligence Gap
                            </h1>
                            <p className="text-lg mt-6 text-white/80 text-justify">
                                The modern investment landscape is broken. While institutional investors leverage sophisticated analytics and real-time <span className="font-bold">Insights</span> across multiple asset classes, individual investors are stuck navigating fragmented platforms, delayed data, and generic advice. We watched talented investors make suboptimal decisions not because they lacked skill, but because they lacked the same quality of information and analysis available to Wall Street.
                            </p>
                            <p className="text-lg mt-6 text-white/80 text-justify">
                                <span className="font-bold">Insights</span> was born from a simple belief: every investor deserves institutional-grade intelligence. We&apos;re building the first platform that unifies stocks, bonds, and crypto analytics with prescriptive AI that doesn&apos;t just show you what&apos;s happening—it tells you what to do about it. Our vision is to democratize investment intelligence, giving individual investors the same analytical edge that institutions have enjoyed for decades.
                            </p>
                            <p className="text-lg mt-6 text-white/80 text-justify">
                                We are committed to democratizing access to premium market intelligence, enabling every investor to compete on a level playing field with institutional players.
                            </p>
                            <p className="text-lg mt-6 text-white/80 text-justify">
                                We achieve this through three core pillars: unified cross-asset data that reveals hidden opportunities, personalized AI that adapts to your unique goals and risk profile, and prescriptive analytics that move beyond predictions to actionable recommendations. Every feature we build is designed to close the information gap between you and institutional investors.
                            </p>
                        </div>

                        <div className="flex-1/3 w-full h-full flex items-end">
                            <div className="w-full h-[90%] flex flex-row justify-start items-stretch gap-6 px-10 py-16">
                                {/* Step indicator */}
                                <div className="relative flex flex-col items-center justify-center min-h-full">
                                    {/* Vertical gradient line */}
                                    <div
                                        className="absolute left-1/2 transform -translate-x-1/2 w-6 h-[90%] top-[5%]"
                                        style={{
                                            background: 'linear-gradient(180deg, #1f2937 0%, #8a7557 50%, #ea580c 100%)'
                                        }}
                                    />

                                    {/* Step indicators */}
                                    <div className="relative z-10 flex flex-col justify-between h-full py-4">
                                        {/* Step 1 */}
                                        <div className="flex items-center justify-center w-12 h-12 bg-white rounded-full shadow-lg">
                                            <span className="text-black font-bold text-lg">1</span>
                                        </div>

                                        {/* Step 2 */}
                                        <div className="flex items-center justify-center w-12 h-12 bg-white rounded-full shadow-lg">
                                            <span className="text-black font-bold text-lg">2</span>
                                        </div>

                                        {/* Step 3 */}
                                        <div className="flex items-center justify-center w-12 h-12 bg-white rounded-full shadow-lg">
                                            <span className="text-black font-bold text-lg">3</span>
                                        </div>
                                    </div>
                                </div>

                                {/* Steps content */}
                                <div className="w-[80%] h-full flex flex-col justify-between items-start">
                                    <div className="flex-1 flex flex-col justify-center items-start">
                                        <h1 className="text-sm font-medium text-gray-400 mb-2">STEP 01</h1>
                                        <h3 className="text-highlight-20 text-3xl font-bold mb-4">
                                            See It in Action
                                        </h3>
                                        <p className="mb-6 text-gray-300 leading-relaxed">
                                            Experience the power of <span className="font-bold">Insights</span> firsthand. Sign up for our early access program to explore our platform and see how it can transform your investment strategy.
                                        </p>
                                        <button>
                                            <span className="text-white px-4 py-2 rounded-full bg-white/15 hover:bg-white/25 transition-all duration-300 font-medium backdrop-blur-sm border border-white/20 hover:border-white/30 shadow-lg text-lg">
                                                View Demo
                                            </span>
                                        </button>
                                    </div>

                                    <div className="flex-1 flex flex-col justify-center items-start">
                                        <h1 className="text-sm font-medium text-gray-400 mb-2">STEP 02</h1>
                                        <h3 className="text-highlight-20 text-3xl font-bold mb-4">
                                            Secure Your Beta Access
                                        </h3>
                                        <p className="mb-6 text-gray-300 leading-relaxed">
                                            Join our exclusive beta program to be among the first to experience <span className="font-bold">Insights</span>. Gain early access to our platform and help shape its future with your feedback.
                                        </p>
                                        <button>
                                            <span className="text-white px-4 py-2 rounded-full bg-white/15 hover:bg-white/25 transition-all duration-300 font-medium backdrop-blur-sm border border-white/20 hover:border-white/30 shadow-lg text-lg">
                                                Get Early Access
                                            </span>
                                        </button>
                                    </div>

                                    <div className="flex-1 flex flex-col justify-center">
                                        <h1 className="text-sm font-medium text-gray-400 mb-2">STEP 03</h1>
                                        <h3 className="text-highlight-20 text-3xl font-bold mb-4">
                                            Start Investing Smarter
                                        </h3>
                                        <p className="text-gray-300 leading-relaxed">
                                            Get immediate access to unified market data, personalized insights, and prescriptive recommendations. No more platform hopping or guesswork—just clear, actionable intelligence.
                                        </p>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>

                {/* Footer */}
                <div className="w-screen h-auto bg-highlight-50 flex flex-col justify-center items-center py-16">
                    {/* Main Content Section */}
                    <div className="w-[90%] flex flex-row justify-between items-start gap-12 mb-16">
                        {/* Left Side - Main Message */}
                        <div className="text-white w-[50%] flex flex-col justify-between items-between gap-6">
                            <h1 className=" text-4xl font-bold leading-tight">
                                Insights intend to be the first platform that unifies cross-asset investing with real-time intelligence.
                            </h1>
                            <p className="mt-10">Projections for 2025 includes:  </p>
                        </div>

                        {/* Right Side - Description & CTA */}
                        <div className="w-[40%] flex flex-col justify-start items-start gap-6">
                            <p className="text-white text-lg leading-relaxed">
                            Starting with our beta community of finance professionals, we&apos;re expanding to serve individual investors worldwide. Our roadmap includes AI-powered risk assessment, advanced portfolio optimization, and expanded asset class coverage.
                            </p>
                            <button className="mt-10 text-white px-4 py-2 rounded-full bg-white/10 hover:bg-white/20 transition-all duration-300 font-medium backdrop-blur-sm border border-white/20 hover:border-white/30 shadow-lg text-lg">
                                Schedule a Call
                            </button>
                        </div>
                    </div>

                    {/* Statistics Section */}
                    <div className="w-[90%] grid grid-cols-4 gap-8">
                        {/* Stat 1 */}
                        <div className="flex flex-col items-center text-center">
                            <h2 className="text-6xl font-bold text-white mb-2">1K+</h2>
                            <p className="text-white text-lg font-medium">Active Users</p>
                        </div>

                        {/* Stat 2 */}
                        <div className="flex flex-col items-center text-center">
                            <h2 className="text-6xl font-bold text-white mb-2">$2M+</h2>
                            <p className="text-white text-lg font-medium">Assets Under Management</p>
                        </div>

                        {/* Stat 3 */}
                        <div className="flex flex-col items-center text-center">
                            <h2 className="text-6xl font-bold text-white mb-2">7+</h2>
                            <p className="text-white text-lg font-medium">Asset Classes</p>
                        </div>

                        {/* Stat 4 */}
                        <div className="flex flex-col items-center text-center">
                            <h2 className="text-6xl font-bold text-white mb-2">24/7</h2>
                            <p className="text-white text-lg font-medium">Market Coverage</p>
                        </div>
                    </div>

                    {/* Bottom Section - Additional Info */}
                    <div className="w-[90%] flex flex-row justify-between items-center mt-16 pt-8 border-t border-white/10">
                        {/* Company Info */}
                        <div className="flex flex-col gap-2">
                            <h3 className="text-2xl font-bold text-highlight-20">Insights</h3>
                            <p className="text-white">Unifying the investment landscape</p>
                        </div>

                        {/* Quick Links */}
                        <div className="flex flex-row gap-8">
                            <div className="flex flex-col gap-2">
                                <h4 className="text-white font-semibold mb-2">Platform</h4>
                                <a href="#" className="text-white hover:text-white/80 transition-colors">Features</a>
                                <a href="#" className="text-white hover:text-white/80 transition-colors">Pricing</a>
                                <a href="#" className="text-white hover:text-white/80 transition-colors">API</a>
                            </div>
                            <div className="flex flex-col gap-2">
                                <h4 className="text-white font-semibold mb-2">Company</h4>
                                <a href="#" className="text-white hover:text-white/80 transition-colors">About</a>
                                <a href="#" className="text-white hover:text-white/80 transition-colors">Careers</a>
                                <a href="#" className="text-white hover:text-white/80 transition-colors">Contact</a>
                            </div>
                            <div className="flex flex-col gap-2">
                                <h4 className="text-white font-semibold mb-2">Resources</h4>
                                <a href="#" className="text-white hover:text-white/80 transition-colors">Blog</a>
                                <a href="#" className="text-white hover:text-white/80 transition-colors">Research</a>
                                <a href="#" className="text-white hover:text-white/80 transition-colors">Support</a>
                            </div>
                        </div>

                        {/* Social Links */}
                        <div className="flex flex-col gap-2">
                            <h4 className="text-white font-semibold mb-2">Connect</h4>
                            <div className="flex gap-4">
                                <a href="#" className="text-white hover:text-white/80 transition-colors">LinkedIn</a>
                                <a href="#" className="text-white hover:text-white/80 transition-colors">Twitter</a>
                                <a href="#" className="text-white hover:text-white/80 transition-colors">GitHub</a>
                            </div>
                        </div>
                    </div>

                    {/* Copyright */}
                    <div className="w-[90%] flex justify-center items-center mt-8 pt-4 border-t border-white">
                        <p className="text-white/40 text-sm">
                            © 2025 Insights. All rights reserved. | Privacy Policy | Terms of Service
                        </p>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default page
