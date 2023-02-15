import { Box, Heading } from "@chakra-ui/react";
import { AdminLayout } from "components/Layout/AdminLayout";
import { NextPageWithLayout } from "pages/_app.page";
import { ReactElement } from "react";

const CreateArticlePage: NextPageWithLayout = () => {
  return (
    <Box>
      <Heading>Create an article</Heading>
    </Box>
  );
};

CreateArticlePage.getLayout = (page: ReactElement) => {
  return (
    <AdminLayout>
      <AdminLayout.Content>{page}</AdminLayout.Content>
    </AdminLayout>
  );
};

export default CreateArticlePage;
