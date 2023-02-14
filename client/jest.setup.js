import "@testing-library/jest-dom";

jest.spyOn(Storage.prototype, "setItem");
Storage.prototype.setItem = jest.fn();
jest.mock("next/router", () => ({
  useRouter() {
    return {
      push: pushMock,
    };
  },
}));
