import {
  FormControl,
  FormErrorMessage,
  FormLabel,
  Input,
  Switch,
  Textarea,
  VStack,
} from "@chakra-ui/react";
import { Select } from "chakra-react-select";
import { Controller, useFormContext } from "react-hook-form";
import { TagInput } from "./TagInput";
import { z } from "zod";
import { useMemo } from "react";
import { restCli } from "libs/axios";
import { Topic } from "libs/api/models/topic";
import useSWR from "swr";

export const articleSchema = z.object({
  title: z.string().min(1, { message: "title is required" }),
  content: z.string().min(1, { message: "content is required" }),
  topicId: z.string().optional(),
  tagIds: z.string().array().max(3, { message: "select up to three tags" }),
  isPublic: z.boolean({ required_error: "isPublic is required" }),
});

export type ArticleFormData = z.infer<typeof articleSchema>;

export const ArticleForm = () => {
  const {
    register,
    control,
    formState: { errors },
  } = useFormContext<ArticleFormData>();
  const fetchTopics = (url: string) =>
    restCli<{ topics: Topic[] }>(url).then((res) => res.data);
  const { data: topicData } = useSWR("/topics", fetchTopics);
  const options = useMemo(() => {
    if (!topicData?.topics.length) return [];

    const topicDataOptions = topicData?.topics.map((topic) => ({
      label: topic.name,
      value: topic.topicId,
    }));

    return topicDataOptions;
  }, [topicData]);

  return (
    <VStack gap={4}>
      <FormControl isInvalid={!!errors.topicId?.message}>
        <FormLabel>Topic</FormLabel>
        {!!options.length && (
          <Controller
            name="topicId"
            control={control}
            render={({ field: { value, onChange } }) => (
              <Select
                value={options.find((option) => option.value === value)}
                placeholder=""
                options={options}
                onChange={(value) => {
                  onChange(value?.value);
                }}
              />
            )}
          />
        )}
        <FormErrorMessage>{errors.topicId?.message}</FormErrorMessage>
      </FormControl>
      <FormControl isInvalid={!!errors.tagIds?.length}>
        <FormLabel>Tags</FormLabel>
        <Controller
          name="tagIds"
          control={control}
          render={({ field: { value, onChange } }) => (
            <TagInput value={value} onChange={onChange} />
          )}
        />
        {Array.isArray(errors.tagIds) &&
          errors.tagIds?.map((error) => (
            <FormErrorMessage key={error?.message}>
              {error?.message}
            </FormErrorMessage>
          ))}
      </FormControl>
      <FormControl isInvalid={!!errors.title?.message}>
        <FormLabel>Title</FormLabel>
        <Input {...register("title")} />
        <FormErrorMessage>{errors.title?.message}</FormErrorMessage>
      </FormControl>
      <FormControl isInvalid={!!errors.content?.message}>
        <FormLabel>Content</FormLabel>
        <Textarea {...register("content")} rows={30} />
        <FormErrorMessage>{errors.content?.message}</FormErrorMessage>
      </FormControl>
      <FormControl>
        <FormLabel>Public</FormLabel>
        <Switch {...register("isPublic")} />
      </FormControl>
    </VStack>
  );
};
