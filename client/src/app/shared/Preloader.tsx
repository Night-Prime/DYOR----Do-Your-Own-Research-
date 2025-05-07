"use client"

import React, { CSSProperties } from 'react';
import { DotLottieReact } from '@lottiefiles/dotlottie-react';

const Preloader = () => {
    const overlayStyle: CSSProperties = {
        position: 'fixed',
        top: 0,
        left: 0,
        width: "100%",
        height: "100%",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        zIndex: 9999,
        backgroundColor: "rgba(255, 255, 255, 0.7)",
    };

    return (
        <div style={overlayStyle}>
            <DotLottieReact
                src="/assets/Animation.lottie"
                loop
                autoplay
            />
        </div>
    );
}


export default Preloader
