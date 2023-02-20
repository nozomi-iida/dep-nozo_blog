import {
  Box,
  Button,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Heading,
  Input,
  Select,
  Switch,
} from "@chakra-ui/react";
import { zodResolver } from "@hookform/resolvers/zod";
import { AdminLayout } from "components/Layout/AdminLayout";
import { Topic } from "libs/api/models/topic";
import { restCli } from "libs/axios";
import { NextPageWithLayout } from "pages/_app.page";
import { ReactElement } from "react";
import { useForm } from "react-hook-form";
import useSWR from "swr";
import { z } from "zod";

const schema = z.object({
  title: z.string().min(1, { message: "title is required" }),
  content: z.string().min(1, { message: "content is required" }),
  topic: z.string().optional(),
  tags: z.string().array().max(3, { message: "select up to three tags" }),
  isPublic: z.boolean(),
});

type FormData = z.infer<typeof schema>;

const CreateArticlePage: NextPageWithLayout = () => {
  const fetchTopics = (url: string) =>
    restCli<{ topics: Topic[] }>(url).then((res) => res.data);
  const { data: topicData } = useSWR("/topics", fetchTopics);

  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<FormData>({ resolver: zodResolver(schema) });
  const onSubmit = handleSubmit(async (params) => {
    console.log(params);
  });

  return (
    <Box>
      <Heading>Create an article</Heading>
      <form onSubmit={onSubmit}>
        <FormControl>
          <FormLabel>Topic</FormLabel>
          <Select {...register("topic")} placeholder=" ">
            {topicData &&
              topicData.topics.map((topic) => (
                <option value={topic.topicID} key={topic.topicID}>
                  {topic.name}
                </option>
              ))}
          </Select>
        </FormControl>
        <FormControl>
          <FormLabel>Tags</FormLabel>
          {/* TODO: 実装 */}
        </FormControl>
        <FormControl isInvalid={!!errors.title?.message}>
          <FormLabel>Title</FormLabel>
          <Input {...register("title")} />
          <FormErrorMessage>{errors.title?.message}</FormErrorMessage>
        </FormControl>
        <FormControl isInvalid={!!errors.content?.message}>
          <FormLabel>Content</FormLabel>
          {/* TODO: 実装 */}
          <FormErrorMessage>{errors.content?.message}</FormErrorMessage>
        </FormControl>
        <FormControl>
          <FormLabel>Public</FormLabel>
          <Switch {...register("isPublic")} />
        </FormControl>
        <Button type="submit" isLoading={isSubmitting}>
          Create
        </Button>
      </form>
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
