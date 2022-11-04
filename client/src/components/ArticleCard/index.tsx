import { StarIcon } from "@chakra-ui/icons";
import { Box, Flex, Heading, HStack, Text } from "@chakra-ui/react";
import Image from "next/image";
import BlogTest from "../../../public/blog_test.jpeg";

export const ArticleCard = () => {
  return (
    <Box>
      <Image alt="" src={BlogTest} width={300} height={200} />
      <Text fontSize="sm">2019, 10 31</Text>
      <Heading>Gallery format - Sticky post</Heading>
      <Text fontSize="md">
        Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo
        ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis
        dis parturient montes, nascetur ridiculus mus….{" "}
      </Text>
      <Text fontSize="md">Read more</Text>
      <Flex justify="right">
        <StarIcon />
        <Text fontSize="sm">100</Text>
      </Flex>
    </Box>
  );
};
