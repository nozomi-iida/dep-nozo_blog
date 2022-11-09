import { Box, Heading, Text, VStack } from "@chakra-ui/react";
import Image from "next/image";
import Link from "next/link";
import { Article } from "libs/strapi/models/article";
import { FC } from "react";
import dayjs from "dayjs";
import { pagesPath } from "libs/pathpida/$path";

type ArticleCardProps = { articleId: number; article: Article };

// TODO: 説明部分にマークダウンが表示されるから表示しない方が良いかも
export const ArticleCard: FC<ArticleCardProps> = ({ articleId, article }) => {
  return (
    <Box backgroundColor="white" as="article">
      <Image
        alt=""
        src={`${process.env.NEXT_PUBLIC_STRAPI_URI}${
          article.thumbnail?.data?.attributes.url ?? ""
        }`}
        width={300}
        height={200}
      />
      <VStack gap={4} p={7} align="left">
        <Text fontSize="sm" color="subInfoText" fontWeight="bold">
          {dayjs(article.publishedAt).format("YYYY-MM-DD")}
        </Text>
        <Link href={pagesPath.articles._id(articleId).$url()}>
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
      </VStack>
    </Box>
  );
};
