import { render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { ChakraProvider } from "@chakra-ui/react";
import SignInPage from "./index.page";

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
        <SignInPage />
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
        <SignInPage />
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
