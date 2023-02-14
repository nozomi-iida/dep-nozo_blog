import { rest } from "msw";

export const handler = {
  signIn: {
    success: () => {
      return rest.post("/sign_in", (req, res, ctx) => {
        return res(
          ctx.status(200),
          ctx.json({
            token: "token",
          })
        );
      });
    },
    userNotFound: (message: string) =>
      rest.post("/sign_in", (_, res, ctx) => {
        return res(ctx.status(404), ctx.json({ error: { message } }));
      }),
  },
};
