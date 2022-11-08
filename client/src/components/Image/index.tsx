import { Box, BoxProps } from "@chakra-ui/react";
import NextImage, { ImageProps as NextImageProps } from "next/image";
import { FC } from "react";

type ImageProps = Pick<NextImageProps, "alt" | "src"> &
  Omit<BoxProps, "position">;

export const Image: FC<ImageProps> = ({ src, alt, ...props }) => {
  return (
    <Box position="relative" {...props}>
      <NextImage fill src={src} alt={alt} />
    </Box>
  );
};
