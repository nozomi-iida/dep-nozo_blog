import { Box, Button, useBoolean } from "@chakra-ui/react";
import { Footer } from "components/Footer";
import { FC, ReactNode } from "react";
import { Header } from "../Header";
import { AiFillCaretUp } from "react-icons/ai";
import { useState } from "react";
import { useEffect } from "react";

type LayoutProps = {
  children: ReactNode;
};

type LayoutSubComponent = {
  // Header: FC<LayoutProps>;
  Content: FC<LayoutProps>;
};

export const Layout: FC<LayoutProps> & LayoutSubComponent = ({ children }) => {
  return (
    <Box minH="100vh" backgroundColor="blackAlpha.50">
      <Box backgroundColor="white">
        <Box mx="auto" maxW={970}>
          <Header />
        </Box>
      </Box>
      {children}
      <Box mx="auto" maxW={970}>
        <Footer />
      </Box>
    </Box>
  );
};

// const HeaderLayout: FC<LayoutProps> = ({ children }) => {
//   return (
//     <Box backgroundColor="white">
//       <Box maxW={970}>{children}</Box>
//     </Box>
//   );
// };

const ContentLayout: FC<LayoutProps> = ({ children }) => {
  const [showTopIcon, setShowTopIcon] = useBoolean(true);
  const onClickTop = () => {
    window.scrollTo({
      top: 0,
      behavior: "smooth",
    });
  };

  useEffect(() => {
    const changeShow = () => {
      if (window.pageYOffset > 200) {
        setShowTopIcon.on();
      } else {
        setShowTopIcon.off();
      }
    };

    window.addEventListener("scroll", changeShow);

    return () => window.removeEventListener("scroll", changeShow);
  }, [setShowTopIcon]);

  return (
    <Box mx="auto" maxW={970} py={14}>
      {children}
      <Box
        as="button"
        w={10}
        h={10}
        lineHeight={0}
        position="fixed"
        right="30px"
        bottom="30px"
        boxShadow="xl"
        borderRadius="50%"
        textAlign="center"
        transition="all .2s"
        opacity={showTopIcon ? 1 : 0}
        visibility={showTopIcon ? "visible" : "hidden"}
        _hover={{ backgroundColor: "activeColor", color: "white" }}
        onClick={onClickTop}
      >
        <AiFillCaretUp style={{ display: "inline" }} />
      </Box>
    </Box>
  );
};

// Layout.Header = HeaderLayout;
Layout.Content = ContentLayout;
