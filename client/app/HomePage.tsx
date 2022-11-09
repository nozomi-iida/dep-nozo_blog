"use client";

import { Article } from "libs/strapi/models/article";
import { StrapiListResponse } from "libs/strapi/types";

const HomePage = ({ articles }: { articles: StrapiListResponse<Article> }) => {
  const { data } = articles;
  return (
    <div>
      {data.map((el) => (
        <div key={el.id}>{el.attributes.title}</div>
      ))}
    </div>
  );
};

export default HomePage;
