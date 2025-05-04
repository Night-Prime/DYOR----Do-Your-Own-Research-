"use client"

import { useState } from "react"
import Register from "./Register";
import Login from "./Login";

const Navbar = () => {
  const [activeModal, setActiveModal] = useState<string | null>(null); // 'login' or 'register' or null

  const openRegModal = () => {
    setActiveModal('register');
  }

  const openLogModal = () => {
    setActiveModal('login');
  }

  const closeModal = () => {
    setActiveModal(null);
  }

  return (
    <>
      {activeModal === 'register' && (
        <div
          className={`w-full h-fit transition-opacity duration-700 ${
            activeModal ? "opacity-100" : "opacity-0"
          }`}
        >
          <Register modal={closeModal} />
        </div>
      )}
      
      {activeModal === 'login' && (
        <div
          className={`w-full h-fit transition-opacity duration-700 ${
            activeModal ? "opacity-100" : "opacity-0"
          }`}
        >
          <Login modal={closeModal} />
        </div>
      )}
       
      {!activeModal && (
        <div className="fixed top-0 w-[95%] h-auto flex flex-row justify-between items-center z-50">
          <div className="mt-10">
            <h1 className="text-2xl font-bold text-lime-800">DYOR</h1>
          </div>
          <div>
            <ul className="flex flex-row gap-5">
              <li className="text-lg font-semibold">
                <button onClick={openLogModal} className="bg-slate-50 text-lime-900 px-4 py-2 rounded-3xl hover:bg-slate-200 transition duration-300 cursor-pointer">
                  Login
                </button>
              </li>
              <li className="text-lg font-semibold">
                <button
                  onClick={openRegModal}
                  className="bg-lime-700 text-slate-100 px-4 py-2 rounded-3xl hover:bg-lime-900 transition duration-300 cursor-pointer"
                >
                  Signup
                </button>
              </li>
            </ul>
          </div>
        </div>
      )}
    </>
  )
}

export default Navbar
