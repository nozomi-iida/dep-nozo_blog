import { ReactElement } from "react";
import { ArticleCard } from "components/ArticleCard";
import { Layout } from "components/Layout";
import { NextPageWithLayout } from "./_app.page";
import { Box, SimpleGrid } from "@chakra-ui/react";
import { GetStaticProps, InferGetStaticPropsType } from "next";
import { StrapiListResponse } from "libs/strapi/types";
import { Article } from "libs/strapi/models/article";
import { strapiClient } from "libs/strapi/api/axios";
import qs from "qs";
import { NextHead } from "components/NextHead";
import { pagesPath } from "libs/pathpida/$path";

export const getStaticProps: GetStaticProps<{
  articles: StrapiListResponse<Article>;
}> = async () => {
  const query = qs.stringify(
    { populate: ["thumbnail"] },
    { encodeValuesOnly: true }
  );
  const res = await strapiClient.get(`articles?${query}`);

  return {
    props: {
      articles: res.data,
    },
  };
};

const Home: NextPageWithLayout<
  InferGetStaticPropsType<typeof getStaticProps>
> = ({ articles }) => {
  const { data } = articles;

  return (
    <Box>
      <NextHead title="Nozo Blog" url={pagesPath.$url().pathname} />
      <SimpleGrid columns={{ sm: 1, md: 2, lg: 3 }} spacing={10}>
        {data.map((el) => (
          <ArticleCard key={el.id} articleId={el.id} article={el.attributes} />
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
