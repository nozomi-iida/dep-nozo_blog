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

  useEffect(() => {
    const query = qs.stringify(
      { populate: ["thumbnail", "topic"], pagination: { pageSize: 5 } },
      { encodeValuesOnly: true }
    );

    strapiClient.get(`articles?${query}`).then((res) => {
      setArticles(res.data);
    });
  }, []);

  return (
    <VStack as="aside" gap={10}>
      <VStack gap={4}>
        <Box w="full" p={4} backgroundColor="white">
          <Text fontSize="lg" fontWeight="bold">
            記事一覧
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
        {articles && (
          <VStack gap={4} align="normal">
            {articles.data.map((el, idx) => (
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
