import { render } from "@testing-library/react";
import EditArticleDPage from "./index.page";

describe("CreateArticlePage", () => {
  it("set default value", async () => {
    render(<EditArticleDPage />);
  });
});
