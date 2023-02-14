import { render, screen } from "@testing-library/react";
import { localStorageKeys } from "utils/localstorageKeys";
import { AdminRouter } from ".";

const pushMock = jest.fn();
jest.mock("next/router", () => ({
  useRouter() {
    return {
      push: pushMock,
    };
  },
}));

describe("admin router", () => {
  it("return sign in page when user does not sign in", async () => {
    render(
      <AdminRouter>
        <h1>Admin Page</h1>
      </AdminRouter>
    );
    expect(await screen.getByText(/loading\.\.\./i)).toBeInTheDocument();
    expect(pushMock).toBeCalled();
  });
  it("show page", async () => {
    const pageTitle = "Admin Page";
    jest.spyOn(Storage.prototype, "getItem").mockImplementation(() => {
      return "JWT_TOKEN";
    });
    render(
      <AdminRouter>
        <h1>{pageTitle}</h1>
      </AdminRouter>
    );
    expect(await screen.findByText(pageTitle)).toBeInTheDocument();
  });
});
