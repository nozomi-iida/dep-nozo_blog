import {
  Box,
  Heading,
  Text,
  useColorModeValue,
  VStack,
} from "@chakra-ui/react";
import Link from "next/link";
import { Article } from "libs/strapi/models/article";
import { FC } from "react";
import dayjs from "dayjs";
import { pagesPath } from "libs/pathpida/$path";
import { Image } from "components/Image";
import { useThemeColor } from "libs/chakra/theme";

type ArticleCardProps = { articleId: number; article: Article };

// TODO: カードのデザイン変えたほうが良いかも
export const ArticleCard: FC<ArticleCardProps> = ({ articleId, article }) => {
  const { bgColor } = useThemeColor();

  return (
    <Box backgroundColor={bgColor} as="article">
      <Image
        alt={article.title}
        src={`${process.env.NEXT_PUBLIC_STRAPI_URI}${
          article.thumbnail?.data?.attributes.url ?? ""
        }`}
        width="full"
        height={200}
      />
      <VStack gap={4} p={7} align="left">
        <Text fontSize="sm" color="subInfoText" fontWeight="bold">
          {dayjs(article.publishedAt).format("YYYY-MM-DD")}
        </Text>
        <Link href={pagesPath.articles._id(articleId).$url()}>
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
        <Link href={pagesPath.articles._id(articleId).$url()}>
          <Text fontSize="md" as="u" _hover={{ color: "activeColor" }}>
            Read more
          </Text>
        </Link>
      </VStack>
    </Box>
  );
};
