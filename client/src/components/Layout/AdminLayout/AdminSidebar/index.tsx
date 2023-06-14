import {
  Box,
  Heading,
  Hide,
  HStack,
  Icon,
  Link,
  Text,
  VStack,
} from "@chakra-ui/react";
import { pagesPath } from "libs/pathpida/$path";
import { RiArticleLine } from "react-icons/ri";
import NextLink from "next/link";

export const AdminSidebar = () => {
  const contens = [
    {
      title: "Articles",
      icon: RiArticleLine,
      path: pagesPath.admin.managements.articles.$url(),
    },
    {
      title: "Topics",
      icon: RiArticleLine,
      path: pagesPath.admin.managements.articles.$url(),
    },
  ];
  return (
    <Hide below="md">
      <Box px={4} as="aside" minW={300} boxSizing="content-box">
        <Link as={NextLink} href={pagesPath.$url() as any} passHref>
          <Heading
            textAlign="center"
            transition="line-height .2s ease"
            whiteSpace="nowrap"
          >
            Nozo blog
          </Heading>
        </Link>
        <VStack>
          {contens.map((content) => (
            <HStack
              as={NextLink}
              key={content.title}
              href={content.path as any}
              gap={4}
              w="full"
            >
              <Icon as={content.icon} />
              <Text as="b">{content.title}</Text>
            </HStack>
          ))}
        </VStack>
      </Box>
    </Hide>
  );
};
