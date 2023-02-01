import { render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { SignInForm } from ".";
import { ChakraProvider } from "@chakra-ui/react";

describe("sign in form", () => {
  const user = userEvent.setup();
  const clickSignIn = () => {
    user.click(screen.getByText("Sign In"));
  };
  const typeUsername = (username = "test") => {
    user.type(screen.getByLabelText("Username"), username);
  };
  const typePassword = (password = "password") => {
    user.type(screen.getByLabelText("Username"), password);
  };
  test.only("Success to sign in", async () => {
    render(
      <ChakraProvider>
        <SignInForm />
      </ChakraProvider>
    );
    typeUsername();
    typePassword();
    user.click(screen.getByText("Sign In"));
    const allByTitle = await screen.findAllByText("Success to sign in");
    expect(allByTitle).toHaveLength(1);
  });
  test("Show Validate Error", () => {
    clickSignIn();
    expect(screen.findByText("Please enter your username")).toBeInTheDocument();
    expect(screen.findByText("Please enter password")).toBeInTheDocument();
  });
  test("Show error message when value was incorrect", () => {
    typeUsername("InValid");
    typePassword();
    expect(screen.findByText("Username was incorrect")).toBeInTheDocument();
  });
});
