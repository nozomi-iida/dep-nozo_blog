import { Article } from "./article";

export type Topic = {
  topicId: string;
  name: string;
  description: string;
  articles?: Article[];
};

export type TopicQuery = {
  associatedWith: "article";
};

export type ListTopicResponse = {
  topics: Topic[];
};
