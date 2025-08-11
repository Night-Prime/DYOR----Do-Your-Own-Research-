"use client";

import { useState } from "react";
import Register from "./Register";
import Login from "./Login";
import { motion, useScroll, useSpring, useTransform } from "framer-motion";

const Navbar = () => {
  const { scrollY } = useScroll();

  // Shrinks width from 50% to 30% after scrolling one full viewport height
  const rawWidth = useTransform(
    scrollY,
    [0, window.innerHeight], // 0 to full screen height
    ["50%", "20%"]
  );
  
  // Apply a spring to smooth it out
  const widthScale = useSpring(rawWidth, {
    stiffness: 250,   // how strong the spring is
    damping: 15,      // how quickly it settles
    mass: 0.5         // lower = more responsive
  });

  const [activeModal, setActiveModal] = useState<null | "login" | "register">(null);

  const openRegModal = () => setActiveModal("register");
  const closeModal = () => setActiveModal(null);

  return (
    <>
      {activeModal === "register" && (
        <div className="w-full h-fit transition-opacity duration-700 opacity-100">
          <Register modal={closeModal} />
        </div>
      )}

      {activeModal === "login" && (
        <div className="w-full h-fit transition-opacity duration-700 opacity-100">
          <Login modal={closeModal} />
        </div>
      )}

      {!activeModal && (
        <motion.div
          style={{
            width: widthScale,
          }}
          className="fixed top-8 left-1/2 transform -translate-x-1/2 h-auto z-50"
        >
          <div
            className="relative px-8 py-4 flex flex-row justify-between items-center gap-8 rounded-full backdrop-blur-xl bg-black/20 border border-white/10 shadow-2xl"
            style={{
              background: "linear-gradient(135deg, rgba(255,255,255,0.1), rgba(255,255,255,0.05))",
              backdropFilter: "blur(20px)",
              WebkitBackdropFilter: "blur(20px)",
              boxShadow: `
                0 8px 32px rgba(0, 0, 0, 0.3),
                inset 0 1px 0 rgba(255, 255, 255, 0.2),
                inset 0 -1px 0 rgba(255, 255, 255, 0.1)
              `,
            }}
          >
            {/* Subtle inner glow */}
            <div
              className="absolute inset-0 rounded-full opacity-50"
              style={{
                background:
                  "linear-gradient(135deg, rgba(255,255,255,0.15) 0%, transparent 50%, rgba(255,255,255,0.05) 100%)",
              }}
            />

            {/* Logo */}
            <div className="relative z-10">
              <h1 className="text-xl font-bold text-white drop-shadow-lg">Insights</h1>
            </div>

            {/* Actions */}
            <div className="relative z-10">
              <ul className="flex flex-row gap-4">
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
        </motion.div>
      )}
    </>
  );
};

export default Navbar;
