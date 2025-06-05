"use client"

import React, { CSSProperties } from 'react';
import { DotLottieReact } from '@lottiefiles/dotlottie-react';

type LottieProps = {
    lottie: string
}

const LottieAnimation: React.FC<LottieProps> = ({ lottie }) => {
    const overlayStyle: CSSProperties = {
        width: "100%",
        height: "100%",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
    };
    return (
            <div style={overlayStyle}>
                <DotLottieReact
                    src={lottie}
                    loop
                    autoplay
                />
            </div>
    )
}

export default LottieAnimation
