import { render, screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { randomUUID } from "crypto";
import { handler } from "mocks/api";
import { setupServer } from "msw/node";
import CreateArticlePage from "./index.page";

const mockTopic = {
  topicID: randomUUID(),
  name: "mock topic",
  description: "",
};

const server = setupServer();
server.use(handler.getTopics.success([mockTopic]));
beforeAll(() => server.listen());
afterEach(() => server.resetHandlers());
afterAll(() => server.close());

describe("CreateArticlePage", () => {
  const user = userEvent.setup();
  const selectTopic = async (topic = mockTopic.name) => {
    await waitFor(() => screen.getByRole("option", { name: topic }));

    await user.selectOptions(
      screen.getByRole("combobox", {
        name: /topic/i,
      }),
      screen.getByRole("option", { name: topic })
    );
  };
  const addTags = async (tag = "tag") => {
    await user.type(screen.getByLabelText("Tag"), `${tag}{enter}`);
  };
  const typeTitle = async (title = "title") => {
    await user.type(screen.getByLabelText("Title"), title);
  };
  const typeContent = async (content = "content") => {
    await user.type(screen.getByLabelText("Content"), content);
  };
  const changeToDraft = async () => {
    await user.click(screen.getByLabelText("Public"));
  };
  const clickPublic = async () => {
    await user.click(screen.getByLabelText("Public"));
  };
  it("show validate error message", async () => {
    render(<CreateArticlePage />);
    await clickPublic();
    await waitFor(() => {
      expect(screen.findByText("content is required")).toBeInTheDocument();
    });
  });
  it("show validate message when add 4 tags", async () => {
    render(<CreateArticlePage />);
    await addTags("tag 1");
    await addTags("tag 2");
    await addTags("tag 3");
    await addTags("tag 4");
    expect(
      await screen.findByText("select up to three tags")
    ).toBeInTheDocument();
  });
  it("not add tag when try to add same name tag", async () => {
    render(<CreateArticlePage />);
    await addTags();
    await addTags();
  });
  it("remove tag", async () => {});
  it("success to create public article", async () => {
    render(<CreateArticlePage />);
    await selectTopic();
    await typeTitle();
    await clickPublic();
  });
  it("success to create draft article", async () => {
    render(<CreateArticlePage />);
    await selectTopic();
    await typeTitle();
  });
});
