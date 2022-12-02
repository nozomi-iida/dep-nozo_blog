import { staticPath } from "libs/pathpida/$path";
import Head from "next/head";
import { FC } from "react";

type NextHeadProps = {
  title?: string;
  description?: string;
  url: string;
  imageUrl?: string;
  type?: "blog" | "article";
};

export const NextHead: FC<NextHeadProps> = ({
  title = "Nozo Blog",
  description,
  url,
  imageUrl = staticPath.nozomi_work_jpeg,
  type = "blog",
}) => {
  return (
    <Head>
      <title>{title}</title>
      <meta property="description" content={description} />
      <meta property="og:type" content={type} />
      <meta property="og:title" content={title} />
      <meta property="og:site_name" content={title} />
      <meta property="og:description" content={description} />
      <meta property="og:image" content={imageUrl} />
      <meta property="og:url" content={url} />
      <meta property="twitter:card" content="summary_large_imagerl" />
      <meta property="twitter:site" content="@nozo_i0106" />
    </Head>
  );
};
