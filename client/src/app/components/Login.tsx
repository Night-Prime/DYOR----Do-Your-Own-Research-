import React, { useState } from 'react';
import { login } from '../utils/auth';
import Image from "next/image";
import { useRouter } from 'next/navigation'
import { useAppDispatch } from '../hooks/hook';
import { checkAuthStatus, loginSuccess } from '../core/authSlice';
import { DyorAlert } from '../shared/Alert';
import { EyeIcon, EyeSlashIcon } from '../shared/icons';
import Preloader from '../shared/Preloader';

interface LoginProps {
  modal: () => void;
}


const Login: React.FC<LoginProps> = ({ modal }) => {
  const [showAlert, setShowAlert] = useState<{
    type: 'success' | 'error';
    message: string;
  } | null>(null);

  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [showPassword, setShowPassword] = useState<boolean>(false);

  const router = useRouter()
  const dispatch = useAppDispatch();

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setIsLoading(true);

    const formData = new FormData(e.currentTarget);
    const result = await login({}, formData);

    if (result.success) {
      dispatch(checkAuthStatus());
      dispatch(loginSuccess(result.data))
      setShowAlert({
        type: 'success',
        message: 'Login successful!'
      });
      setTimeout(() => router.push('/dashboard'), 2000);
    } else {
      setShowAlert({
        type: 'error',
        message: 'Login failed!'
      });
      setIsLoading(false)
    }
  };

  return (
    <>
    {isLoading && <Preloader />}
      {showAlert && (
        <DyorAlert
          type={showAlert.type}
          message={showAlert.message}
          open={true}
          autoClose={true}
          onClose={() => setShowAlert(null)}
        />
      )}
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
        <div
          className="relative bg-white w-full max-w-sm max-h-2/3 rounded-3xl p-12 shadow-lg z-10 overflow-y-auto"
          style={{ scrollbarWidth: "none", msOverflowStyle: "none" }}
        >
          <style jsx>{`
            div::-webkit-scrollbar {
              display: none;
            }
          `}</style>
          <div className="w-full flex flex-row justify-between items-center mb-4">
            <h2 className="text-2xl font-bold text-gray-800">
              Login into DYOR
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
                type="email"
                id="email"
                name="email"
                placeholder="Email"
                className="rounded-3xl border-gray-300 shadow-sm px-4 py-2 focus:border-lime-500 focus:ring-lime-500 sm:text-sm"
                required
              />
            </div>
            <div className="flex flex-col space-y-2 relative">
              <input
                type={showPassword ? 'text' : 'password'}
                id="password"
                name="password"
                placeholder="*********"
                className="rounded-3xl border-gray-300 shadow-sm px-4 py-2 focus:border-lime-500 focus:ring-lime-500 sm:text-sm pr-10"
                required
              />
              <button
                type="button"
                className="absolute right-3 top-[40%] transform -translate-y-1/2 text-gray-400 hover:text-gray-600 focus:outline-none"
                onClick={() => setShowPassword(!showPassword)}
                aria-label={showPassword ? 'Hide password' : 'Show password'}
              >
                {showPassword ? (
                  <EyeSlashIcon />
                ) : (
                  <EyeIcon />
                )}
              </button>
            </div>

            <div className="flex items-center justify-between">
              <button
                type="submit"
                disabled={isLoading}
                className="w-1/3 py-2 px-4 bg-lime-700 text-white font-medium text-sm rounded-3xl shadow-sm hover:bg-lime-900 focus:outline-none focus:ring-2 focus:ring-lime-900 focus:ring-offset-2 transition duration-300"
              >
                {isLoading ? 'Logging in...' : 'Login'}
              </button>

            </div>
          </form>
        </div>
      </div>
    </>
  );
};

export default Login;