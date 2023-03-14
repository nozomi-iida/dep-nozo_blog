import { Box, Button, Heading, HStack, SimpleGrid } from "@chakra-ui/react";
import { AdminRouter } from "components/AdminRouter";
import { ArticleCard } from "components/ArticleCard";
import { AdminLayout } from "components/Layout/AdminLayout";
import { Article } from "libs/api/models/article";
import { restCli } from "libs/axios";
import { pagesPath } from "libs/pathpida/$path";
import { useRouter } from "next/router";
import { NextPageWithLayout } from "pages/_app.page";
import { ReactElement } from "react";
import { AiOutlinePlus } from "react-icons/ai";
import useSWR from "swr";

const ManagementArticlesPage: NextPageWithLayout = () => {
  const router = useRouter();
  const fetchArticles = (url: string) =>
    restCli<{ articles: Article[] }>(url).then((res) => res.data);
  const { data: articleData } = useSWR("/articles", fetchArticles);
  console.log(articleData);

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
        <SimpleGrid columns={{ sm: 1, md: 2, lg: 3 }} spacing={10}>
          {/* {data.map((el) => (
          <ArticleCard key={el.id} articleId={el.id} article={el.attributes} />
        ))} */}
        </SimpleGrid>
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
