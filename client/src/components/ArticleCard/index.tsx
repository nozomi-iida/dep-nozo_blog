import { HiHeart } from "react-icons/hi";
import { Box, Flex, Heading, HStack, Text, VStack } from "@chakra-ui/react";
import Image from "next/image";
import Link from "next/link";
import BlogTest from "../../../public/blog_test.jpeg";

export const ArticleCard = () => {
  return (
    <Box backgroundColor="white" as="article">
      <Image alt="" src={BlogTest} width={300} height={200} />
      <VStack gap={4} p={7} align="left">
        <Text fontSize="sm" color="subInfoText" fontWeight="bold">
          2019, 10 31
        </Text>
        <Link href="">
          <Heading size="lg" _hover={{ textDecoration: "underline" }}>
            Gallery format - Sticky post
          </Heading>
        </Link>
        <Box>
          <Text fontSize="md">
            Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean
            commodo ligula eget dolor. Aenean massa. Cum sociis natoque
            penatibus et magnis dis parturient montes, nascetur ridiculus mus….{" "}
          </Text>
          <Link href="">
            <Text fontSize="md" as="u" _hover={{ color: "activeColor" }}>
              Read more
            </Text>
          </Link>
        </Box>
        {/* FIXME: ハートが揃ってない */}
        <Flex justify="right" color="subInfoText">
          <HiHeart />
          <Text fontSize="sm">100</Text>
        </Flex>
      </VStack>
    </Box>
  );
};
