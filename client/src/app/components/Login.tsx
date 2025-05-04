import React, { useState } from 'react';
import { login } from '../utils/auth';
import Image from "next/image";
import { useRouter } from 'next/navigation'
import { useAppDispatch } from '../hooks/hook';
import { checkAuthStatus } from '../core/authSlice';

interface LoginProps {
    modal: () => void; // Function to close the modal
}


const Login: React.FC<LoginProps> = ({ modal }) => {
    const [isLoading, setIsLoading] = useState(false);
    const router = useRouter()
    const dispatch = useAppDispatch();

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
      e.preventDefault();
      setIsLoading(true);
  
      const formData = new FormData(e.currentTarget);
      const result = await login({}, formData);
      
      if (result?.success) {
        dispatch(checkAuthStatus())
        router.push('/dashboard')
      } else {
          console.log("Error Occurred here!", result.errors)
      }
      setIsLoading(false);
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
              The Path to Financial Abundance
            </h2>
            <button
              className="absolute top-2 right-5 text-red-600 text-xl font-bold hover:text-red-800 transition duration-200 focus:outline-none"
              onClick={modal}
              aria-label="Close"
            >
              &times;
            </button>
          </div>

                <form onSubmit={handleSubmit}>
                    <div className="mb-4">
                        <label className="block text-gray-700 mb-2" htmlFor="email">
                            Email
                        </label>
                        <input
                            type="email"
                            id="email"
                            name="email"
                            className="w-full p-2 border rounded"
                            autoComplete='false'
                            required
                        />
                    </div>

                    <div className="mb-6">
                        <label className="block text-gray-700 mb-2" htmlFor="password">
                            Password
                        </label>
                        <input
                            type="password"
                            id="password"
                            name="password"
                            className="w-full p-2 border rounded"
                            required
                        />
                    </div>

                    <div className="flex items-center justify-between">
                        <button
                            type="submit"
                            disabled={isLoading}
                             className="w-1/3 py-2 px-4 bg-lime-700 text-white font-medium text-sm rounded-md shadow-sm hover:bg-lime-900 focus:outline-none focus:ring-2 focus:ring-lime-900 focus:ring-offset-2 transition duration-300"
                        >
                            {isLoading ? 'Logging in...' : 'Login'}
                        </button>
                        
                    </div>
                </form>
            </div>
        </div>
        </div>
    );
};

export default Login;