import { Box, Grid, Heading } from "@chakra-ui/react";
import { ArticleCard } from "components/ArticleCard";
import { Layout } from "components/Layout";
import { strapiClient } from "libs/strapi/api/axios";
import { Article } from "libs/strapi/models/article";
import { StrapiListResponse } from "libs/strapi/types";
import { useRouter } from "next/router";
import { NextPageWithLayout } from "pages/_app";
import { useEffect, useState } from "react";
import { ReactElement } from "react-markdown/lib/react-markdown";
import qs from "qs";

const Search: NextPageWithLayout = () => {
  const router = useRouter();
  const keyword = router.query.keyword;
  const [articles, setArticles] = useState<StrapiListResponse<Article>["data"]>(
    []
  );

  useEffect(() => {
    const query = qs.stringify(
      {
        populate: "*",
        filters: { title: { $contains: keyword } },
      },
      { encodeValuesOnly: true }
    );
    strapiClient
      .get<StrapiListResponse<Article>>(`articles?${query}`)
      .then((res) => {
        setArticles(res.data.data);
      });
  }, [keyword]);

  return (
    <Box>
      <Heading fontSize="2xl" mb={4}>
        検索ワード: {keyword}
      </Heading>
      <Grid templateColumns="repeat(3, 1fr)" gap={10}>
        {articles?.map((article) => (
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
