import { randomUUID } from "crypto";
import { rest } from "msw";
import { Topic } from "libs/api/models/topic";
import { Tag } from "libs/api/models/tag";

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
          topicId: randomUUID(),
          name: "topic 1",
          description: "topic 1",
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
  getTags: {
    success: (tags: Tag[] = [{ tagId: randomUUID(), name: "tag1" }]) => {
      return rest.get("/tags", (_, res, ctx) => {
        return res(
          ctx.status(200),
          ctx.json({
            tags,
          })
        );
      });
    },
  },
};
