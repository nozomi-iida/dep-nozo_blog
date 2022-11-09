import { FileType, RelationDataType, RelationManyDataType } from "../types";
import { Tag } from "./tag";
import { Topic } from "./topic";

export type Article = {
  title: string;
  content: string;
  thumbnail?: FileType;
  publishedAt?: string;
  topic?: RelationDataType<Topic>;
  tags?: RelationManyDataType<Tag>;
};
