import { HiHeart } from "react-icons/hi";
import { Box, Flex, Heading, Text, VStack } from "@chakra-ui/react";
import Image from "next/image";
import Link from "next/link";
import BlogTest from "../../../public/blog_test.jpeg";
import { Article } from "libs/strapi/models/article";
import { FC } from "react";
import dayjs from "dayjs";

type ArticleCardProps = { article: Article };

export const ArticleCard: FC<ArticleCardProps> = ({ article }) => {
  return (
    <Box backgroundColor="white" as="article">
      <Image alt="" src={BlogTest} width={300} height={200} />
      <VStack gap={4} p={7} align="left">
        <Text fontSize="sm" color="subInfoText" fontWeight="bold">
          {dayjs(article.publishedAt).format("YYYY-MM-DD")}
        </Text>
        <Link href="">
          <Heading size="lg" _hover={{ textDecoration: "underline" }}>
            {article.title}
          </Heading>
        </Link>
        <Box>
          <Text
            fontSize="md"
            overflow="hidden"
            style={{
              WebkitLineClamp: 6,
              WebkitBoxOrient: "vertical",
              display: "-webkit-box",
            }}
          >
            {article.content}
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
