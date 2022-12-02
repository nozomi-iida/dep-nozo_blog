import { Box, Heading, SimpleGrid } from "@chakra-ui/react";
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
import { NextHead } from "components/NextHead";
import { pagesPath } from "libs/pathpida/$path";

const Search: NextPageWithLayout = () => {
  const router = useRouter();
  const keyword = (router.query.keyword ?? "") as string;
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
      <NextHead title={keyword} url={pagesPath.search.$url().pathname} />
      <Heading fontSize="2xl" mb={4}>
        検索ワード: {keyword}
      </Heading>

      <SimpleGrid columns={{ sm: 1, md: 2, lg: 3 }} spacing={10}>
        {searchArticlesData?.map((article) => (
          <ArticleCard
            key={article.id}
            article={article.attributes}
            articleId={article.id}
          />
        ))}
      </SimpleGrid>
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
