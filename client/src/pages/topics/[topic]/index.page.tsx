import { Box, Heading, SimpleGrid } from "@chakra-ui/react";
import { Layout } from "components/Layout";
import { GetStaticPaths, GetStaticProps, InferGetStaticPropsType } from "next";
import { NextPageWithLayout } from "pages/_app.page";
import { ReactElement } from "react";
import { ArticleCard } from "components/ArticleCard";
import { useRouter } from "next/router";
import { NextHead } from "components/NextHead";
import { pagesPath } from "libs/pathpida/$path";
import { restCli } from "libs/axios";
import { ListTopicResponse, Topic, TopicQuery } from "libs/api/models/topic";

export const getStaticPaths: GetStaticPaths = async () => {
  const res = await restCli.get<ListTopicResponse>("/topics");

  const paths = res.data.topics.map((topic) => ({
    params: {
      topic: topic.name,
    },
  }));

  return { paths, fallback: false };
};

export const getStaticProps: GetStaticProps<{
  data: Topic;
}> = async (context) => {
  const query: TopicQuery = {
    associatedWith: "article",
  };
  const res = await restCli.get<Topic>(`/topics/${context.params?.topic}`, {
    params: query,
  });

  return {
    props: { data: res.data },
  };
};

const Topic: NextPageWithLayout<
  InferGetStaticPropsType<typeof getStaticProps>
> = ({ data }) => {
  const articles = data.articles;
  const router = useRouter();
  const topic = router.query.topic as string;

  return (
    <Box>
      <NextHead
        title={topic}
        url={pagesPath.topics._topic(topic).$url().pathname}
      />
      <Heading fontSize="2xl" mb={4}>
        Topic: {data?.name}
      </Heading>

      <SimpleGrid columns={{ sm: 1, md: 2, lg: 3 }} spacing={10}>
        {articles?.map((article) => (
          <ArticleCard key={article.articleId} article={article} />
        ))}
      </SimpleGrid>
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
