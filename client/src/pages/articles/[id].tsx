import { Box } from "@chakra-ui/react";
import { Layout } from "components/Layout";
import { NextPage } from "next";
import { useRouter } from "next/router";

const Article: NextPage = () => {
  const router = useRouter();
  const { id } = router.query;
  return <Layout.Content>Hello</Layout.Content>;
};

export default Article;
