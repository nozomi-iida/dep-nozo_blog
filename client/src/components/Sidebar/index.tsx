import { Box, Text, VStack } from "@chakra-ui/react";
import { ArticleWidget } from "components/ArticleWidget";
import { strapiClient } from "libs/strapi/api/axios";
import { Article } from "libs/strapi/models/article";
import { StrapiListResponse } from "libs/strapi/types";
import qs from "qs";
import { ArticleMedia } from "components/ ArticleMedia";
import { useThemeColor } from "libs/chakra/theme";
import useSWR from "swr";

export const Sidebar = () => {
  const { bgColor } = useThemeColor();
  const query = qs.stringify(
    {
      populate: ["thumbnail", "topic"],
      pagination: { pageSize: 5 },
      sort: ["publishedAt:desc"],
    },
    { encodeValuesOnly: true }
  );
  const popularQuery = qs.stringify(
    {
      populate: ["thumbnail", "topic"],
      pagination: { pageSize: 5 },
      sort: ["likeCount:desc"],
    },
    { encodeValuesOnly: true }
  );

  const latestArticlesFetcher = () =>
    strapiClient
      .get<StrapiListResponse<Article>>(`articles?${query}`)
      .then((res) => {
        return res.data;
      });

  const popularArticlesFetcher = () =>
    strapiClient
      .get<StrapiListResponse<Article>>(`articles?${popularQuery}`)
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
            {latestArticlesData.data.map((el) => (
              <ArticleWidget key={el.id} id={el.id} article={el.attributes} />
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
          <VStack gap={4} align="normal">
            {popularArticlesData.data.map((el, idx) => (
              <ArticleMedia
                key={el.id}
                id={el.id}
                article={el.attributes}
                popularNum={idx + 1}
              />
            ))}
          </VStack>
        )}
      </VStack>
    </VStack>
  );
};
