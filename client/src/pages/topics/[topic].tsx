import { Box, Grid, Heading } from "@chakra-ui/react";
import { Layout } from "components/Layout";
import { strapiClient } from "libs/strapi/api/axios";
import { Topic, Topic as TopicType } from "libs/strapi/models/topic";
import { StrapiGetResponse, StrapiListResponse } from "libs/strapi/types";
import { GetStaticPaths, GetStaticProps, InferGetStaticPropsType } from "next";
import { NextPageWithLayout } from "pages/_app";
import { ReactElement } from "react";
import qs from "qs";
import { ArticleCard } from "components/ArticleCard";

export const getStaticPaths: GetStaticPaths = async () => {
  const res = await strapiClient.get<StrapiListResponse<TopicType>>("topics");
  const paths = res.data.data.map((el) => ({
    params: {
      topic: el.attributes.name,
    },
  }));

  return { paths, fallback: false };
};

export const getStaticProps: GetStaticProps<{
  data: StrapiGetResponse<Topic>["data"];
}> = async (context) => {
  const query = qs.stringify(
    {
      populate: { articles: { populate: "*" } },
      filters: { name: context.params?.topic },
    },
    { encodeValuesOnly: true }
  );
  const res = await strapiClient.get<StrapiListResponse<Topic>>(
    `topics?${query}`
  );

  return {
    props: { data: res.data.data[0] },
  };
};

const Topic: NextPageWithLayout<
  InferGetStaticPropsType<typeof getStaticProps>
> = ({ data }) => {
  const articles = data.attributes.articles?.data;

  return (
    <Box>
      <Heading fontSize="2xl" mb={4}>
        Topic: {data.attributes.name}
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

Topic.getLayout = (page: ReactElement) => {
  return (
    <Layout>
      <Layout.Content>{page}</Layout.Content>
    </Layout>
  );
};

export default Topic;
