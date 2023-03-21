import {
  FormControl,
  FormErrorMessage,
  FormLabel,
  Input,
  Textarea,
  VStack,
} from "@chakra-ui/react";
import { useFormContext } from "react-hook-form";
import { z } from "zod";

export const topicSchema = z.object({
  name: z.string().min(1, { message: "name is required" }),
  description: z.string().min(1, { message: "description is required" }),
});

export type TopicFormData = z.infer<typeof topicSchema>;

export const TopicForm = () => {
  const {
    register,
    formState: { errors },
  } = useFormContext<TopicFormData>();
  return (
    <VStack gap={4} w="full">
      <FormControl isInvalid={!!errors.name?.message}>
        <FormLabel>Name</FormLabel>
        <Input {...register("name")} />
        <FormErrorMessage>{errors.name?.message}</FormErrorMessage>
      </FormControl>
      <FormControl isInvalid={!!errors.description?.message}>
        <FormLabel>Description</FormLabel>
        <Textarea {...register("description")} />
        <FormErrorMessage>{errors.description?.message}</FormErrorMessage>
      </FormControl>
    </VStack>
  );
};
