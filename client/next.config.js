/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
  images: {
    domains: [
      "localhost",
      "127.0.0.1",
      "nozo-blog-images.s3.ap-northeast-1.amazonaws.com",
    ],
  },
  experimental: {
    appDir: true,
  },
  pageExtensions: ["page.tsx"],
};

module.exports = nextConfig;
