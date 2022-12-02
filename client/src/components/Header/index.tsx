import { CloseIcon, MoonIcon, SearchIcon, SunIcon } from "@chakra-ui/icons";
import {
  Box,
  Button,
  Center,
  Drawer,
  DrawerBody,
  DrawerCloseButton,
  DrawerContent,
  DrawerOverlay,
  Fade,
  Flex,
  Heading,
  Hide,
  Input,
  InputGroup,
  InputRightElement,
  Popover,
  PopoverArrow,
  PopoverBody,
  PopoverContent,
  PopoverTrigger,
  Show,
  SlideFade,
  Text,
  useBoolean,
  useBreakpointValue,
  useColorMode,
  useColorModeValue,
  VStack,
} from "@chakra-ui/react";
import { pagesPath } from "libs/pathpida/$path";
import { strapiClient } from "libs/strapi/api/axios";
import { Topic } from "libs/strapi/models/topic";
import { StrapiListResponse } from "libs/strapi/types";
import Link from "next/link";
import { useRouter } from "next/router";
import { FormEvent, useEffect, useState } from "react";
import useSWR from "swr";

export const Header = () => {
  const responsiveIsMenuOpen = useBreakpointValue({
    base: false,
    lg: true,
  });
  const [isMenuOpen, setMenuOpen] = useBoolean();
  const [isSearchOpen, setSearchOpen] = useBoolean();
  const [showShadow, setShowShadow] = useBoolean();
  const [keyword, setKeyword] = useState("");
  const { colorMode, toggleColorMode } = useColorMode();
  const router = useRouter();

  const color = useColorModeValue("white", "#18191b");
  const fetchTopics = () => {
    return strapiClient
      .get<StrapiListResponse<Topic>>("topics")
      .then((res) => res.data.data.map((el) => el.attributes.name));
  };
  const { data: topics } = useSWR("topics", fetchTopics);

  const onSearch = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    router.push(`${pagesPath.search.$url().pathname}?keyword=${keyword}`);
    setSearchOpen.off();
    setTimeout(() => {
      setKeyword("");
    }, 100);
  };

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
    if (responsiveIsMenuOpen) {
      setMenuOpen.on();
    } else {
      setMenuOpen.off();
    }
  }, [setMenuOpen, responsiveIsMenuOpen]);

  return (
    <Box
      backgroundColor={color}
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
              whiteSpace="nowrap"
            >
              Nozo blog
            </Heading>
          </Link>
          <Flex gap={10}>
            <Hide below="md">
              <SlideFade in={isMenuOpen} offsetX={5} offsetY={0}>
                <Flex align="center" fontWeight="bold" h="full" gap={5}>
                  <Link href={pagesPath.$url()}>
                    <Text
                      lineHeight={showShadow ? "70px" : "100px"}
                      transition="color 0.2s, line-height .2s ease"
                      fontSize="sm"
                      _hover={{ color: "activeColor" }}
                    >
                      Home
                    </Text>
                  </Link>
                  {topics?.map((topic) => (
                    <Link
                      key={topic}
                      href={pagesPath.topics._topic(topic).$url()}
                    >
                      <Text
                        lineHeight={showShadow ? "70px" : "100px"}
                        transition="color 0.2s, line-height .2s ease"
                        fontSize="sm"
                        _hover={{ color: "activeColor" }}
                      >
                        {topic}
                      </Text>
                    </Link>
                  ))}
                </Flex>
              </SlideFade>
            </Hide>
            <Flex>
              <Center
                w={12}
                as="button"
                transition="color 0.2s"
                _hover={{ color: "activeColor" }}
                color={isMenuOpen ? "activeColor" : undefined}
                onClick={() => {
                  setMenuOpen.toggle();
                }}
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
                onClick={toggleColorMode}
              >
                <Fade in={colorMode === "dark"} hidden={colorMode === "light"}>
                  <SunIcon w={17} h={17} />
                </Fade>

                <Fade in={colorMode === "light"} hidden={colorMode === "dark"}>
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
                    <form onSubmit={onSearch}>
                      <InputGroup>
                        {/* TODO: autoFocusできるようにしたい */}
                        <Input
                          variant="flushed"
                          placeholder="Enter your search query..."
                          autoFocus
                          _focusVisible={{ borderColor: "gray.200" }}
                          value={keyword}
                          onChange={(e) => setKeyword(e.target.value)}
                        />
                        <InputRightElement>
                          <Button
                            type="submit"
                            backgroundColor="transparent"
                            _hover={{ backgroundColor: "transparent" }}
                          >
                            <SearchIcon />
                          </Button>
                        </InputRightElement>
                      </InputGroup>
                    </form>
                  </PopoverBody>
                </PopoverContent>
              </Popover>
            </Flex>
          </Flex>
        </Flex>
      </Box>
      <Show below="md">
        <Drawer isOpen={isMenuOpen} onClose={setMenuOpen.off}>
          <DrawerOverlay />
          <DrawerContent>
            <DrawerCloseButton />
            <DrawerBody>
              <VStack gap="md">
                <Link href={pagesPath.$url()}>
                  <Text
                    transition="color 0.2s"
                    fontWeight="bold"
                    _hover={{ color: "activeColor" }}
                  >
                    Home
                  </Text>
                </Link>
                {topics?.map((topic) => (
                  <Link
                    key={topic}
                    href={pagesPath.topics._topic(topic).$url()}
                  >
                    <Text
                      transition="color 0.2s"
                      fontWeight="bold"
                      _hover={{ color: "activeColor" }}
                    >
                      {topic}
                    </Text>
                  </Link>
                ))}
              </VStack>
            </DrawerBody>
          </DrawerContent>
        </Drawer>
      </Show>
    </Box>
  );
};
