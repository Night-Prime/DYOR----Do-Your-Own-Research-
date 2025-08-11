import React, { useEffect, useRef } from 'react';

import Navbar from "../components/Navbar";
import { motion, useInView, useMotionValue, useSpring, useTransform, Variants } from "framer-motion";
import { useLenis } from '../hooks/useLenis';

interface StatItemProps {
    value: string;
    label: string;
    delay?: number;
}

const AnimatedCounter: React.FC<{ value: string; delay?: number }> = ({ value, delay = 0 }) => {
    const ref = useRef<HTMLHeadingElement>(null);
    const isInView = useInView(ref, { once: true });

    // Extract numeric part for animation
    const numericValue = value.match(/\d+/)?.[0] || "0";
    const suffix = value.replace(numericValue, "");

    const motionValue = useMotionValue(0);
    const springValue = useSpring(motionValue, { duration: 2000 });
    const displayValue = useTransform(springValue, (latest) =>
        Math.round(latest) + suffix
    );

    useEffect(() => {
        if (isInView) {
            setTimeout(() => {
                motionValue.set(parseInt(numericValue));
            }, delay);
        }
    }, [isInView, motionValue, numericValue, delay]);

    return (
        <motion.h2
            ref={ref}
            className="text-6xl font-bold text-white mb-2"
            initial={{ opacity: 0, scale: 0.5 }}
            animate={isInView ? { opacity: 1, scale: 1 } : {}}
            transition={{ duration: 0.8, delay: delay / 1000, type: "spring", bounce: 0.3 }}
        >
            {displayValue}
        </motion.h2>
    );
};

const StatItem: React.FC<StatItemProps> = ({ value, label, delay = 0 }) => {
    return (
        <motion.div
            className="flex flex-col items-center text-center"
            initial={{ opacity: 0, y: 50 }}
            whileInView={{ opacity: 1, y: 0 }}
            viewport={{ once: true, amount: 0.5 }}
            transition={{
                duration: 0.8,
                delay: delay / 1000,
                ease: [0.25, 0.25, 0.25, 0.75]
            }}
        >
            <AnimatedCounter value={value} delay={delay} />
            <motion.p
                className="text-white text-lg font-medium"
                initial={{ opacity: 0 }}
                whileInView={{ opacity: 1 }}
                viewport={{ once: true }}
                transition={{ duration: 0.6, delay: (delay + 400) / 1000 }}
            >
                {label}
            </motion.p>
        </motion.div>
    );
};

const Homepage = () => {
    useLenis();
    const container = {
        hidden: { opacity: 0 },
        show: {
            opacity: 1,
            transition: {
                staggerChildren: .25, // delay between each child
            },
        },
    };

    const item: Variants = {
        hidden: { opacity: 0, y: 50 },
        show: { opacity: 1, y: 0, transition: { duration: .5, ease: "easeOut" } },
    };


    const cardContainer: Variants = {
        hidden: {},
        show: {
            transition: {
                staggerChildren: 0.25,
                delayChildren: 0.1, 
            },
        },
    };

    const card: Variants = {
        hidden: {
            y: 80,        // Reduced from 300 for subtler movement
            opacity: 0,
            scale: 0.9,   // Slight scale for extra smoothness
            rotateX: 15,  // Subtle 3D rotation
        },
        show: {
            y: 0,
            opacity: 1,
            scale: 1,
            rotateX: 0,
            transition: {
                type: "spring",
                stiffness: 100,      // Lower stiffness for smoother motion
                damping: 20,         // Higher damping reduces oscillation
                mass: 1,             // Controls weight feel
                duration: 1.2,       // Slightly longer for elegance
                ease: [0.25, 0.46, 0.45, 0.94], // Custom cubic-bezier for silk-smooth motion
            },
        },
    };

    const lineVariants: Variants = {
        hidden: {
            opacity: 0,
            y: 30, // Start 30px below
        },
        visible: {
            opacity: 1,
            y: 0,
            transition: {
                duration: 0.8,
                ease: [0.25, 0.25, 0.25, 0.75], // Custom easing
            },
        },
    };

    const itemVariants: Variants = {
        hidden: {
            opacity: 0,
            y: 30,
        },
        visible: {
            opacity: 1,
            y: 0,
            transition: {
                duration: 0.8,
                ease: [0.25, 0.25, 0.25, 0.75],
            },
        },
    };

    const linkColumnVariants: Variants = {
        hidden: {},
        visible: {
            transition: {
                staggerChildren: 0.1,
            },
        },
    };

    const linkVariants: Variants = {
        hidden: {
            opacity: 0,
            x: -20,
        },
        visible: {
            opacity: 1,
            x: 0,
            transition: {
                duration: 0.5,
            },
        },
    }

    // Split text into lines for animation
    const textLines = [
        "The modern investment landscape is broken. While institutional investors leverage sophisticated analytics and real-time Insights across multiple asset classes, individual investors are stuck navigating fragmented platforms, delayed data, and generic advice. We watched talented investors make suboptimal decisions not because they lacked skill, but because they lacked the same quality of information and analysis available to Wall Street.",

        "Insights was born from a simple belief: every investor deserves institutional-grade intelligence. We're building the first platform that unifies stocks, bonds, and crypto analytics with prescriptive AI that doesn't just show you what's happening—it tells you what to do about it. Our vision is to democratize investment intelligence, giving individual investors the same analytical edge that institutions have enjoyed for decades.",

        "We are committed to democratizing access to premium market intelligence, enabling every investor to compete on a level playing field with institutional players.",

        "We achieve this through three core pillars: unified cross-asset data that reveals hidden opportunities, personalized AI that adapts to your unique goals and risk profile, and prescriptive analytics that move beyond predictions to actionable recommendations. Every feature we build is designed to close the information gap between you and institutional investors."
    ];

    return (
        <div className="w-full h-auto bg-zinc-950">
            <div className="w-full h-full flex flex-col justify-center items-center gap-10">
                <Navbar />
                {/* Hero-section */}
                <motion.div
                    variants={container}
                    initial="hidden"
                    animate="show"
                    transition={{ delayChildren: .1, staggerChildren: 0.1 }}
                    className="w-screen h-screen flex justify-center items-center relative"
                    style={{
                        backgroundImage: 'url("/assets/img/Gemini_Generated_Image_m8cxeem8cxeem8cx.png")',
                        backgroundSize: 'cover',
                        backgroundPosition: 'center',
                    }}
                >
                    <div className="absolute inset-0 bg-black opacity-[40%]"></div>

                    <motion.div
                        className="h-[50%] w-[80%] m-auto text-center flex flex-col justify-center items-center gap-4 z-10"
                    >
                        <motion.h1
                            variants={item}
                            className="text-7xl font-bold text-white"
                        >
                            Make Smarter Investment Decisions
                        </motion.h1>

                        <motion.p
                            variants={item}
                            className="text-lg mt-6 font-semibold text-white/80"
                        >
                            Comprehensive market data, portfolio analytics, and cross-asset{" "}
                            <span className="font-bold">Insights</span> in one powerful platform
                            built for investors and finance professionals.
                        </motion.p>

                        <motion.div
                            variants={item}
                            className="w-[60%] flex flex-row justify-center items-center gap-4 mt-6"
                        >
                            <button className="text-white px-8 py-4 rounded-full bg-white/15 hover:bg-white/25 transition-all duration-300 font-medium backdrop-blur-sm border border-white/20 hover:border-white/30 shadow-lg text-lg">
                                Get Early Access
                            </button>
                            <button className="text-white px-8 py-4 rounded-full bg-white/15 hover:bg-white/25 transition-all duration-300 font-medium backdrop-blur-sm border border-white/20 hover:border-white/30 shadow-lg text-lg">
                                View Demo
                            </button>
                        </motion.div>
                    </motion.div>

                    {/* Bottom labels */}
                    <motion.div
                        className="w-[40%] absolute bottom-20 left-1/2 transform -translate-x-1/2 text-white text-sm flex flex-row justify-between items-center gap-2"
                    >
                        <motion.div variants={item}>
                            <p>Portfolio Intelligence</p>
                        </motion.div>
                        <motion.div variants={item}>
                            <p>Insight Engine</p>
                        </motion.div>
                        <motion.div variants={item}>
                            <p>Cross Asset Analytics</p>
                        </motion.div>
                    </motion.div>

                </motion.div>

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
            <motion.div
                variants={cardContainer}
                initial="hidden"
                whileInView="show"
                viewport={{ once: true, amount: 0.2 }}
                className="w-[90%] h-[90%] flex flex-col justify-evenly items-center"
            >
                <div className="w-full h-full flex flex-row justify-between items-center gap-6">
                    {/* Card 1 */}
                    <motion.div
                        variants={card}
                        className="w-[30%] h-[80%] rounded-2xl shadow-sm flex flex-col justify-start items-start p-4 gap-8 relative cursor-pointer"
                        style={{
                            backgroundImage: 'url("/assets/img/Explore X Image (1).jpeg")',
                            backgroundSize: "cover",
                            backgroundPosition: "center",
                        }}
                        whileHover={{
                            y: -10,
                            scale: 1.02,
                            transition: { duration: 0.3, ease: "easeOut" }
                        }}
                    >
                        <div className="absolute inset-0 bg-yellow-800 opacity-50 rounded-2xl"></div>
                        <span className="m-2 text-white px-6 py-2 rounded-full bg-white/15 hover:bg-white/25 transition-all duration-300 font-medium backdrop-blur-sm border border-white/20 hover:border-white/30 shadow-lg glow-top-left glow-bottom-right z-10">
                            <h2 className="text-xl font-bold">Smart Portfolio Management</h2>
                        </span>
                        <p className="w-[85%] mx-auto text-xl text-white font-semibold text-center z-10 my-[30%]">
                            Insight enables you to monitor performance across stocks and crypto with intelligent analysis that spots weaknesses before they become losses. Get specific recommendations based on your risk profile and goals.
                        </p>
                    </motion.div>

                    {/* Card 2 */}
                    <motion.div
                        variants={card}
                        className="w-[30%] h-[80%] bg-highlight-50 rounded-2xl shadow-sm flex flex-col justify-start items-start p-4 gap-8 cursor-pointer"
                        whileHover={{
                            y: -10,
                            scale: 1.02,
                            transition: { duration: 0.3, ease: "easeOut" }
                        }}
                    >
                        <span className="m-2 text-white px-6 py-2 rounded-full bg-white/15 hover:bg-white/25 transition-all duration-300 font-medium backdrop-blur-sm border border-white/20 hover:border-white/30 shadow-lg glow-top-left glow-bottom-right">
                            <h2 className="text-xl font-bold">Insight Engine</h2>
                        </span>
                        <p className="w-[85%] mx-auto text-xl text-white font-semibold text-center my-[30%] z-10">
                            We are building an intelligent engine that continuously recalibrates recommendations based on evolving market dynamics and your personal investment profile—ensuring you&apos;re always positioned for optimal outcomes, not yesterday&apos;s opportunities.
                        </p>
                    </motion.div>

                    {/* Card 3 */}
                    <motion.div
                        variants={card}
                        className="w-[30%] h-[80%] rounded-2xl shadow-sm flex flex-col justify-start items-start p-4 gap-8 relative overflow-hidden cursor-pointer"
                        whileHover={{
                            y: -10,
                            scale: 1.02,
                            transition: { duration: 0.3, ease: "easeOut" }
                        }}
                    >
                        <div
                            className="absolute inset-0 bg-cover bg-center rounded-2xl"
                            style={{
                                backgroundImage: 'url("/assets/img/Explore X Image (1).jpeg")',
                                transform: "scaleX(-1)",
                            }}
                        />
                        <div className="absolute inset-0 bg-yellow-800 opacity-50 rounded-2xl"></div>
                        <span className="m-2 text-white px-6 py-2 rounded-full bg-white/15 hover:bg-white/25 transition-all duration-300 font-medium backdrop-blur-sm border border-white/20 hover:border-white/30 shadow-lg glow-top-left glow-bottom-right z-10">
                            <h2 className="text-xl font-bold">Cross Asset Analytics</h2>
                        </span>
                        <p className="w-[85%] mx-auto text-xl text-white font-semibold text-center my-[30%] z-10">
                            Traditional platforms show you trees. We intend to show you the forest. Understand how macro trends flow between traditional equities and digital assets—spotting arbitrage opportunities, rotation signals, and correlation breakdowns that create alpha for informed investors willing to think beyond silos.
                        </p>
                    </motion.div>
                </div>
            </motion.div>
        </div>

                {/* What we do */}
                <div className="w-screen h-screen flex justify-center items-center">
                    <div className="w-full h-full flex flex-row justify-between items-end">
                        <div className="flex-2/3 w-full h-full flex flex-col justify-center items-start gap-6 p-16">
                            {/* Animated heading */}
                            <motion.h1
                                className="text-highlight-20 text-6xl text-left font-semibold"
                                initial={{ opacity: 0, y: 50 }}
                                whileInView={{ opacity: 1, y: 0 }}
                                viewport={{ once: true, amount: 0.3 }}
                                transition={{ duration: 1, ease: "easeOut" }}
                            >
                                Bridging the Investment Intelligence Gap
                            </motion.h1>

                            {/* Animated text lines */}
                            <motion.div
                                variants={cardContainer}
                                initial="hidden"
                                whileInView="visible"
                                viewport={{ once: true, amount: 0.3 }}
                                className="space-y-6"
                            >
                                {textLines.map((line, index) => (
                                    <motion.p
                                        key={index}
                                        variants={lineVariants}
                                        className="text-lg text-white/80 text-justify"
                                    >
                                        {index === 0 && (
                                            <>
                                                The modern investment landscape is broken. While institutional investors leverage sophisticated analytics and real-time <span className="font-bold">Insights</span> across multiple asset classes, individual investors are stuck navigating fragmented platforms, delayed data, and generic advice. We watched talented investors make suboptimal decisions not because they lacked skill, but because they lacked the same quality of information and analysis available to Wall Street.
                                            </>
                                        )}
                                        {index === 1 && (
                                            <>
                                                <span className="font-bold">Insights</span> was born from a simple belief: every investor deserves institutional-grade intelligence. We&apos;re building the first platform that unifies stocks, bonds, and crypto analytics with prescriptive AI that doesn&apos;t just show you what&apos;s happening—it tells you what to do about it. Our vision is to democratize investment intelligence, giving individual investors the same analytical edge that institutions have enjoyed for decades.
                                            </>
                                        )}
                                        {index === 2 && line}
                                        {index === 3 && line}
                                    </motion.p>
                                ))}
                            </motion.div>
                        </div>

                        {/* Right side with steps - also animated */}
                        <div className="flex-1/3 w-full h-full flex items-end">
                            <motion.div
                                className="w-full h-[90%] flex flex-row justify-start items-stretch gap-6 px-10 py-16"
                                initial={{ opacity: 0, x: 100 }}
                                whileInView={{ opacity: 1, x: 0 }}
                                viewport={{ once: true, amount: 0.3 }}
                                transition={{ duration: 1, delay: 0.5, ease: "easeOut" }}
                            >
                                {/* Step indicator */}
                                <div className="relative flex flex-col items-center justify-center min-h-full">
                                    {/* Vertical gradient line */}
                                    <motion.div
                                        className="absolute left-1/2 transform -translate-x-1/2 w-6 top-[5%]"
                                        style={{
                                            background: 'linear-gradient(180deg, #1f2937 0%, #8a7557 50%, #ea580c 100%)'
                                        }}
                                        initial={{ height: 0 }}
                                        whileInView={{ height: '90%' }}
                                        viewport={{ once: true }}
                                        transition={{ duration: 1.5, delay: 0.8 }}
                                    />

                                    {/* Step indicators */}
                                    <div className="relative z-10 flex flex-col justify-between h-full py-4">
                                        {[1, 2, 3].map((step, index) => (
                                            <motion.div
                                                key={step}
                                                className="flex items-center justify-center w-12 h-12 bg-white rounded-full shadow-lg"
                                                initial={{ scale: 0, opacity: 0 }}
                                                whileInView={{ scale: 1, opacity: 1 }}
                                                viewport={{ once: true }}
                                                transition={{
                                                    duration: 0.5,
                                                    delay: 1.2 + (index * 0.2),
                                                    type: "spring",
                                                    bounce: 0.4
                                                }}
                                            >
                                                <span className="text-black font-bold text-lg">{step}</span>
                                            </motion.div>
                                        ))}
                                    </div>
                                </div>

                                {/* Steps content */}
                                <motion.div
                                    className="w-[80%] h-full flex flex-col justify-between items-start"
                                    variants={cardContainer}
                                    initial="hidden"
                                    whileInView="visible"
                                    viewport={{ once: true, amount: 0.3 }}
                                >
                                    <motion.div variants={lineVariants} className="flex-1 flex flex-col justify-center items-start">
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
                                    </motion.div>

                                    <motion.div variants={lineVariants} className="flex-1 flex flex-col justify-center items-start">
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
                                    </motion.div>

                                    <motion.div variants={lineVariants} className="flex-1 flex flex-col justify-center">
                                        <h1 className="text-sm font-medium text-gray-400 mb-2">STEP 03</h1>
                                        <h3 className="text-highlight-20 text-3xl font-bold mb-4">
                                            Start Investing Smarter
                                        </h3>
                                        <p className="text-gray-300 leading-relaxed">
                                            Get immediate access to unified market data, personalized insights, and prescriptive recommendations. No more platform hopping or guesswork—just clear, actionable intelligence.
                                        </p>
                                    </motion.div>
                                </motion.div>
                            </motion.div>
                        </div>
                    </div>
                </div>

                {/* Footer */}
                <div className="w-screen h-auto bg-highlight-50 flex flex-col justify-center items-center py-16">
                    {/* Main Content Section */}
                    <motion.div
                        className="w-[90%] flex flex-row justify-between items-start gap-12 mb-16"
                        variants={cardContainer}
                        initial="hidden"
                        whileInView="visible"
                        viewport={{ once: true, amount: 0.3 }}
                    >
                        {/* Left Side - Main Message */}
                        <div className="text-white w-[50%] flex flex-col justify-between items-between gap-6">
                            <motion.h1
                                variants={itemVariants}
                                className="text-4xl font-bold leading-tight"
                            >
                                Insights intend to be the first platform that unifies cross-asset investing with real-time intelligence.
                            </motion.h1>
                            <motion.p
                                variants={itemVariants}
                                className="mt-10"
                            >
                                Projections for 2025 includes:
                            </motion.p>
                        </div>

                        {/* Right Side - Description & CTA */}
                        <motion.div
                            className="w-[40%] flex flex-col justify-start items-start gap-6"
                            variants={cardContainer}
                        >
                            <motion.p
                                variants={itemVariants}
                                className="text-white text-lg leading-relaxed"
                            >
                                Starting with our beta community of finance professionals, we&apos;re expanding to serve individual investors worldwide. Our roadmap includes AI-powered risk assessment, advanced portfolio optimization, and expanded asset class coverage.
                            </motion.p>
                            <motion.button
                                variants={itemVariants}
                                className="mt-10 text-white px-4 py-2 rounded-full bg-white/10 hover:bg-white/20 transition-all duration-300 font-medium backdrop-blur-sm border border-white/20 hover:border-white/30 shadow-lg text-lg"
                                whileHover={{ scale: 1.05 }}
                                whileTap={{ scale: 0.95 }}
                            >
                                Schedule a Call
                            </motion.button>
                        </motion.div>
                    </motion.div>

                    {/* Statistics Section */}
                    <motion.div
                        className="w-[90%] grid grid-cols-4 gap-8"
                        initial={{ opacity: 0 }}
                        whileInView={{ opacity: 1 }}
                        viewport={{ once: true, amount: 0.2 }}
                        transition={{ duration: 0.6 }}
                    >
                        <StatItem value="1K+" label="Active Users" delay={200} />
                        <StatItem value="2M+" label="Assets Under Management" delay={400} />
                        <StatItem value="7+" label="Asset Classes" delay={600} />
                        <StatItem value="24/7" label="Market Coverage" delay={800} />
                    </motion.div>

                    {/* Bottom Section - Additional Info */}
                    <motion.div
                        className="w-[90%] flex flex-row justify-between items-center mt-16 pt-8 border-t border-white/10"
                        initial={{ opacity: 0, y: 50 }}
                        whileInView={{ opacity: 1, y: 0 }}
                        viewport={{ once: true, amount: 0.3 }}
                        transition={{ duration: 1, delay: 0.2 }}
                    >
                        {/* Company Info */}
                        <motion.div
                            className="flex flex-col gap-2"
                            initial={{ opacity: 0, x: -30 }}
                            whileInView={{ opacity: 1, x: 0 }}
                            viewport={{ once: true }}
                            transition={{ duration: 0.8, delay: 0.4 }}
                        >
                            <h3 className="text-2xl font-bold text-highlight-20">Insights</h3>
                            <p className="text-white">Unifying the investment landscape</p>
                        </motion.div>

                        {/* Quick Links */}
                        <motion.div
                            className="flex flex-row gap-8"
                            variants={linkColumnVariants}
                            initial="hidden"
                            whileInView="visible"
                            viewport={{ once: true, amount: 0.3 }}
                        >
                            <motion.div className="flex flex-col gap-2" variants={linkColumnVariants}>
                                <motion.h4 variants={linkVariants} className="text-white font-semibold mb-2">
                                    Platform
                                </motion.h4>
                                <motion.a variants={linkVariants} href="#" className="text-white hover:text-white/80 transition-colors">
                                    Features
                                </motion.a>
                                <motion.a variants={linkVariants} href="#" className="text-white hover:text-white/80 transition-colors">
                                    Pricing
                                </motion.a>
                                <motion.a variants={linkVariants} href="#" className="text-white hover:text-white/80 transition-colors">
                                    API
                                </motion.a>
                            </motion.div>

                            <motion.div className="flex flex-col gap-2" variants={linkColumnVariants}>
                                <motion.h4 variants={linkVariants} className="text-white font-semibold mb-2">
                                    Company
                                </motion.h4>
                                <motion.a variants={linkVariants} href="#" className="text-white hover:text-white/80 transition-colors">
                                    About
                                </motion.a>
                                <motion.a variants={linkVariants} href="#" className="text-white hover:text-white/80 transition-colors">
                                    Careers
                                </motion.a>
                                <motion.a variants={linkVariants} href="#" className="text-white hover:text-white/80 transition-colors">
                                    Contact
                                </motion.a>
                            </motion.div>

                            <motion.div className="flex flex-col gap-2" variants={linkColumnVariants}>
                                <motion.h4 variants={linkVariants} className="text-white font-semibold mb-2">
                                    Resources
                                </motion.h4>
                                <motion.a variants={linkVariants} href="#" className="text-white hover:text-white/80 transition-colors">
                                    Blog
                                </motion.a>
                                <motion.a variants={linkVariants} href="#" className="text-white hover:text-white/80 transition-colors">
                                    Research
                                </motion.a>
                                <motion.a variants={linkVariants} href="#" className="text-white hover:text-white/80 transition-colors">
                                    Support
                                </motion.a>
                            </motion.div>
                        </motion.div>

                        {/* Social Links */}
                        <motion.div
                            className="flex flex-col gap-2"
                            initial={{ opacity: 0, x: 30 }}
                            whileInView={{ opacity: 1, x: 0 }}
                            viewport={{ once: true }}
                            transition={{ duration: 0.8, delay: 0.6 }}
                        >
                            <h4 className="text-white font-semibold mb-2">Connect</h4>
                            <motion.div
                                className="flex gap-4"
                                variants={linkColumnVariants}
                                initial="hidden"
                                whileInView="visible"
                                viewport={{ once: true }}
                            >
                                <motion.a variants={linkVariants} href="#" className="text-white hover:text-white/80 transition-colors">
                                    LinkedIn
                                </motion.a>
                                <motion.a variants={linkVariants} href="#" className="text-white hover:text-white/80 transition-colors">
                                    Twitter
                                </motion.a>
                                <motion.a variants={linkVariants} href="#" className="text-white hover:text-white/80 transition-colors">
                                    GitHub
                                </motion.a>
                            </motion.div>
                        </motion.div>
                    </motion.div>

                    {/* Copyright */}
                    <motion.div
                        className="w-[90%] flex justify-center items-center mt-8 pt-4 border-t border-white"
                        initial={{ opacity: 0 }}
                        whileInView={{ opacity: 1 }}
                        viewport={{ once: true }}
                        transition={{ duration: 1, delay: 0.8 }}
                    >
                        <p className="text-white/40 text-sm">
                            © 2025 Insights. All rights reserved. | Privacy Policy | Terms of Service
                        </p>
                    </motion.div>
                </div>
            </div>
        </div >
    )
}

export default Homepage
