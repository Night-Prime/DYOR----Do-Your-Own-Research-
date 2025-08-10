"use client"

import { useState } from "react"
import Register from "./Register";
import Login from "./Login";

const Navbar = () => {
  const [activeModal, setActiveModal] = useState<string | null>(null); // 'login' or 'register' or null

  const openRegModal = () => {
    setActiveModal('register');
  }

  // const openLogModal = () => {
  //   setActiveModal('login');
  // }

  const closeModal = () => {
    setActiveModal(null);
  }

  return (
    <>
      {activeModal === 'register' && (
        <div
          className={`w-full h-fit transition-opacity duration-700 ${activeModal ? "opacity-100" : "opacity-0"
            }`}
        >
          <Register modal={closeModal} />
        </div>
      )}

      {activeModal === 'login' && (
        <div
          className={`w-full h-fit transition-opacity duration-700 ${activeModal ? "opacity-100" : "opacity-0"
            }`}
        >
          <Login modal={closeModal} />
        </div>
      )}

      {!activeModal && (
        <div className="fixed top-8 left-1/2 transform -translate-x-1/2 w-[50%] h-auto z-50">
          <div
            className="relative px-8 py-4 flex flex-row justify-between items-center gap-8 rounded-full backdrop-blur-xl bg-black/20 border border-white/10 shadow-2xl"
            style={{
              background: 'linear-gradient(135deg, rgba(255,255,255,0.1), rgba(255,255,255,0.05))',
              backdropFilter: 'blur(20px)',
              WebkitBackdropFilter: 'blur(20px)',
              boxShadow: `
              0 8px 32px rgba(0, 0, 0, 0.3),
              inset 0 1px 0 rgba(255, 255, 255, 0.2),
              inset 0 -1px 0 rgba(255, 255, 255, 0.1)
            `
            }}
          >
            {/* Subtle inner glow */}
            <div
              className="absolute inset-0 rounded-full opacity-50"
              style={{
                background: 'linear-gradient(135deg, rgba(255,255,255,0.15) 0%, transparent 50%, rgba(255,255,255,0.05) 100%)'
              }}
            />

            {/* Content */}
            <div className="relative z-10">
              <h1 className="text-xl font-bold text-white drop-shadow-lg">Insights</h1>
            </div>

            <div className="relative z-10">
              <ul className="flex flex-row gap-4">
                {/* <li>
                  <button
                    onClick={openLogModal}
                    className="text-white/90 hover:text-white px-6 py-2 rounded-full hover:bg-white/10 transition-all duration-300 font-medium backdrop-blur-sm border border-white/5 hover:border-white/20 shadow-lg"
                  >
                    Login
                  </button>
                </li> */}
                <li>
                  <button
                    onClick={openRegModal}
                    className="text-white px-6 py-2 rounded-full bg-white/15 hover:bg-white/25 transition-all duration-300 font-medium backdrop-blur-sm border border-white/20 hover:border-white/30 shadow-lg"
                  >
                    Join Beta
                  </button>
                </li>
              </ul>
            </div>
          </div>
        </div>
      )}
    </>
  )
}

export default Navbar
