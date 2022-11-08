import { Box, Heading, HStack, Text, VStack } from "@chakra-ui/react";
import { Image } from "components/Image";
import { Layout } from "components/Layout";
import dayjs from "dayjs";
import { strapiClient } from "libs/strapi/api/axios";
import { Article } from "libs/strapi/models/article";
import { StrapiGetResponse, StrapiListResponse } from "libs/strapi/types";
import { GetStaticPaths, GetStaticProps, InferGetStaticPropsType } from "next";
import { NextPageWithLayout } from "pages/_app";
import qs from "qs";
import { ReactElement } from "react";
import ReactMarkdown from "react-markdown";
import rehypeRaw from "rehype-raw";

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
  return (
    <Box>
      {data.attributes.thumbnail?.data?.attributes.url && (
        <Image
          src={`${process.env.NEXT_PUBLIC_STRAPI_URI}${data.attributes.thumbnail.data.attributes.url}`}
          alt={data.id.toString()}
          w="100%"
          h="100%"
        />
      )}
      <VStack gap={4} align="left">
        <HStack gap={2}>
          <Text fontSize="sm" color="subInfoText" fontWeight="bold">
            {dayjs(data.attributes.publishedAt).format("YYYY-MM-DD")}
          </Text>
          {data.attributes.topic?.name && (
            <Text fontSize="sm" color="subInfoText" fontWeight="bold">
              {data.attributes.topic.name}
            </Text>
          )}
        </HStack>
        <Heading>{data.attributes.title}</Heading>
        {/* TODO: underlineが未実装 */}
        <ReactMarkdown
          rehypePlugins={[rehypeRaw]}
          components={{
            h1({ children, ...props }) {
              return <Heading fontSize="3xl">{children}</Heading>;
            },
            h2({ children }) {
              return <Heading fontSize="2xl">{children}</Heading>;
            },
            h3({ children }) {
              return <Heading fontSize="xl">{children}</Heading>;
            },
            h4({ children }) {
              return <Heading fontSize="lg">{children}</Heading>;
            },
          }}
        >
          {data.attributes.content}
        </ReactMarkdown>
      </VStack>
    </Box>
  );
};

Article.getLayout = (page: ReactElement) => {
  return (
    <Layout>
      <Layout.Content>{page}</Layout.Content>
    </Layout>
  );
};

export default Article;
