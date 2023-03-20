import {
  extendTheme,
  StyleFunctionProps,
  useColorModeValue,
} from "@chakra-ui/react";
import { colors } from "./colors";

export const useThemeColor = () => {
  const bgColor = useColorModeValue("white", "#18191b");

  return { bgColor };
};

const chakraTheme = extendTheme({
  colors,
  components: {
    Text: {
      baseStyle: (props: StyleFunctionProps) => ({
        color: props.colorMode === "light" ? "#2f3235" : "#eaedf1",
      }),
    },
    Heading: {
      baseStyle: (props: StyleFunctionProps) => ({
        color: props.colorMode === "light" ? "#2f3235" : "#eaedf1",
      }),
    },
    Alert: {
      variants: {},
    },
  },
  breakpoints: {
    sm: "320px",
    md: "768px",
    lg: "960px",
    xl: "1200px",
    "2xl": "1536px",
  },
  config: {
    initialColorMode: "light",
    useSystemColorMode: false,
  },
});

export default chakraTheme;
