import { Box, Button, Heading, VStack } from "@chakra-ui/react";
import { zodResolver } from "@hookform/resolvers/zod";
import { AdminRouter } from "components/AdminRouter";
import { AdminLayout } from "components/Layout/AdminLayout";
import { TopicForm, TopicFormData, topicSchema } from "components/TafForm";
import { restAdminCli } from "libs/axios/restAdminCli";
import { useCustomToast } from "libs/chakra/useCustomToast";
import { NextPageWithLayout } from "pages/_app.page";
import { ReactElement } from "react";
import { FormProvider, useForm } from "react-hook-form";

const CreateTopicPage: NextPageWithLayout = () => {
  const toast = useCustomToast();
  const methods = useForm<TopicFormData>({
    resolver: zodResolver(topicSchema),
  });
  const onSubmit = methods.handleSubmit(async (params) => {
    try {
      await restAdminCli.post("/topics", params);
      toast({ title: "Topic created", status: "success" });
    } catch (error) {
      toast({ title: "Error", status: "error" });
    }
  });
  return (
    <AdminRouter>
      <FormProvider {...methods}>
        <Heading mb={6}>Create a topic</Heading>
        <VStack gap={4} w="full">
          <TopicForm />
          <Button onClick={onSubmit}>Create</Button>
        </VStack>
      </FormProvider>
    </AdminRouter>
  );
};

CreateTopicPage.getLayout = (page: ReactElement) => {
  return (
    <AdminLayout>
      <AdminLayout.Sidebar />
      <AdminLayout.Content>{page}</AdminLayout.Content>
    </AdminLayout>
  );
};

export default CreateTopicPage;
