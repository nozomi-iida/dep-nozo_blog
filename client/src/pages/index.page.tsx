import { ReactElement } from "react";
import { ArticleCard } from "components/ArticleCard";
import { Layout } from "components/Layout";
import { NextPageWithLayout } from "./_app.page";
import { Box, SimpleGrid } from "@chakra-ui/react";
import { GetStaticProps, InferGetStaticPropsType } from "next";
import { NextHead } from "components/NextHead";
import { pagesPath } from "libs/pathpida/$path";
import { Article, ListArticleResponse } from "libs/api/models/article";
import { restCli } from "libs/axios";

export const getStaticProps: GetStaticProps<{
  articles: Article[];
}> = async () => {
  const res = await restCli.get<ListArticleResponse>("/articles");

  return {
    props: {
      articles: res.data.articles,
    },
  };
};

const Home: NextPageWithLayout<
  InferGetStaticPropsType<typeof getStaticProps>
> = ({ articles }) => {
  return (
    <Box>
      <NextHead title="Nozo Blog" url={pagesPath.$url().pathname} />
      <SimpleGrid columns={{ sm: 1, md: 2, lg: 3 }} spacing={10}>
        {articles.map((article) => (
          <ArticleCard key={article.articleId} article={article} />
        ))}
      </SimpleGrid>
    </Box>
  );
};

Home.getLayout = (page: ReactElement) => {
  return (
    <Layout>
      <Layout.Content>{page}</Layout.Content>
    </Layout>
  );
};

export default Home;
