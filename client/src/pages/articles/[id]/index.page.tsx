import {
  Box,
  Divider,
  Heading,
  HStack,
  Link,
  Text,
  useColorMode,
  VStack,
} from "@chakra-ui/react";
import { Image } from "components/Image";
import { Layout } from "components/Layout";
import dayjs from "dayjs";
import { strapiClient } from "libs/strapi/api/axios";
import { Article } from "libs/strapi/models/article";
import { StrapiGetResponse, StrapiListResponse } from "libs/strapi/types";
import { GetStaticPaths, GetStaticProps, InferGetStaticPropsType } from "next";
import { NextPageWithLayout } from "pages/_app.page";
import qs from "qs";
import { ReactElement, useEffect, useState } from "react";
import ReactMarkdown from "react-markdown";
import rehypeRaw from "rehype-raw";
import gfm from "remark-gfm";
import NextLink from "next/link";
import { AiFillCaretLeft, AiFillCaretRight } from "react-icons/ai";
import { pagesPath, staticPath } from "libs/pathpida/$path";
import { Sidebar } from "components/Sidebar";
import { motion, useAnimationControls } from "framer-motion";
import { useThemeColor } from "libs/chakra/theme";
import { NextHead } from "components/NextHead";
import { markdown2content } from "utils/helpers";

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
  const [likeCount, setLikeCount] = useState(0);
  const { bgColor } = useThemeColor();
  const { colorMode } = useColorMode();

  const controls = useAnimationControls();
  const onClickClap = async () => {
    controls.start({ scale: 1.3 });
    setTimeout(() => {
      controls.start({ scale: 1.0 });
    }, 100);
    const newLikeCount = likeCount + 1;
    strapiClient
      .put(`articles/${data.id}`, {
        data: { likeCount: newLikeCount },
      })
      .then(() => {
        setLikeCount(newLikeCount);
      });
  };

  useEffect(() => {
    setLikeCount(Number(data.attributes.likeCount));
  }, [data]);

  return (
    <Box>
      <NextHead
        title={data.attributes.title}
        imageUrl={data.attributes.thumbnail?.data?.attributes.url}
        url={pagesPath.articles._id(data.id).$url().pathname}
        description={markdown2content(data.attributes.content)}
      />
      {data.attributes.thumbnail?.data?.attributes.url && (
        <Image
          src={data.attributes.thumbnail.data.attributes.url}
          alt={data.id.toString()}
          w="full"
          h={430}
        />
      )}
      <VStack gap={4} align="left" backgroundColor={bgColor} p={8}>
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
        {/* TODO: 共有機能をつけたい */}
        <HStack gap={8} w="full" justify="space-between">
          {data.attributes.tags?.data.length && (
            <HStack gap={0.5}>
              <Text>Tags:</Text>
              {data.attributes.tags.data.map((tag) => (
                // TODO: hrefを設置
                <NextLink key={tag.id} href={pagesPath.$url()}>
                  <Text
                    display="inline"
                    _hover={{
                      color: "activeColor",
                      textDecoration: "underline",
                    }}
                  >
                    #{tag.attributes.name}
                  </Text>
                </NextLink>
              ))}
            </HStack>
          )}
          <HStack
            w="full"
            justify="right"
            cursor="pointer"
            onClick={onClickClap}
            userSelect="none"
          >
            <motion.button animate={controls}>
              <Image
                alt={data.attributes.title}
                src={
                  colorMode === "light"
                    ? staticPath.clap_png
                    : staticPath.clap_dark_png
                }
                h={6}
                w={6}
              />
            </motion.button>
            <Text fontSize="sm">{likeCount}</Text>
          </HStack>
        </HStack>
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
