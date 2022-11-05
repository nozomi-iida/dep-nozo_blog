import { extendTheme } from "@chakra-ui/react";

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
      baseStyle: {
        color: "#574F53",
      },
    },
    Heading: {
      baseStyle: {
        color: "#2f3235",
      },
    },
  },
});

export default chakraTheme;
