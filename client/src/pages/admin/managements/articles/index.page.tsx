import { Box, Button, Heading, HStack } from "@chakra-ui/react";
import { AdminRouter } from "components/AdminRouter";
import { AdminLayout } from "components/Layout/AdminLayout";
import { pagesPath } from "libs/pathpida/$path";
import { useRouter } from "next/router";
import { NextPageWithLayout } from "pages/_app.page";
import { ReactElement } from "react";
import { AiOutlinePlus } from "react-icons/ai";

const ManagementArticlesPage: NextPageWithLayout = () => {
  const router = useRouter();
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
      </Box>
    </AdminRouter>
  );
};

ManagementArticlesPage.getLayout = (page: ReactElement) => {
  return (
    <AdminLayout>
      <AdminLayout.Content>{page}</AdminLayout.Content>
    </AdminLayout>
  );
};

export default ManagementArticlesPage;
