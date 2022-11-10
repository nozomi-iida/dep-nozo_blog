import {
  Box,
  Divider,
  Heading,
  HStack,
  Link,
  Text,
  VStack,
} from "@chakra-ui/react";
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
import gfm from "remark-gfm";
import NextLink from "next/link";
import { AiFillCaretLeft, AiFillCaretRight } from "react-icons/ai";
import { pagesPath } from "libs/pathpida/$path";
import { Sidebar } from "components/Sidebar";

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
      <VStack gap={4} align="left" backgroundColor="white" p={8}>
        <HStack gap={2}>
          <Text fontSize="sm" color="subInfoText" fontWeight="bold">
            {dayjs(data.attributes.publishedAt).format("YYYY-MM-DD")}
          </Text>
          {data.attributes.topic?.data && (
            <Text fontSize="sm" color="subInfoText" fontWeight="bold">
              {data.attributes.topic?.data.attributes.name}
            </Text>
          )}
        </HStack>
        <Heading>{data.attributes.title}</Heading>
        {/* TODO: 画像は横幅いっぱいに表示したい */}
        <ReactMarkdown
          rehypePlugins={[rehypeRaw]}
          remarkPlugins={[gfm]}
          components={{
            h1({ children }) {
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
            a({ children }) {
              return (
                <Link
                  isExternal
                  display="block"
                  href={children.toString()}
                  _hover={{ color: "activeColor", textDecoration: "underline" }}
                >
                  {children}
                </Link>
              );
            },
          }}
        >
          {data.attributes.content}
        </ReactMarkdown>
        {data.attributes.tags?.data.length && (
          <HStack gap={0.5}>
            <Text>Tags:</Text>
            {data.attributes.tags.data.map((tag) => (
              // TODO: hrefを設置
              <NextLink key={tag.id} href={pagesPath.$url()}>
                <Text
                  display="inline"
                  _hover={{ color: "activeColor", textDecoration: "underline" }}
                >
                  #{tag.attributes.name}
                </Text>
              </NextLink>
            ))}
          </HStack>
        )}
        <Divider borderColor="borderColor" />
        {/* TODO: 次と前の記事のタイトルを取得し、表示したい */}
        <HStack justify="space-between">
          {/* TODO: Linkのコンポーネントを作る */}
          <NextLink href={pagesPath.articles._id(data.id).$url()}>
            <VStack>
              <Text
                fontSize="sm"
                color="subInfoText"
                fontWeight="bold"
                display="flex"
                alignItems="center"
              >
                <AiFillCaretLeft /> 前のページ
              </Text>
            </VStack>
          </NextLink>
          <NextLink href={pagesPath.articles._id(data.id).$url()}>
            <VStack>
              <Text
                fontSize="sm"
                color="subInfoText"
                fontWeight="bold"
                display="flex"
                alignItems="center"
              >
                次のページ <AiFillCaretRight />
              </Text>
            </VStack>
          </NextLink>
        </HStack>
      </VStack>
    </Box>
  );
};

Article.getLayout = (page: ReactElement) => {
  return (
    <Layout>
      <Layout.Content>{page}</Layout.Content>
      <Layout.Sidebar>
        <Sidebar />
      </Layout.Sidebar>
    </Layout>
  );
};

export default Article;
