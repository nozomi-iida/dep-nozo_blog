import { Box } from "@chakra-ui/react";
import { Footer } from "components/Footer";
import { FC, ReactNode } from "react";
import { Header } from "../Header";

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
  return (
    <Box mx="auto" maxW={970} py={14}>
      {children}
    </Box>
  );
};

// Layout.Header = HeaderLayout;
Layout.Content = ContentLayout;
