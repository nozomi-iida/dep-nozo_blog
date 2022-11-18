import { ReactElement } from "react";
import { ArticleCard } from "components/ArticleCard";
import { Layout } from "components/Layout";
import { NextPageWithLayout } from "./_app";
import { Grid } from "@chakra-ui/react";
import { GetStaticProps, InferGetStaticPropsType } from "next";
import { StrapiListResponse } from "libs/strapi/types";
import { Article } from "libs/strapi/models/article";
import { strapiClient } from "libs/strapi/api/axios";
import qs from "qs";

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
    <Grid
      templateColumns={{
        sm: "repeat(1, 1fr)",
        md: "repeat(2, 1fr)",
        lg: "repeat(3, 1fr)",
      }}
      gap={10}
    >
      {data.map((el) => (
        <ArticleCard key={el.id} articleId={el.id} article={el.attributes} />
      ))}
    </Grid>
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
