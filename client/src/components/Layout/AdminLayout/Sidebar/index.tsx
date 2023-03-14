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
import { RiArticleFill } from "react-icons/ri";
import { MdTopic } from "react-icons/md";
import { useRouter } from "next/router";

export const AdminSidebar = () => {
  const menus = [
    {
      name: "Articles",
      icon: RiArticleFill,
      path: pagesPath.admin.managements.articles.$url().pathname,
    },
    {
      name: "Topics",
      icon: MdTopic,
      path: pagesPath.admin.managements.articles.$url().pathname,
    },
  ];
  const router = useRouter();

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
            _hover={{
              textDecoration: "none",
              bgColor: "gray.100",
              color: "black",
            }}
            fontSize="xl"
            fontWeight="bold"
            p={2}
            borderRadius="sm"
            key={menu.name}
            color={menu.path === router.pathname ? "black" : "gray.600"}
            bgColor={menu.path === router.pathname ? "gray.100" : "transparent"}
          >
            <HStack justifyContent="flex-start">
              <Icon as={menu.icon} />
              <Box as="p">{menu.name}</Box>
            </HStack>
          </Link>
        ))}
      </VStack>
    </Box>
  );
};
