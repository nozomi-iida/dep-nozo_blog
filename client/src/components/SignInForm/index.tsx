import { FormControl, FormLabel, Input } from "@chakra-ui/react";
import { FC } from "react";

export const SignInForm: FC = () => {
  return (
    <form>
      <FormControl>
        <FormLabel>Username</FormLabel>
        <Input />
      </FormControl>
    </form>
  );
};
