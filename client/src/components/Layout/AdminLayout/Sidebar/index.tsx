import {
  Box,
  Heading,
  HStack,
  Icon,
  Link,
  Text,
  VStack,
} from "@chakra-ui/react";
import { pagesPath } from "libs/pathpida/$path";
import NextLink from "next/link";
import { RiArticleLine } from "react-icons/ri";

export const AdminSidebar = () => {
  const menus = [
    {
      name: "Articles",
      icon: RiArticleLine,
      path: pagesPath.admin.managements.articles.$url().pathname,
    },
  ];
  return (
    <Box py={14} px={10} as="aside" minW={330} bgColor="white">
      <NextLink href={pagesPath.admin.managements.articles.$url()}>
        <Heading transition="line-height .2s ease" whiteSpace="nowrap">
          Nozo blog
        </Heading>
      </NextLink>
      <VStack spacing={4} mt={8}>
        {menus.map((menu) => (
          <Link
            as={NextLink}
            passHref
            href={menu.path}
            w="full"
            _hover={{ textDecoration: "none", backgroundColor: "gray.100" }}
            fontSize="xl"
            fontWeight="bold"
            p={2}
            borderRadius="sm"
            key={menu.name}
          >
            <HStack justifyContent="flex-start">
              <Icon as={menu.icon} />
              <Text>{menu.name}</Text>
            </HStack>
          </Link>
        ))}
      </VStack>
    </Box>
  );
};
