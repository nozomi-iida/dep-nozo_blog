import { Box, Grid, Heading } from "@chakra-ui/react";
import { ArticleCard } from "components/ArticleCard";
import { Layout } from "components/Layout";
import { strapiClient } from "libs/strapi/api/axios";
import { Article } from "libs/strapi/models/article";
import { StrapiListResponse } from "libs/strapi/types";
import { useRouter } from "next/router";
import { NextPageWithLayout } from "pages/_app";
import { ReactElement } from "react-markdown/lib/react-markdown";
import qs from "qs";
import useSWR from "swr";

const Search: NextPageWithLayout = () => {
  const router = useRouter();
  const keyword = router.query.keyword;
  const query = qs.stringify(
    {
      populate: "*",
      filters: { title: { $contains: keyword } },
    },
    { encodeValuesOnly: true }
  );
  const searchArticlesFetcher = () =>
    strapiClient
      .get<StrapiListResponse<Article>>(`articles?${query}`)
      .then((res) => res.data.data);

  const { data: searchArticlesData } = useSWR(
    `articles?${query}`,
    searchArticlesFetcher
  );

  return (
    <Box>
      <Heading fontSize="2xl" mb={4}>
        検索ワード: {keyword}
      </Heading>
      <Grid templateColumns="repeat(3, 1fr)" gap={10}>
        {searchArticlesData?.map((article) => (
          <ArticleCard
            key={article.id}
            article={article.attributes}
            articleId={article.id}
          />
        ))}
      </Grid>
    </Box>
  );
};

Search.getLayout = (page: ReactElement) => {
  return (
    <Layout>
      <Layout.Content>{page}</Layout.Content>
    </Layout>
  );
};

export default Search;
