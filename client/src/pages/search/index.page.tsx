import { Box, Heading, SimpleGrid } from "@chakra-ui/react";
import { ArticleCard } from "components/ArticleCard";
import { Layout } from "components/Layout";
import { useRouter } from "next/router";
import { NextPageWithLayout } from "pages/_app.page";
import useSWR from "swr";
import { NextHead } from "components/NextHead";
import { pagesPath } from "libs/pathpida/$path";
import { ReactElement } from "react";
import { restCli } from "libs/axios";
import { ListArticleResponse } from "libs/api/models/article";
import qs from "qs";

export type OptionalQuery = {
  keyword?: string;
};

const Search: NextPageWithLayout = () => {
  const router = useRouter();
  const keyword = (router.query.keyword ?? "") as string;
  const query = qs.stringify({ keyword });
  const searchArticlesFetcher = (url: string) =>
    restCli.get<ListArticleResponse>(url).then((res) => res.data);

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
        {searchArticlesData?.articles.map((article) => (
          <ArticleCard key={article.articleId} article={article} />
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
