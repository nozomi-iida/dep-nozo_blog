import { Box, HStack, Text } from "@chakra-ui/react";
import { Image } from "components/Image";
import dayjs from "dayjs";
import { Article } from "libs/strapi/models/article";
import Link from "next/link";
import { FC } from "react";
import { pagesPath } from "libs/pathpida/$path";
import { useThemeColor } from "libs/chakra/theme";

type ArticleWidgetProps = {
  id: number;
  article: Article;
};

export const ArticleWidget: FC<ArticleWidgetProps> = ({ id, article }) => {
  const { bgColor } = useThemeColor();

  return (
    <Link href={pagesPath.articles._id(id).$url()}>
      <Box
        position="relative"
        sx={{
          ":hover": {
            ".article_widget_content_box": {
              bottom: 8,
            },
            ".article_widget_content_overlay": {
              opacity: 0.35,
            },
            ".article_widget_content_title": {
              textDecoration: "underline",
              color: "activeColor",
            },
          },
        }}
        cursor="pointer"
      >
        <Image
          src={
            article.thumbnail?.data
              ? article.thumbnail?.data?.attributes.url
              : undefined
          }
          w={300}
          h={200}
          alt={article.title}
        />
        <Box
          position="absolute"
          top={0}
          left={0}
          backgroundColor="black"
          w="full"
          h="full"
          opacity={0.15}
          className="article_widget_content_overlay"
          transition="opacity 0.2s"
        />
        <Box
          p={4}
          backgroundColor={bgColor}
          boxShadow="2xl"
          position="absolute"
          bottom={5}
          transition="bottom 0.2s"
          className="article_widget_content_box"
          w="80%"
        >
          <Text
            fontWeight="bold"
            className="article_widget_content_title"
            textOverflow="ellipsis"
            whiteSpace="nowrap"
            overflow="hidden"
          >
            {article.title}
          </Text>
          <HStack>
            <Text fontSize="xs" fontWeight="bold" color="subInfoText">
              {dayjs(article.publishedAt).format("YYYY-MM-DD")}
            </Text>
            {article.topic?.data && (
              <Text fontSize="xs" fontWeight="bold" color="subInfoText">
                {article.topic.data.attributes.name}
              </Text>
            )}
          </HStack>
        </Box>
      </Box>
    </Link>
  );
};
