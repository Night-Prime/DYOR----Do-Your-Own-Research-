"use client";

import React, { useState } from "react";
import Image from "next/image";
import { signup } from "../utils/auth";
import SuccessModal from "../shared/SucessModal";

interface RegisterProps {
  modal: () => void;
}

const Register: React.FC<RegisterProps> = ({ modal }) => {
  const [success, setSuccess] = useState<boolean>(false);

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    formData.append("role", "user");

    const result = await signup({}, formData);
    if (result.success) {
      setSuccess(true);
    } else {
      console.log("Registration failed:", result.errors);
    }
  };

  return (
    <div className="w-full h-full fixed inset-0 flex backdrop-blur-sm items-center justify-center overflow-hidden z-40">
      <div className="absolute inset-0 flex items-center justify-center">
        <Image
          src="https://ik.imagekit.io/0y99xuz0yp/Task%2001%20Image%200.png?updatedAt=1745784176302"
          alt="Hero-image"
          width={900}
          height={900}
          objectFit="contain"
        />
      </div>
      <div>
        {success ? (
          <div className="fixed inset-0 flex items-center justify-center z-50">
            <SuccessModal />
          </div>
        ) : (
          <div
            className="relative bg-white w-full max-w-md max-h-2/3 rounded-3xl p-12 shadow-lg z-10 overflow-y-auto"
            style={{ scrollbarWidth: "none", msOverflowStyle: "none" }}
          >
            <style jsx>{`
              div::-webkit-scrollbar {
                display: none;
              }
            `}</style>

            <div className="w-full flex flex-row justify-between items-center mb-4">
              <h2 className="text-2xl font-bold text-gray-800">
                Smarter trades start here.
              </h2>
              <button
                className="absolute top-2 right-5 text-red-600 text-xl font-bold hover:text-red-800 transition duration-200 focus:outline-none"
                onClick={modal}
                aria-label="Close"
              >
                &times;
              </button>
            </div>
            <form onSubmit={handleSubmit} className="flex flex-col space-y-6">
              <div className="flex flex-col space-y-2">
                <input
                  type="text"
                  id="first_name"
                  name="first_name"
                  className="rounded-md border-gray-300 shadow-sm px-4 py-2 focus:border-lime-500 focus:ring-lime-500 sm:text-sm"
                  placeholder="First Name"
                />
              </div>
              <div className="flex flex-col space-y-2">
                <input
                  type="text"
                  id="last_name"
                  name="last_name"
                  className="rounded-md border-gray-300 shadow-sm px-4 py-2 focus:border-lime-500 focus:ring-lime-500 sm:text-sm"
                  placeholder="Last Name"
                />
              </div>
              {/* <div className="flex flex-col space-y-2">
                <input
                  type="text"
                  id="avatar"
                  name="avatar"
                  className="rounded-md border-gray-300 shadow-sm px-4 py-2 focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                  placeholder="Avatar"
                />
              </div> */}
              <div className="flex flex-col space-y-2">
                <input
                  type="email"
                  id="email"
                  name="email"
                  className="rounded-md border-gray-300 shadow-sm px-4 py-2 focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                  placeholder="Email Address"
                />
              </div>
              <div className="flex flex-col space-y-2">
                <input
                  type="password"
                  id="password"
                  name="password"
                  className="rounded-md border-gray-300 shadow-sm px-4 py-2 focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                  placeholder="Password"
                />
              </div>
              <div className="flex items-center justify-between">
                <button
                  type="submit"
                  className="w-1/3 py-2 px-4 bg-lime-700 text-white font-medium text-sm rounded-md shadow-sm hover:bg-lime-900 focus:outline-none focus:ring-2 focus:ring-lime-900 focus:ring-offset-2 transition duration-300"
                >
                  Register
                </button>

                <button
                  onClick={modal}
                  className="w-1/3 py-2 px-4 bg-red-700 text-white font-medium text-sm rounded-md shadow-sm hover:bg-red-900 focus:outline-none focus:ring-2 focus:ring-red-900 focus:ring-offset-2 transition duration-300"
                >
                  Cancel
                </button>
              </div>
            </form>
          </div>
        )}
      </div>
    </div>
  );
};

export default Register;
