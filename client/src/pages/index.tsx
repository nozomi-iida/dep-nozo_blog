import { FC } from "react";
import { ArticleCard } from "../components/ArticleCard";
import { Header } from "../components/Header";
import { Layout } from "../components/Layout";

const Home: FC = () => {
  return (
    <Layout.Content>
      <Header />
      <ArticleCard />
    </Layout.Content>
  );
};

export default Home;
