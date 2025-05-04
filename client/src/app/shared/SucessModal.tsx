import { useState, useEffect } from 'react';
import Login from '../components/Login';

// plan to make it reuseable
export default function SuccessModal() {
  const [animationComplete, setAnimationComplete] = useState(false);
  const [login, setLogin] = useState(false);

  const openLogin = () => {
    setLogin(true)
  }
  const closeLogin = () => {
    setLogin(false)
  }
  
  useEffect(() => {
    const timer = setTimeout(() => {
      setAnimationComplete(true);
    }, 1000);
    
    return () => clearTimeout(timer);
  }, []);
  
  return (
    <>
    {login ? (
      <Login modal={closeLogin} />
    ): (
      <div className="w-1/4 h-full flex flex-col items-center justify-center text-center z-10">
      <div className="max-w-md w-full bg-white rounded-lg shadow-lg p-8 space-y-6 transition-all duration-500">
        <div className={`w-24 h-24 mx-auto relative ${animationComplete ? 'scale-100' : 'scale-0'} transition-transform duration-500`}>
          <div className="w-24 h-24 rounded-full bg-green-100 flex items-center justify-center">
            <svg 
              xmlns="http://www.w3.org/2000/svg" 
              className="h-12 w-12 text-green-500" 
              fill="none" 
              viewBox="0 0 24 24" 
              stroke="currentColor" 
              strokeWidth={3}
            >
              <path 
                strokeLinecap="round" 
                strokeLinejoin="round" 
                d="M5 13l4 4L19 7" 
                className={`${animationComplete ? 'opacity-100' : 'opacity-0'} transition-opacity duration-300 delay-500`}
              />
            </svg>
          </div>
        </div>
        
        <h1 className="text-3xl font-bold text-gray-800 mt-6">
          Registration Successful!
        </h1>
        
        <p className="text-gray-600 mt-2">
          Your account has been created. You can now log in to access your dashboard.
        </p>
        
        <div className="pt-4">
          <button
            onClick={openLogin}
             className="w-2/3 py-2 px-4 bg-lime-700 text-white font-medium text-sm rounded-md shadow-sm hover:bg-lime-900 focus:outline-none focus:ring-2 focus:ring-lime-900 focus:ring-offset-2 transition duration-300"
          >
            Proceed to Login
          </button>
        </div>
      </div>
    </div>
    )}
    
    </>
  );
}