import { FC } from "react";
import { ArticleCard } from "../components/ArticleCard";
import { Layout } from "../components/Layout";

const Home: FC = () => {
  return (
    <Layout.Content>
      <ArticleCard />
    </Layout.Content>
  );
};

export default Home;
