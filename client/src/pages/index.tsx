import { ReactElement } from "react";
import { ArticleCard } from "components/ArticleCard";
import { Layout } from "components/Layout";
import { NextPageWithLayout } from "./_app";
import { Grid } from "@chakra-ui/react";

const Home: NextPageWithLayout = () => {
  return (
    <Layout.Content>
      <Grid templateColumns="repeat(3, 1fr)" gap={10}>
        <ArticleCard />
        <ArticleCard />
        <ArticleCard />
        <ArticleCard />
        <ArticleCard />
        <ArticleCard />
        <ArticleCard />
        <ArticleCard />
      </Grid>
    </Layout.Content>
  );
};

// TODO: Layoutの状態が保持されているかの確認をする
Home.getLayout = (page: ReactElement) => {
  return <Layout>{page}</Layout>;
};

export default Home;
