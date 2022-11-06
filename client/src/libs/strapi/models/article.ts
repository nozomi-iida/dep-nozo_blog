import { Tag } from "./tag";
import { Topic } from "./topic";

export type Article = {
  title: string;
  content: string;
  thumbnail?: FileType;
  publishedAt?: string;
  topic?: Topic;
  tags?: Tag[];
};

// TODO: 必要に応じて追加
type FileType = {
  data?: {
    attributes: {
      url: string;
    };
  };
};
