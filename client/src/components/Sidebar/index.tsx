import { Box, Text, VStack } from "@chakra-ui/react";
import { ArticleWidget } from "components/ArticleWidget";
import { strapiClient } from "libs/strapi/api/axios";
import { StrapiListResponse } from "libs/strapi/types";
import qs from "qs";
import { ArticleMedia } from "components/ ArticleMedia";
import { useThemeColor } from "libs/chakra/theme";
import useSWR from "swr";
import { restCli } from "libs/axios";
import {
  Article,
  ArticleOrderBy,
  ArticleQueryParams,
} from "libs/api/models/article";

export const Sidebar = () => {
  const { bgColor } = useThemeColor();
  const query: ArticleQueryParams = {
    orderBy: ArticleOrderBy.PUBLISHED_AT_ASC,
  };
  const popularQuery = qs.stringify(
    {
      populate: ["thumbnail", "topic"],
      pagination: { pageSize: 5 },
      sort: ["likeCount:desc"],
    },
    { encodeValuesOnly: true }
  );

  const latestArticlesFetcher = () =>
    restCli
      .get<{ articles: Article[] }>("/articles", {
        params: query,
      })
      .then((res) => {
        return res.data;
      });

  const popularArticlesFetcher = () =>
    restCli
      .get<{ articles: Article[] }>("/articles", {
        params: query,
      })
      .then((res) => {
        return res.data;
      });

  const { data: latestArticlesData } = useSWR(
    "latestArticles",
    latestArticlesFetcher
  );
  const { data: popularArticlesData } = useSWR(
    "latestArticles",
    popularArticlesFetcher
  );

  return (
    <VStack gap={10}>
      <VStack gap={4} w="full">
        <Box w="full" p={4} backgroundColor={bgColor}>
          <Text fontSize="lg" fontWeight="bold">
            最新の記事一覧
          </Text>
        </Box>
        {latestArticlesData && (
          <VStack gap={4}>
            {latestArticlesData.articles.map((article) => (
              <ArticleWidget key={article.articleId} article={article} />
            ))}
          </VStack>
        )}
      </VStack>
      <VStack gap={4} w="full">
        <Box w="full" p={4} backgroundColor={bgColor}>
          <Text fontSize="lg" fontWeight="bold">
            人気の記事
          </Text>
        </Box>
        {popularArticlesData && (
          <VStack w="full" gap={4} align="normal">
            {popularArticlesData.articles.map((article, idx) => (
              <ArticleMedia
                key={article.articleId}
                article={article}
                popularNum={idx + 1}
              />
            ))}
          </VStack>
        )}
      </VStack>
    </VStack>
  );
};
