import { strapiClient } from "libs/strapi/api/axios";
import HomePage from "../HomePage";

const getArticles = async () => {
  const res = await strapiClient.get(`articles`);
  return res.data;
};

const Page = async () => {
  const articles = await getArticles();
  return <HomePage articles={articles} />;
};

export default Page;
