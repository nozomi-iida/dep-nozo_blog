import { ReactElement } from "react";
import { ArticleCard } from "components/ArticleCard";
import { Layout } from "components/Layout";
import { NextPageWithLayout } from "./_app";
import { Grid } from "@chakra-ui/react";
import { GetStaticProps, InferGetStaticPropsType } from "next";
import { StrapiListResponse } from "libs/strapi/types";
import { Article } from "libs/strapi/models/article";
import { strapiClient } from "libs/strapi/api/axios";

export const getStaticProps: GetStaticProps<{
  articles: StrapiListResponse<Article>;
}> = async () => {
  const res = await strapiClient.get("articles");

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
    <Layout.Content>
      <Grid templateColumns="repeat(3, 1fr)" gap={10}>
        {data.map((el) => (
          <ArticleCard key={el.id} article={el.attributes} />
        ))}
      </Grid>
    </Layout.Content>
  );
};

// TODO: Layoutの状態が保持されているかの確認をする
Home.getLayout = (page: ReactElement) => {
  return <Layout>{page}</Layout>;
};

export default Home;
