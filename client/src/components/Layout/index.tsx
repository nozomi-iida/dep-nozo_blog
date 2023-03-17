import {
  Box,
  Flex,
  Hide,
  useBoolean,
  useColorModeValue,
} from "@chakra-ui/react";
import { Footer } from "components/Footer";
import { FC, ReactNode } from "react";
import { Header } from "../Header";
import { AiFillCaretUp } from "react-icons/ai";
import { useEffect } from "react";
import { useThemeColor } from "libs/chakra/theme";

type LayoutProps = {
  children: ReactNode;
};

type LayoutSubComponent = {
  Content: FC<LayoutProps>;
  Sidebar: FC<LayoutProps>;
};

export const Layout: FC<LayoutProps> & LayoutSubComponent = ({ children }) => {
  const color = useColorModeValue("#f7f8f8", "#0b0b0c");
  return (
    <Flex flexDir="column" backgroundColor={color} pt="100px" minH="100vh">
      <Header />
      <Flex flex={1} mx="auto" maxW={970} w="full" py={14}>
        {children}
      </Flex>
      <Box mx="auto" maxW={970}>
        <Footer />
      </Box>
    </Flex>
  );
};

export const SidebarLayout: FC<LayoutProps> = ({ children }) => {
  return (
    <Hide below="md">
      <Box px={4} as="aside" minW={300} boxSizing="content-box">
        {children}
      </Box>
    </Hide>
  );
};

const ContentLayout: FC<LayoutProps> = ({ children }) => {
  const [showTopIcon, setShowTopIcon] = useBoolean(true);
  const { bgColor } = useThemeColor();
  const onClickTop = () => {
    window.scrollTo({
      top: 0,
      behavior: "smooth",
    });
  };

  useEffect(() => {
    const changeShow = () => {
      if (window.pageYOffset === 0) {
        setShowTopIcon.off();
      } else {
        setShowTopIcon.on();
      }
    };

    window.addEventListener("scroll", changeShow);

    return () => window.removeEventListener("scroll", changeShow);
  }, [setShowTopIcon]);

  return (
    <Box px={4} as="main" w="full">
      {children}
      <Box
        as="button"
        w={10}
        h={10}
        lineHeight={0}
        position="fixed"
        right="30px"
        bottom="10px"
        boxShadow="xl"
        borderRadius="50%"
        textAlign="center"
        transition="all .2s"
        opacity={showTopIcon ? 1 : 0}
        backgroundColor={bgColor}
        visibility={showTopIcon ? "visible" : "hidden"}
        _hover={{ backgroundColor: "activeColor", color: "white" }}
        onClick={onClickTop}
        zIndex="docked"
      >
        <AiFillCaretUp style={{ display: "inline" }} />
      </Box>
    </Box>
  );
};

Layout.Content = ContentLayout;
Layout.Sidebar = SidebarLayout;
