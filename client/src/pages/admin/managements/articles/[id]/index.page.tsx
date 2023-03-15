import { Box } from "@chakra-ui/react";
import { AdminRouter } from "components/AdminRouter";
import { AdminLayout } from "components/Layout/AdminLayout";
import { NextPageWithLayout } from "pages/_app.page";
import { ReactElement } from "react";

const ArticleDetailPage: NextPageWithLayout = () => {
  return (
    <AdminRouter>
      <Box>ArticleDetailPage</Box>
    </AdminRouter>
  );
};

ArticleDetailPage.getLayout = (page: ReactElement) => {
  return (
    <AdminLayout>
      <AdminLayout.Sidebar />
      <AdminLayout.Content>{page}</AdminLayout.Content>
    </AdminLayout>
  );
};

export default ArticleDetailPage;
