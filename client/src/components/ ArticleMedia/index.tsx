import { Box, HStack, Text, textDecoration, VStack } from "@chakra-ui/react";
import { Image } from "components/Image";
import { useThemeColor } from "libs/chakra/theme";
import { pagesPath } from "libs/pathpida/$path";
import { Article } from "libs/strapi/models/article";
import Link from "next/link";
import { FC } from "react";

type ArticleMediaProps = {
  id: number;
  article: Article;
  popularNum: number;
};

export const ArticleMedia: FC<ArticleMediaProps> = ({
  id,
  article,
  popularNum,
}) => {
  const { bgColor } = useThemeColor();

  return (
    <Link href={pagesPath.articles._id(id).$url()}>
      <HStack
        gap={4}
        sx={{
          ":hover": {
            ".article_media_num": {
              boxShadow: "dark-lg",
            },
          },
        }}
      >
        {/* TODO: 画像を正方形にできない場合がある */}
        <Box position="relative">
          <Box
            position="absolute"
            top="-10px"
            left="-10px"
            w={7}
            h={7}
            lineHeight={7}
            borderRadius="full"
            textAlign="center"
            zIndex="docked"
            backgroundColor={bgColor}
            className="article_media_num"
          >
            <Text fontSize="sm" fontWeight="bold">
              {popularNum}
            </Text>
          </Box>
          <Image
            src={
              article.thumbnail?.data
                ? article.thumbnail.data.attributes.url
                : undefined
            }
            alt={article.title}
            w={100}
            h={100}
          />
          <Box
            position="absolute"
            top={0}
            left={0}
            w="full"
            h="full"
            _hover={{ opacity: 0.2, backgroundColor: "black" }}
          />
        </Box>
        <VStack justify="left">
          <Text
            fontWeight="bold"
            _hover={{
              color: "activeColor",
              textDecoration: "underline",
            }}
            overflow="hidden"
            sx={{
              WebkitLineClamp: 2,
              WebkitBoxOrient: "vertical",
              display: "-webkit-box",
            }}
          >
            {article.title}
          </Text>
          <HStack align="normal" w="full">
            {article.topic?.data && (
              <Text fontSize="xs" fontWeight="bold" color="subInfoText">
                {article.topic.data.attributes.name}
              </Text>
            )}
          </HStack>
        </VStack>
      </HStack>
    </Link>
  );
};
