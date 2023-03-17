import { Tag } from "./tag";
import { Topic } from "./topic";

export enum ArticleOrderBy {
  PUBLISHED_AT_DESC = "published_at_desc",
  PUBLISHED_AT_ASC = "published_at_asc",
}

export type ArticleQueryParams = {
  orderBy?: ArticleOrderBy;
  keyword?: string;
};

export type Article = {
  articleId: string;
  title: string;
  content: string;
  publishedAt?: string;
  tags: Tag[];
  topic?: Topic;
};
