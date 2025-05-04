"use client"

import React from 'react'
import Image from "next/image"

const SubHero = () => {
  return (
    <section className="w-full h-full relative">
            <h1 className="absolute bottom-10 right-10 text-4xl font-bold text-slate-50 z-10">
                DYOR
            </h1>
                 <Image 
                   src="https://ik.imagekit.io/0y99xuz0yp/Aerial%20Farm%20View.png?updatedAt=1745512149498" 
                   alt="Hero-image" 
                   layout="fill" 
                   objectFit="cover" 
                 />
        </section>
  )
}

export default SubHero
