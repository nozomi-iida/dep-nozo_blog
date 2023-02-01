import {
  Button,
  FormControl,
  FormLabel,
  Input,
  useToast,
} from "@chakra-ui/react";
import { pagesPath } from "libs/pathpida/$path";
import { useRouter } from "next/router";
import { FC, FormEvent } from "react";

export const SignInForm: FC = () => {
  const toast = useToast();
  const router = useRouter();
  const onSubmit = (e: FormEvent) => {
    e.preventDefault();
    toast({ title: "Success to sign in" });
    router.push(pagesPath.admin.management.$url());
  };
  return (
    <form onSubmit={onSubmit}>
      <FormControl>
        <FormLabel>Username</FormLabel>
        <Input />
      </FormControl>
      <FormControl>
        <FormLabel>Password</FormLabel>
        <Input />
      </FormControl>
      <Button type="submit">Sign In</Button>
    </form>
  );
};
