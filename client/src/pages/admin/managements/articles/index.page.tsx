import { Box } from "@chakra-ui/react";
import { AdminRouter } from "components/AdminRouter";
import { AdminLayout } from "components/Layout/AdminLayout";
import { NextPageWithLayout } from "pages/_app.page";
import { ReactElement } from "react";

const ManagementArticlesPage: NextPageWithLayout = () => {
  return (
    <AdminRouter>
      <Box>Managements</Box>
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
