import { render, screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { SignInForm } from ".";
import { ChakraProvider } from "@chakra-ui/react";

// TODO: 共通化したい
const pushMock = jest.fn();
jest.mock("next/router", () => ({
  useRouter() {
    return {
      route: "/",
      pathname: "/",
      query: "",
      asPath: "",
      push: pushMock,
    };
  },
}));
describe("sign in form", () => {
  const user = userEvent.setup();
  const typeUsername = async (username = "username") => {
    await user.type(screen.getByLabelText("Username"), username);
  };
  const typePassword = async (password = "password") => {
    await user.type(screen.getByLabelText("Password"), password);
  };
  const clickSignIn = async () => {
    await user.click(
      screen.getByRole("button", {
        name: /sign in/i,
      })
    );
  };
  test("Success to sign in", async () => {
    render(
      <ChakraProvider>
        <SignInForm />
      </ChakraProvider>
    );
    await typeUsername("hoge");
    await typePassword();
    await clickSignIn();
    expect(await screen.findByText("Success to sign in")).toBeInTheDocument();
    expect(pushMock).toBeCalled();
  });
  test("Show Validate Error", async () => {
    render(
      <ChakraProvider>
        <SignInForm />
      </ChakraProvider>
    );
    await clickSignIn();
    expect(
      await screen.findByText("Please enter your username")
    ).toBeInTheDocument();
    expect(
      await screen.findByText("Please enter your password")
    ).toBeInTheDocument();
  });
  test.skip("Show error message when value was incorrect", () => {
    typeUsername("InValid");
    typePassword();
    expect(screen.findByText("Username was incorrect")).toBeInTheDocument();
  });
});
