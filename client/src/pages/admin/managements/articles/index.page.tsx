import {
  Box,
  Button,
  Heading,
  HStack,
  Table,
  Tbody,
  Td,
  Th,
  Thead,
  Tr,
} from "@chakra-ui/react";
import { AdminRouter } from "components/AdminRouter";
import { AdminLayout } from "components/Layout/AdminLayout";
import dayjs from "dayjs";
import { ListArticleResponse } from "libs/api/models/article";
import { restAdminCli } from "libs/axios/restAdminCli";
import { pagesPath } from "libs/pathpida/$path";
import Link from "next/link";
import { useRouter } from "next/router";
import { NextPageWithLayout } from "pages/_app.page";
import { ReactElement } from "react";
import { AiOutlinePlus } from "react-icons/ai";
import useSWR from "swr";

const ManagementArticlesPage: NextPageWithLayout = () => {
  const router = useRouter();
  const fetchArticles = (url: string) =>
    restAdminCli.get<ListArticleResponse>(url).then((res) => res.data);
  const { data: articleData } = useSWR("/articles", fetchArticles);

  return (
    <AdminRouter>
      <Box>
        <HStack justify="space-between">
          <Heading>Articles</Heading>
          <Button
            onClick={() =>
              router.push(pagesPath.admin.managements.articles.create.$url())
            }
            leftIcon={<AiOutlinePlus />}
          >
            Create
          </Button>
        </HStack>
        <Table>
          <Thead>
            <Tr>
              <Th>Title</Th>
              <Th>Published At</Th>
              <Th>Action</Th>
            </Tr>
          </Thead>
          <Tbody>
            {articleData?.articles.map((article) => (
              <Tr key={article.articleId}>
                <Td>{article.title}</Td>
                <Td>
                  {article.publishedAt
                    ? dayjs(article.publishedAt).format("YYYY-MM-DD")
                    : "UnPublished"}
                </Td>
                <Td>
                  <Link
                    href={pagesPath.admin.managements.articles
                      ._id(article.articleId)
                      .edit.$url()}
                  >
                    <Button>編集</Button>
                  </Link>
                </Td>
              </Tr>
            ))}
          </Tbody>
        </Table>
      </Box>
    </AdminRouter>
  );
};

ManagementArticlesPage.getLayout = (page: ReactElement) => {
  return (
    <AdminLayout>
      <AdminLayout.Sidebar />
      <AdminLayout.Content>{page}</AdminLayout.Content>
    </AdminLayout>
  );
};

export default ManagementArticlesPage;
