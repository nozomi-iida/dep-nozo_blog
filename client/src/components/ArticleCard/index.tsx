import { Box, Heading, Text, VStack } from "@chakra-ui/react";
import Link from "next/link";
import { FC } from "react";
import dayjs from "dayjs";
import { pagesPath } from "libs/pathpida/$path";
import { Image } from "components/Image";
import { useThemeColor } from "libs/chakra/theme";
import { markdown2content } from "utils/helpers";
import { Article } from "libs/api/models/article";

type ArticleCardProps = { article: Article };

export const ArticleCard: FC<ArticleCardProps> = ({ article }) => {
  const { bgColor } = useThemeColor();
  const normalContent = markdown2content(article.content);

  return (
    <Box backgroundColor={bgColor} as="article">
      <Image alt={article.title} src={undefined} width="full" height={200} />
      <VStack gap={4} p={7} align="left">
        <Text fontSize="sm" color="subInfoText" fontWeight="bold">
          {dayjs(article.publishedAt).format("YYYY-MM-DD")}
        </Text>
        <Link href={pagesPath.articles._id(article.articleId).$url()}>
          <Heading
            fontSize="2xl"
            size="lg"
            _hover={{ textDecoration: "underline" }}
            overflow="hidden"
            sx={{
              WebkitLineClamp: 2,
              WebkitBoxOrient: "vertical",
              display: "-webkit-box",
            }}
          >
            {article.title}
          </Heading>
        </Link>
        <Text
          overflow="hidden"
          sx={{
            WebkitLineClamp: 6,
            WebkitBoxOrient: "vertical",
            display: "-webkit-box",
          }}
        >
          {normalContent}
        </Text>
        <Link href={pagesPath.articles._id(article.articleId).$url()}>
          <Text fontSize="md" as="u" _hover={{ color: "activeColor" }}>
            Read more
          </Text>
        </Link>
      </VStack>
    </Box>
  );
};
