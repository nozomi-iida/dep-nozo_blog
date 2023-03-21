import { Button, Heading, useToast, VStack } from "@chakra-ui/react";
import { zodResolver } from "@hookform/resolvers/zod";
import { AdminRouter } from "components/AdminRouter";
import {
  ArticleForm,
  ArticleFormData,
  articleSchema,
} from "components/ArticleForm";
import { AdminLayout } from "components/Layout/AdminLayout";
import { restAdminCli } from "libs/axios/restAdminCli";
import { NextPageWithLayout } from "pages/_app.page";
import { ReactElement } from "react";
import { FormProvider, useForm } from "react-hook-form";

const CreateArticlePage: NextPageWithLayout = () => {
  const toast = useToast();
  const methods = useForm<ArticleFormData>({
    resolver: zodResolver(articleSchema),
    defaultValues: { tagNames: [] },
  });

  const onSubmit = methods.handleSubmit(async (params) => {
    try {
      await restAdminCli.post("/articles", params);
      toast({ position: "top", title: "Article created", status: "success" });
    } catch (error) {
      toast({ position: "top", title: "Error", status: "error" });
    }
  });

  return (
    <AdminRouter>
      <FormProvider {...methods}>
        <Heading mb={6}>Create an article</Heading>
        <VStack gap={4}>
          <ArticleForm />
          <Button onClick={onSubmit} isLoading={methods.formState.isSubmitting}>
            Create
          </Button>
        </VStack>
      </FormProvider>
    </AdminRouter>
  );
};

CreateArticlePage.getLayout = (page: ReactElement) => {
  return (
    <AdminLayout>
      <AdminLayout.Sidebar />
      <AdminLayout.Content>{page}</AdminLayout.Content>
    </AdminLayout>
  );
};

export default CreateArticlePage;
