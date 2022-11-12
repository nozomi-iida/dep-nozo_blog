import { Box, Text, VStack } from "@chakra-ui/react";
import { ArticleWidget } from "components/ArticleWidget";
import { strapiClient } from "libs/strapi/api/axios";
import { Article } from "libs/strapi/models/article";
import { StrapiListResponse } from "libs/strapi/types";
import { useEffect, useState } from "react";
import qs from "qs";
import { ArticleMedia } from "components/ ArticleMedia";

export const Sidebar = () => {
  const [articles, setArticles] = useState<StrapiListResponse<Article>>();
  const [popularArticles, setPopularArticles] =
    useState<StrapiListResponse<Article>>();

  useEffect(() => {
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

    strapiClient.get(`articles?${query}`).then((res) => {
      setArticles(res.data);
    });
    strapiClient.get(`articles?${popularQuery}`).then((res) => {
      setPopularArticles(res.data);
    });
  }, []);

  return (
    <VStack as="aside" gap={10}>
      <VStack gap={4}>
        <Box w="full" p={4} backgroundColor="white">
          <Text fontSize="lg" fontWeight="bold">
            最新の記事一覧
          </Text>
        </Box>
        {articles && (
          <VStack gap={4}>
            {articles.data.map((el) => (
              <ArticleWidget key={el.id} id={el.id} article={el.attributes} />
            ))}
          </VStack>
        )}
      </VStack>
      <VStack gap={4}>
        <Box w="full" p={4} backgroundColor="white">
          <Text fontSize="lg" fontWeight="bold">
            人気の記事
          </Text>
        </Box>
        {popularArticles && (
          <VStack gap={4} align="normal">
            {popularArticles.data.map((el, idx) => (
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
