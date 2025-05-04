"use client"

import React from 'react'
import Feed from '../components/Feed'
import ProtectedRoute from '../shared/ProtectedRoute'

const page = () => {
  return (
    <ProtectedRoute>
      <div className="bg-white w-full h-screen flex flex-col justify-center items-center">
        <h1 className="text-4xl text-lime-800">
          DYOR Dashboard
          <Feed />
        </h1>
      </div>
    </ProtectedRoute>
  );
};

export default page
