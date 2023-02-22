import {
  ChangeEvent,
  FC,
  KeyboardEvent,
  useEffect,
  useMemo,
  useRef,
  useState,
} from "react";
import useSWR from "swr";
import { restCli } from "libs/axios";
import { Tag as TagType } from "libs/api/models/tag";
import { MultiValue, Select } from "chakra-react-select";
import { Box, Text } from "@chakra-ui/react";

type TagInputProps = {
  max?: number;
};

type Option = { label: string; value: string };

export const TagInput: FC<TagInputProps> = ({ max = 3 }) => {
  const fetchTopics = (url: string) =>
    restCli<{ tags: TagType[] }>(url).then((res) => res.data);
  const { data: tagData } = useSWR("/tags", fetchTopics);

  const [tags, setTags] = useState<MultiValue<Option>>([]);
  const [tagName, setTagName] = useState("");
  const [message, setMessage] = useState("");
  const options = useMemo(() => {
    if (!tagData?.tags.length) return [];

    const tagDataOptions = tagData?.tags.map((tag) => ({
      label: tag.name,
      value: tag.name,
    }));
    tagDataOptions?.push({ value: tagName, label: tagName });
    return tagDataOptions;
  }, [tagData, tagName]);

  const onChange = (options: MultiValue<Option>) => {
    if (options.length >= 3) {
      setMessage("cannot add more than three tags");
      setTagName("");
    } else {
      setTags(options);
    }
  };
  const onKeydown = (e: KeyboardEvent) => {
    if (e.key === "Backspace" && !tagName) {
      const cloneTags = tags.concat();
      cloneTags.pop();

      setTags(cloneTags);
      setMessage("");
    }
  };

  const onInputChange = (value: string) => {
    setTagName(value);
    // setMessage("");
  };

  // const onTagDelete = (index: number) => {
  //   const filteredTags = tags.filter((_, idx) => idx !== index);
  //   setTags(filteredTags);
  // };

  // const onClickMenuItem = (name: string) => {
  //   addTags(name);
  // };

  // useEffect(() => {
  //   if (tags.length < 3) {
  //     setMessage("");
  //   }
  // }, [tags]);

  return (
    <Box>
      <Select
        value={tags}
        onInputChange={onInputChange}
        onChange={onChange}
        onKeyDown={onKeydown}
        isMulti
        backspaceRemovesValue
        options={options}
        placeholder=""
      />
      <Text color="red.500">{message}</Text>
    </Box>
  );
};
