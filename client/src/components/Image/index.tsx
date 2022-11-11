import { Box, BoxProps } from "@chakra-ui/react";
import { staticPath } from "libs/pathpida/$path";
import NextImage, { ImageProps as NextImageProps } from "next/image";
import { FC } from "react";

type ImageProps = Omit<BoxProps, "position"> & {
  src?: string;
  alt: string;
};

export const Image: FC<ImageProps> = ({ src, alt, ...props }) => {
  return (
    <Box position="relative" {...props}>
      <NextImage
        objectFit="contain"
        fill
        src={src ?? staticPath.blog_test_jpeg}
        alt={alt}
      />
    </Box>
  );
};
