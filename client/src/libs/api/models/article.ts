import { Topic } from "./topic";

export type Article = {
  articleId: string;
  title: string;
  content: string;
  publishedAt?: string;
  tags: string[];
  topic?: Topic;
};
