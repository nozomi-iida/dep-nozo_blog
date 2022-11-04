import { Box } from "@chakra-ui/react";
import { FC, ReactNode } from "react";

type LayoutProps = {
  children: ReactNode;
};

type LayoutSubComponent = {
  Header: FC<LayoutProps>;
  Content: FC<LayoutProps>;
};

export const Layout: FC<LayoutProps> & LayoutSubComponent = ({ children }) => {
  return (
    <Box minH="100vh" backgroundColor="blackAlpha.50">
      {children}
    </Box>
  );
};

const HeaderLayout: FC<LayoutProps> = ({ children }) => {
  return (
    <Box backgroundColor="white">
      <Box maxW={970}>{children}</Box>
    </Box>
  );
};

const ContentLayout: FC<LayoutProps> = ({ children }) => {
  return <Box maxW={970}>{children}</Box>;
};

Layout.Header = HeaderLayout;
Layout.Content = ContentLayout;
