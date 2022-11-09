"use client";

import chakraTheme from "./theme";
import { ChakraProvider } from "@chakra-ui/react";
import { ReactNode } from "react";

const NextChakraProvier = ({ children }: { children: ReactNode }) => {
  return <ChakraProvider theme={chakraTheme}>{children}</ChakraProvider>;
};

export default NextChakraProvier;
