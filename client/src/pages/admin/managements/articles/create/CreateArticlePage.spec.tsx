import { render, screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import CreateArticlePage from "./index.page";

describe("CreateArticlePage", () => {
  const user = userEvent.setup();
  const selectTopic = async (topic = "topic") => {
    await user.selectOptions(screen.getByLabelText("Topic"), topic);
  };
  const addTags = async (tag = "tag") => {
    await user.type(screen.getByLabelText("Tag"), `${tag}{enter}`);
  };
  const typeContent = async (content = "content") => {
    await user.type(screen.getByLabelText("Content"), content);
  };
  const changeToDraft = async () => {
    await user.click(screen.getByLabelText("Public"));
  };
  const clickPublic = async () => {
    await user.click(screen.getByRole("button", { name: "create" }));
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
  it("success to create public article", () => {
    render(<CreateArticlePage />);
  });
  it("success to create draft article", () => {
    render(<CreateArticlePage />);
  });
});
