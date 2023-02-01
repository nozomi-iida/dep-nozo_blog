import { Box } from "@chakra-ui/react";
import { SignInForm } from "components/SignInForm";
import { NextPageWithLayout } from "pages/_app";

const SignIn: NextPageWithLayout = () => {
  return (
    <Box>
      <SignInForm />
    </Box>
  );
};

export default SignIn;
