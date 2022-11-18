import {
  extendTheme,
  StyleFunctionProps,
  useColorModeValue,
} from "@chakra-ui/react";

export const useThemeColor = () => {
  const bgColor = useColorModeValue("white", "#18191b");

  return { bgColor };
};

const chakraTheme = extendTheme({
  colors: {
    // theme color
    activeColor: "#6ca4db",
    borderColor: "#e3e5e6",
    // text color
    subInfoText: "#989ea6",
  },
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
