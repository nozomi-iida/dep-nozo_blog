import { render, screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { TagInput } from ".";
import { setupServer } from "msw/node";
import { handler } from "mocks/api";
import { randomUUID } from "crypto";
import { ChakraProvider } from "@chakra-ui/react";

const mockTag = {
  tagID: randomUUID(),
  name: "mockTag",
};

const server = setupServer();
server.use(handler.getTags.success([mockTag]));
beforeAll(() => server.listen());
afterEach(() => server.resetHandlers());
afterAll(() => server.close());

describe("TagInput", () => {
  const user = userEvent.setup();
  const addTag = async (name = "tag") => {
    await user.type(screen.getByRole("combobox"), `${name}{enter}`);
  };
  const clickDelete = async (index = 0) => {
    await user.click(screen.getAllByRole("button", { name: /close/i })[index]);
  };

  it("add tag", async () => {
    render(
      <ChakraProvider>
        <TagInput />
      </ChakraProvider>
    );
    const name = "add";
    await addTag(name);
    expect(await screen.findByText(name)).toBeInTheDocument();
  });
  it.skip("delete tag", async () => {
    render(<TagInput />);
    const name = "add";
    await addTag(name);
    await clickDelete();
    expect(await screen.queryByText(name)).not.toBeInTheDocument();
  });
  it("delete tag when backspace", async () => {
    render(
      <ChakraProvider>
        <TagInput />
      </ChakraProvider>
    );
    const name = "add";
    await addTag(name);
    await user.type(screen.getByRole("combobox"), "{backspace}");
    expect(await screen.queryByText(name)).not.toBeInTheDocument();
  });
  it("show message when try to add 4tags", async () => {
    render(
      <ChakraProvider>
        <TagInput />
      </ChakraProvider>
    );
    for (let index = 0; index < 4; index++) {
      await addTag(`tag${index}`);
    }
    expect(
      await screen.findByText("cannot add more than three tags")
    ).toBeInTheDocument();
    await user.type(screen.getByRole("combobox"), "{backspace}");
    expect(
      await screen.queryByText("cannot add more than three tags")
    ).not.toBeInTheDocument();
  });
});
