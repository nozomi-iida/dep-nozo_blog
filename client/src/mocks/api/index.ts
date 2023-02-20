import { randomUUID } from "crypto";
import { rest } from "msw";
import { Topic } from "libs/api/models/topic";

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
  getTopics: {
    success: (
      topics: Topic[] = [
        {
          topicID: randomUUID(),
          name: "test 1",
          description: "test 1",
        },
      ]
    ) => {
      return rest.get("/topics", (_, res, ctx) => {
        return res(
          ctx.status(200),
          ctx.json({
            topics,
          })
        );
      });
    },
  },
};
