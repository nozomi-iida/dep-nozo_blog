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
import Link from "next/link";

export const Header = () => {
  const [isMenuOpen, setIsMenuOpen] = useBoolean(true);
  const [isSearchOpen, setIsSearchOpen] = useBoolean();
  const [isDarkTheme, setIsDarkTheme] = useBoolean();

  return (
    <Flex margin="auto" px={15} justify="space-between">
      <Link href="">
        <Heading lineHeight="100px">Nozo blog</Heading>
      </Link>
      <Flex gap={10}>
        <SlideFade in={isMenuOpen} offsetX={5} offsetY={0}>
          <Flex align="center" fontWeight="bold" h="full" gap={5}>
            <Link href="">
              <Text transition="color 0.2s" _hover={{ color: "blue.400" }}>
                Engineer
              </Text>
            </Link>
            <Link href="">
              <Text transition="color 0.2s" _hover={{ color: "blue.400" }}>
                Life
              </Text>
            </Link>
            <Link href="">
              <Text transition="color 0.2s" _hover={{ color: "blue.400" }}>
                About
              </Text>
            </Link>
          </Flex>
        </SlideFade>
        <Flex>
          <Center
            w={12}
            as="a"
            href="#"
            transition="color 0.2s"
            _hover={{ color: "blue.400" }}
            color={isMenuOpen ? "blue.400" : undefined}
            onClick={setIsMenuOpen.toggle}
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
                transition: "all 0.2s",
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
                transition: "all 0.2s",
              }}
            />
          </Center>
          <Center
            w={12}
            as="a"
            href="#"
            transition="color 0.2s"
            _hover={{ color: "blue.400" }}
            onClick={setIsDarkTheme.toggle}
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
                as="a"
                href="#"
                transition="color 0.2s"
                color={isSearchOpen ? "blue.400" : undefined}
                _hover={{ color: "blue.400" }}
                onClick={setIsSearchOpen.toggle}
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
  );
};
