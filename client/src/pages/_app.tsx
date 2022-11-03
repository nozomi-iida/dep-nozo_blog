import "../styles/globals.css";
import type { AppProps } from "next/app";
import { Box, ChakraProvider } from "@chakra-ui/react";

export default function App({ Component, pageProps }: AppProps) {
  return (
    <ChakraProvider>
      <Box minH="100vh" backgroundColor="blackAlpha.50">
        <Component {...pageProps} />
      </Box>
    </ChakraProvider>
  );
}
