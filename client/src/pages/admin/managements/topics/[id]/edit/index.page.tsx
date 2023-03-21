import { Box } from "@chakra-ui/react";
import { AdminRouter } from "components/AdminRouter";
import { AdminLayout } from "components/Layout/AdminLayout";
import { NextPageWithLayout } from "pages/_app.page";
import { ReactElement } from "react";

const EditTopicPage: NextPageWithLayout = () => {
  return (
    <AdminRouter>
      <Box>Edit</Box>
    </AdminRouter>
  );
};

EditTopicPage.getLayout = (page: ReactElement) => {
  return (
    <AdminLayout>
      <AdminLayout.Sidebar />
      <AdminLayout.Content>{page}</AdminLayout.Content>
    </AdminLayout>
  );
};

export default EditTopicPage;
