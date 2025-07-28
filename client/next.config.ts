import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  // reactStrictMode: false,
  /* config options here */
  devIndicators: false,
  images: {
    remotePatterns: [
      {
        protocol: 'https',
        hostname: '**',
        pathname: '/**',
      },
    ],
  }
};

export default nextConfig;
