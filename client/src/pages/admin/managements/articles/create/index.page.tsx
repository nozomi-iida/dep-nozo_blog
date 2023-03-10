import {
  Box,
  Button,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Heading,
  Input,
  Switch,
  Textarea,
  useToast,
  VStack,
} from "@chakra-ui/react";
import { zodResolver } from "@hookform/resolvers/zod";
import { Select } from "chakra-react-select";
import { AdminLayout } from "components/Layout/AdminLayout";
import { Topic } from "libs/api/models/topic";
import { restCli } from "libs/axios";
import { NextPageWithLayout } from "pages/_app.page";
import { ReactElement, useMemo } from "react";
import { Controller, useForm } from "react-hook-form";
import useSWR from "swr";
import { z } from "zod";
import { TagInput } from "./TagInput";

const schema = z.object({
  title: z.string().min(1, { message: "title is required" }),
  content: z.string().min(1, { message: "content is required" }),
  topic: z.string().optional(),
  tags: z.string().array().max(3, { message: "select up to three tags" }),
  isPublic: z.boolean({ required_error: "isPublic is required" }),
});

type FormData = z.infer<typeof schema>;

const CreateArticlePage: NextPageWithLayout = () => {
  const toast = useToast();
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

  const {
    register,
    handleSubmit,
    control,
    formState: { errors, isSubmitting },
  } = useForm<FormData>({
    resolver: zodResolver(schema),
    defaultValues: { tags: [] },
  });

  const onSubmit = handleSubmit(async (params) => {
    try {
      await restCli.post("/articles", params);
      toast({ position: "top", title: "Article created", status: "success" });
    } catch (error) {
      toast({ position: "top", title: "Error", status: "error" });
    }
  });

  return (
    <Box>
      <Heading mb={6}>Create an article</Heading>
      <VStack gap={4}>
        <FormControl>
          <FormLabel>Topic</FormLabel>
          <Controller
            name="topic"
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
        </FormControl>
        <FormControl>
          <FormLabel>Tags</FormLabel>
          <Controller
            name="tags"
            control={control}
            render={({ field: { value, onChange } }) => (
              <TagInput
                value={value}
                onChange={(value) => {
                  console.log("value", value);

                  onChange(value);
                }}
              />
            )}
          />
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
        <Button onClick={onSubmit} isLoading={isSubmitting}>
          Create
        </Button>
      </VStack>
    </Box>
  );
};

CreateArticlePage.getLayout = (page: ReactElement) => {
  return (
    <AdminLayout>
      <AdminLayout.Content>{page}</AdminLayout.Content>
    </AdminLayout>
  );
};

export default CreateArticlePage;
