import { FC, useEffect, useId, useMemo, useState } from "react";
import useSWR from "swr";
import { restCli } from "libs/axios";
import { Tag as TagType } from "libs/api/models/tag";
import { MultiValue, Select } from "chakra-react-select";
import { Box, Text } from "@chakra-ui/react";

type TagInputProps = {
  max?: number;
  value?: string[];
  onChange?: (value: string[]) => void;
};

type Option = { label: string; value: string };

export const TagInput: FC<TagInputProps> = ({
  max = 3,
  value: valueProps,
  onChange: onChangeProps,
}) => {
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
    if (options.length > max) {
      setMessage("cannot add more than three tags");
      setTagName("");
    } else {
      setTags(options);
      onChangeProps?.(options.map((option) => option.value));
    }
  };

  const onInputChange = (value: string) => {
    setTagName(value);
  };

  useEffect(() => {
    if (valueProps?.length) {
      const options = valueProps.map((value) => ({
        label: value,
        value,
      }));
      setTags(options);
    }
  }, []);

  return (
    <Box>
      <Select
        value={tags}
        onInputChange={onInputChange}
        onChange={onChange}
        isMulti
        backspaceRemovesValue
        options={options}
        placeholder=""
        instanceId={useId()}
      />
      <Text color="red.500">{message}</Text>
    </Box>
  );
};
