import { render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { ChakraProvider } from "@chakra-ui/react";
import SignInPage from "./index.page";
import { setupServer } from "msw/node";
import { handler } from "mocks/api";

// TODO: 共通化したい
const pushMock = jest.fn();
jest.mock("next/router", () => ({
  useRouter() {
    return {
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
  describe("post sign in", () => {
    const server = setupServer();
    beforeAll(() => server.listen());
    afterEach(() => server.resetHandlers());
    afterAll(() => server.close());
    test("Success to sign in", async () => {
      server.use(handler.signIn.success());
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
      expect(localStorage.setItem).toHaveBeenCalled();
    });
    test("Show error message when value was incorrect", async () => {
      const errorMessage = "Username was incorrect";
      server.use(handler.signIn.userNotFound(errorMessage));
      render(
        <ChakraProvider>
          <SignInPage />
        </ChakraProvider>
      );
      await typeUsername("InValid");
      await typePassword();
      await clickSignIn();
      expect(await screen.findByText(errorMessage)).toBeInTheDocument();
    });
  });
});
