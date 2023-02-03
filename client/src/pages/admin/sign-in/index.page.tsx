import { NextPageWithLayout } from "pages/_app.page";
import {
  Box,
  Button,
  Center,
  Divider,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Heading,
  HStack,
  Input,
  Text,
  useToast,
  VStack,
} from "@chakra-ui/react";
import { pagesPath } from "libs/pathpida/$path";
import { useRouter } from "next/router";
import { FC, ReactElement } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { Layout } from "components/Layout";

const schema = z.object({
  username: z.string().min(1, { message: "Please enter your username" }),
  password: z.string().min(1, { message: "Please enter your password" }),
});
type FormData = z.infer<typeof schema>;

const SignInPage: NextPageWithLayout = () => {
  const toast = useToast();
  const router = useRouter();
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>({ resolver: zodResolver(schema) });
  const onSubmit = handleSubmit((params) => {
    toast({ title: "Success to sign in", status: "success" });
    router.push(pagesPath.admin.management.$url());
  });

  return (
    <Box px={8} w="400px" mx="auto" boxShadow="2xl" borderRadius="base">
      <Text fontWeight="bold" fontSize="lg" py={6} color="subInfoText">
        SIGN IN
      </Text>
      <Divider />
      <VStack gap={6} as="form" onSubmit={onSubmit} py={8}>
        <FormControl isInvalid={!!errors.username}>
          <FormLabel>Username</FormLabel>
          <Input {...register("username")} />
          {errors.username && (
            <FormErrorMessage>{errors.username.message}</FormErrorMessage>
          )}
        </FormControl>
        <FormControl isInvalid={!!errors.password}>
          <FormLabel>Password</FormLabel>
          <Input {...register("password")} />
          {errors.password && (
            <FormErrorMessage>{errors.password.message}</FormErrorMessage>
          )}
        </FormControl>
        <Button type="submit">Sign In</Button>
      </VStack>
    </Box>
  );
};

SignInPage.getLayout = (page: ReactElement) => {
  return (
    <Layout>
      <Layout.Content>{page}</Layout.Content>
    </Layout>
  );
};

export default SignInPage;
