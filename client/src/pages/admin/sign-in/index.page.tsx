import { NextPageWithLayout } from "pages/_app.page";
import {
  Box,
  Button,
  Divider,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Input,
  Text,
  VStack,
} from "@chakra-ui/react";
import { pagesPath } from "libs/pathpida/$path";
import { useRouter } from "next/router";
import { ReactElement } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { AdminLayout } from "components/Layout/AdminLayout";
import { getRestErrorMessage } from "libs/axios/errorHandler";
import { useCustomToast } from "libs/chakra/useCustomToast";
import { localStorageKeys } from "utils/localstorageKeys";
import { restAuthCli } from "libs/axios/restAuthCli";

const schema = z.object({
  username: z.string().min(1, { message: "Please enter your username" }),
  password: z.string().min(1, { message: "Please enter your password" }),
});
type FormData = z.infer<typeof schema>;

const SignInPage: NextPageWithLayout = () => {
  const toast = useCustomToast();
  const router = useRouter();
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<FormData>({ resolver: zodResolver(schema) });
  const onSubmit = handleSubmit(async (params) => {
    try {
      const res = await restAuthCli.post<{ token: string }>("/sign_in", params);
      localStorage.setItem(
        localStorageKeys.JWT_TOKEN,
        res.data.token as string
      );
      toast({ title: "Success to sign in" });
      router.push(pagesPath.admin.managements.articles.$url());
    } catch (error) {
      toast({ title: getRestErrorMessage(error), status: "error" });
    }
  });

  return (
    <Box
      bgColor="white"
      w="400px"
      mx="auto"
      boxShadow="2xl"
      borderRadius="base"
    >
      <Text px={8} fontWeight="bold" fontSize="lg" py={6}>
        SIGN IN
      </Text>
      <Divider />
      <VStack px={8} gap={6} as="form" onSubmit={onSubmit} py={8}>
        <FormControl isInvalid={!!errors.username}>
          <FormLabel>Username</FormLabel>
          <Input {...register("username")} />
          {errors.username && (
            <FormErrorMessage>{errors.username.message}</FormErrorMessage>
          )}
        </FormControl>
        <FormControl isInvalid={!!errors.password}>
          <FormLabel>Password</FormLabel>
          <Input type="password" {...register("password")} />
          {errors.password && (
            <FormErrorMessage>{errors.password.message}</FormErrorMessage>
          )}
        </FormControl>
        {/* TODO make button style */}
        <Button type="submit" isLoading={isSubmitting}>
          Sign In
        </Button>
      </VStack>
    </Box>
  );
};

SignInPage.getLayout = (page: ReactElement) => {
  return (
    <AdminLayout>
      <AdminLayout.Content>{page}</AdminLayout.Content>
    </AdminLayout>
  );
};

export default SignInPage;
