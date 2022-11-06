import { Box, HStack, Link, Text } from "@chakra-ui/react";
import { SiZenn, SiTwitter } from "react-icons/si";

export const Footer = () => {
  return (
    <Box
      borderTop="1px"
      borderColor="borderColor"
      position="relative"
      py={14}
      _before={{
        content: `""`,
        h: 0.5,
        w: 10,
        position: "absolute",
        top: -0.5,
        left: 0,
        backgroundColor: "black",
      }}
    >
      <HStack align="center" justify="space-between">
        <Text>
          Thanks for the visit{" "}
          <Text as="span" fontWeight="bold">
            Nozo Blog
          </Text>
        </Text>
        <HStack>
          <Link _hover={{ color: "activeColor" }}>
            <SiTwitter size={14} />
          </Link>
          <Link _hover={{ color: "activeColor" }}>
            <SiZenn size={14} />
          </Link>
        </HStack>
      </HStack>
    </Box>
  );
};
