import { Tag } from "./tag";
import { Topic } from "./topic";

export type Article = {
  id: number;
  title: string;
  content: string;
  thumbnail?: string;
  publishedAt?: string;
  topic?: Topic;
  tags?: Tag[];
};
