import { Box } from "@chakra-ui/react";
import { AdminRouter } from "components/AdminRouter";
import { NextPageWithLayout } from "pages/_app.page";

const Managements: NextPageWithLayout = () => {
  return (
    <AdminRouter>
      <Box>Managements</Box>
    </AdminRouter>
  );
};

export default Managements;
