import { Box, useBoolean } from "@chakra-ui/react";
import { Footer } from "components/Footer";
import { FC, ReactNode } from "react";
import { Header } from "../Header";
import { AiFillCaretUp } from "react-icons/ai";
import { useEffect } from "react";

type LayoutProps = {
  children: ReactNode;
};

type LayoutSubComponent = {
  Content: FC<LayoutProps>;
};

export const Layout: FC<LayoutProps> & LayoutSubComponent = ({ children }) => {
  return (
    <Box backgroundColor="blackAlpha.50">
      <Header />
      {children}
      <Box mx="auto" maxW={970}>
        <Footer />
      </Box>
    </Box>
  );
};

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
      if (window.pageYOffset > 600) {
        setShowTopIcon.on();
      } else {
        setShowTopIcon.off();
      }
    };

    window.addEventListener("scroll", changeShow);

    return () => window.removeEventListener("scroll", changeShow);
  }, [setShowTopIcon]);

  return (
    <Box mx="auto" maxW={970} py={14} as="main" minH="100vh">
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

Layout.Content = ContentLayout;
