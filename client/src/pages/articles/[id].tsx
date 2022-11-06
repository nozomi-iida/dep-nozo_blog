import { Box } from "@chakra-ui/react";
import { Layout } from "components/Layout";
import { strapiClient } from "libs/strapi/api/axios";
import { Article } from "libs/strapi/models/article";
import { StrapiGetResponse, StrapiListResponse } from "libs/strapi/types";
import { GetStaticPaths, GetStaticProps, InferGetStaticPropsType } from "next";
import { NextPageWithLayout } from "pages/_app";
import { ReactElement } from "react";
import qs from "qs";

export const getStaticPaths: GetStaticPaths = async () => {
  const res = await strapiClient.get<StrapiListResponse>("articles");
  const paths = res.data.data.map((el) => ({
    params: {
      id: el.id.toString(),
    },
  }));

  return { paths, fallback: false };
};

export const getStaticProps: GetStaticProps<
  StrapiGetResponse<Article>
> = async (context) => {
  const query = qs.stringify({ populate: "*" }, { encodeValuesOnly: true });
  const res = await strapiClient.get(`articles/${context.params?.id}?${query}`);

  return {
    props: res.data,
  };
};

const Article: NextPageWithLayout<
  InferGetStaticPropsType<typeof getStaticProps>
> = ({ data }) => {
  return <Box>Hello</Box>;
};

Article.getLayout = (page: ReactElement) => {
  return (
    <Layout>
      <Layout.Content>{page}</Layout.Content>
    </Layout>
  );
};

export default Article;
