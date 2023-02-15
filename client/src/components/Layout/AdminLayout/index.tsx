import {
  Box,
  Flex,
  Hide,
  useBoolean,
  useColorModeValue,
} from "@chakra-ui/react";
import { Footer } from "components/Footer";
import { FC, ReactNode } from "react";

type LayoutProps = {
  children: ReactNode;
};

type LayoutSubComponent = {
  Content: FC<LayoutProps>;
  Sidebar: FC<LayoutProps>;
};

export const AdminLayout: FC<LayoutProps> & LayoutSubComponent = ({
  children,
}) => {
  const color = useColorModeValue("#f7f8f8", "#0b0b0c");
  return (
    <Flex flexDir="column" backgroundColor={color} minH="100vh">
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
  return (
    <Box px={4} as="main" w="full">
      {children}
    </Box>
  );
};

AdminLayout.Sidebar = SidebarLayout;
AdminLayout.Content = ContentLayout;
