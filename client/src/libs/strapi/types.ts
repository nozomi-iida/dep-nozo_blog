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
