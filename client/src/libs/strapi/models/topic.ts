import { RelationManyDataType } from "../types";
import { Article } from "./article";

export type Topic = {
  name: string;
  articles?: RelationManyDataType<Article>;
};
