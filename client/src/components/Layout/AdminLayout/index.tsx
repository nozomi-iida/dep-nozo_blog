import { Box, Flex, useColorModeValue } from "@chakra-ui/react";
import { Footer } from "components/Footer";
import { FC, ReactNode } from "react";
import { AdminSidebar } from "./AdminSidebar";

type LayoutProps = {
  children: ReactNode;
};

type LayoutSubComponent = {
  Content: FC<LayoutProps>;
  Sidebar: FC;
};

export const AdminLayout: FC<LayoutProps> & LayoutSubComponent = ({
  children,
}) => {
  const color = useColorModeValue("#f7f8f8", "#0b0b0c");
  return (
    <Flex flexDir="column" backgroundColor={color} minH="100vh">
      <Flex flex={1}>{children}</Flex>
    </Flex>
  );
};

const ContentLayout: FC<LayoutProps> = ({ children }) => {
  return (
    <Flex flexDir="column" mx="auto" maxW={970} w="full">
      <Box flex={1} mx="auto" maxW={970} w="full" py={4}>
        {children}
      </Box>
      <Box mx="auto">
        <Footer />
      </Box>
    </Flex>
  );
};

AdminLayout.Sidebar = AdminSidebar;
AdminLayout.Content = ContentLayout;
