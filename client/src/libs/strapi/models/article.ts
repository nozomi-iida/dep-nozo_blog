import { Tag } from "./tag";
import { Topic } from "./topic";

export type Article = {
  title: string;
  content: string;
  thumbnail?: string;
  publishedAt?: string;
  topic?: Topic;
  tags?: Tag[];
};
