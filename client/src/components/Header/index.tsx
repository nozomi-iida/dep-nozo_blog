import { CloseIcon, MoonIcon, SearchIcon, SunIcon } from "@chakra-ui/icons";
import {
  Box,
  Center,
  Fade,
  Flex,
  Heading,
  Input,
  InputGroup,
  InputRightElement,
  Popover,
  PopoverArrow,
  PopoverBody,
  PopoverContent,
  PopoverTrigger,
  SlideFade,
  Text,
  useBoolean,
} from "@chakra-ui/react";
import { pagesPath } from "libs/pathpida/$path";
import { strapiClient } from "libs/strapi/api/axios";
import { Topic } from "libs/strapi/models/topic";
import { StrapiListResponse } from "libs/strapi/types";
import Link from "next/link";
import { useEffect, useState } from "react";

export const Header = () => {
  const [isMenuOpen, setMenuOpen] = useBoolean(true);
  const [isSearchOpen, setSearchOpen] = useBoolean();
  const [isDarkTheme, setDarkTheme] = useBoolean();
  const [showShadow, setShowShadow] = useBoolean();
  const [topics, setTopics] = useState<string[]>([]);

  useEffect(() => {
    const addShadow = () => {
      if (window.pageYOffset === 0) {
        setShowShadow.off();
      } else {
        setShowShadow.on();
      }
    };

    window.addEventListener("scroll", addShadow);

    return () => window.removeEventListener("scroll", addShadow);
  }, [setShowShadow]);

  useEffect(() => {
    strapiClient.get<StrapiListResponse<Topic>>("topics").then((res) => {
      const topics = res.data.data.map((el) => el.attributes.name);
      setTopics(topics);
    });
  }, []);

  return (
    <Box
      backgroundColor="white"
      position="fixed"
      w="full"
      top={0}
      transition="box-shadow .2s"
      boxShadow={showShadow ? "md" : undefined}
      as="header"
      zIndex="sticky"
    >
      <Box maxW={970} mx="auto">
        <Flex margin="auto" px={4} justify="space-between">
          <Link href={pagesPath.$url()}>
            <Heading
              lineHeight={showShadow ? "70px" : "100px"}
              transition="line-height .2s ease"
            >
              Nozo blog
            </Heading>
          </Link>
          <Flex gap={10}>
            <SlideFade in={isMenuOpen} offsetX={5} offsetY={0}>
              <Flex align="center" fontWeight="bold" h="full" gap={5}>
                <Link href={pagesPath.$url()}>
                  <Text
                    transition="color 0.2s"
                    _hover={{ color: "activeColor" }}
                  >
                    Home
                  </Text>
                </Link>
                {topics.map((topic) => (
                  <Link
                    key={topic}
                    href={pagesPath.topics._topic(topic).$url()}
                  >
                    <Text
                      transition="color 0.2s"
                      _hover={{ color: "activeColor" }}
                    >
                      {topic}
                    </Text>
                  </Link>
                ))}
              </Flex>
            </SlideFade>
            <Flex>
              <Center
                w={12}
                as="button"
                transition="color 0.2s"
                _hover={{ color: "activeColor" }}
                color={isMenuOpen ? "activeColor" : undefined}
                onClick={setMenuOpen.toggle}
              >
                <Box
                  w={4}
                  h="3px"
                  backgroundColor="currentcolor"
                  position="absolute"
                  _before={{
                    h: "3px",
                    w: isMenuOpen ? 2 : 4,
                    mx: isMenuOpen ? 1 : 0,
                    content: `""`,
                    display: "block",
                    backgroundColor: "currentcolor",
                    position: "absolute",
                    top: "-7px",
                    transition: "width 0.2s, margin 0.2s",
                  }}
                  _after={{
                    h: "3px",
                    w: isMenuOpen ? 2 : 4,
                    mx: isMenuOpen ? 1 : 0,
                    content: `""`,
                    display: "block",
                    backgroundColor: "currentcolor",
                    position: "absolute",
                    top: "7px",
                    transition: "width 0.2s, margin 0.2s",
                  }}
                />
              </Center>
              <Center
                w={12}
                as="button"
                transition="color 0.2s"
                _hover={{ color: "activeColor" }}
                onClick={setDarkTheme.toggle}
              >
                <Fade in={isDarkTheme} hidden={!isDarkTheme}>
                  <SunIcon w={17} h={17} />
                </Fade>

                <Fade in={!isDarkTheme} hidden={isDarkTheme}>
                  <MoonIcon w={17} h={17} />
                </Fade>
              </Center>
              <Popover isOpen={isSearchOpen}>
                <PopoverTrigger>
                  <Center
                    w={12}
                    as="button"
                    transition="color 0.2s"
                    color={isSearchOpen ? "activeColor" : undefined}
                    _hover={{ color: "activeColor" }}
                    onClick={setSearchOpen.toggle}
                  >
                    <Box>
                      <Fade in={isSearchOpen} hidden={!isSearchOpen}>
                        <CloseIcon w={17} h={17} />
                      </Fade>
                      <Fade in={!isSearchOpen} hidden={isSearchOpen}>
                        <SearchIcon w={17} h={17} />
                      </Fade>
                    </Box>
                  </Center>
                </PopoverTrigger>
                <PopoverContent
                  onClick={(e) => e.stopPropagation()}
                  backgroundColor="black"
                  color="gray"
                  py={6}
                  px={7}
                >
                  <PopoverArrow backgroundColor="black" />
                  <PopoverBody>
                    <InputGroup>
                      {/* autoFocusできるようにしたい */}
                      <Input
                        variant="flushed"
                        placeholder="Enter your search query..."
                        _focusVisible={{ borderColor: "gray.200" }}
                      />
                      <InputRightElement>
                        <SearchIcon />
                      </InputRightElement>
                    </InputGroup>
                  </PopoverBody>
                </PopoverContent>
              </Popover>
            </Flex>
          </Flex>
        </Flex>
      </Box>
    </Box>
  );
};
