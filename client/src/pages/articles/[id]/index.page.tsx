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
import { GetStaticPaths, GetStaticProps, InferGetStaticPropsType } from "next";
import { NextPageWithLayout } from "pages/_app.page";
import { ReactElement, useState } from "react";
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
import { restCli } from "libs/axios";
import { Article } from "libs/api/models/article";

export const getStaticPaths: GetStaticPaths = async () => {
  const res = await restCli.get<{ articles: Article[] }>("/articles");

  const paths = res.data.articles.map((article) => ({
    params: {
      id: article.articleId,
    },
  }));

  return { paths, fallback: false };
};

export const getStaticProps: GetStaticProps<{ article: Article }> = async (
  context
) => {
  const res = await restCli.get<Article>(`articles/${context.params?.id}`);

  return {
    props: {
      article: res.data,
    },
  };
};

const Article: NextPageWithLayout<
  InferGetStaticPropsType<typeof getStaticProps>
> = ({ article }) => {
  const [likeCount, setLikeCount] = useState(0);
  const { bgColor } = useThemeColor();
  const { colorMode } = useColorMode();

  const controls = useAnimationControls();
  const onClickClap = async () => {
    controls.start({ scale: 1.3 });
    setTimeout(() => {
      controls.start({ scale: 1.0 });
    }, 100);
    // const newLikeCount = likeCount + 1;
    // strapiClient
    //   .put(`articles/${article.id}`, {
    //     article: { likeCount: newLikeCount },
    //   })
    //   .then(() => {
    //     setLikeCount(newLikeCount);
    //   });
  };

  return (
    <Box>
      {/* <NextHead
        title={article.title}
        url={pagesPath.articles._id(article.articleId).$url().pathname}
        description={markdown2content(article.content)}
      /> */}
      {/* {article.thumbnail??.attributes.url && (
        <Image
          src={article.thumbnail.attributes.url}
          alt={article.id.toString()}
          w="full"
          h={430}
        />
      )} */}
      <VStack gap={4} align="left" backgroundColor={bgColor} p={8}>
        <HStack gap={2}>
          <Text fontSize="sm" color="subInfoText" fontWeight="bold">
            {dayjs(article.publishedAt).format("YYYY-MM-DD")}
          </Text>
          {article.topic && (
            <Text fontSize="sm" color="subInfoText" fontWeight="bold">
              {article.topic?.name}
            </Text>
          )}
        </HStack>
        <Heading>{article.title}</Heading>
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
          {article.content}
        </ReactMarkdown>
        {/* TODO: 共有機能をつけたい */}
        <HStack gap={8} w="full" justify="space-between">
          {article.tags?.length && (
            <HStack gap={0.5}>
              <Text>Tags:</Text>
              {article.tags.map((tag) => (
                // TODO: hrefを設置
                <NextLink key={tag.tagId} href={pagesPath.$url()}>
                  <Text
                    display="inline"
                    _hover={{
                      color: "activeColor",
                      textDecoration: "underline",
                    }}
                  >
                    #{tag.name}
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
                alt={article.title}
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
          <NextLink href={pagesPath.articles._id(article.articleId).$url()}>
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
          <NextLink href={pagesPath.articles._id(article.articleId).$url()}>
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
