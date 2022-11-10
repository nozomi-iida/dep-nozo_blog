export type StrapiListResponse<T = unknown> = {
  data: {
    id: number;
    attributes: T;
  }[];
  meta: {
    pagination: {
      page: number;
      pageCount: number;
      pageSize: number;
      total: number;
    };
  };
};

export type StrapiGetResponse<T = unknown> = {
  data: {
    id: number;
    attributes: T;
  };
  meta: {};
};

// TODO: 必要に応じて追加
export type FileType = {
  data?: {
    attributes: {
      url: string;
    };
  };
};

export type RelationDataType<T> = {
  data?: {
    id: number;
    attributes: T;
  };
};

export type RelationManyDataType<T> = {
  data: {
    id: number;
    attributes: T;
  }[];
};
