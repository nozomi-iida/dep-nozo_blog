import { Button } from "@chakra-ui/react";
import { zodResolver } from "@hookform/resolvers/zod";
import { AdminRouter } from "components/AdminRouter";
import {
  ArticleForm,
  ArticleFormData,
  articleSchema,
} from "components/ArticleForm";
import { AdminLayout } from "components/Layout/AdminLayout";
import { Article } from "libs/api/models/article";
import { restAdminCli } from "libs/axios/restAdminCli";
import { useCustomToast } from "libs/chakra/useCustomToast";
import { useRouter } from "next/router";
import { NextPageWithLayout } from "pages/_app.page";
import { ReactElement, useEffect } from "react";
import { FormProvider, useForm } from "react-hook-form";

const EditArticleDPage: NextPageWithLayout = () => {
  const router = useRouter();
  const toast = useCustomToast();
  const { id } = router.query;
  const fetchArticle = (url: string) =>
    restAdminCli.get<Article>(url).then((res) => res.data);
  const methods = useForm<ArticleFormData>({
    resolver: zodResolver(articleSchema),
  });
  const onSubmit = methods.handleSubmit(async (params) => {
    try {
      await restAdminCli.patch(`/articles/${id}`, params);
      toast({ title: "Success", status: "success" });
    } catch (error) {}
  });

  useEffect(() => {
    (async () => {
      if (id) {
        const article = await fetchArticle(`/articles/${id}`);
        methods.reset({
          title: article.title,
          content: article.content,
          topicId: article.topic?.topicId,
          tagIds: article.tags.map((tag) => tag.tagId),
          isPublic: !!article?.publishedAt,
        });
      }
    })();
  }, [id]);

  return (
    <AdminRouter>
      <FormProvider {...methods}>
        <ArticleForm />
        <Button onClick={onSubmit} display="block" mx="auto">
          Edit
        </Button>
      </FormProvider>
    </AdminRouter>
  );
};

EditArticleDPage.getLayout = (page: ReactElement) => {
  return (
    <AdminLayout>
      <AdminLayout.Sidebar />
      <AdminLayout.Content>{page}</AdminLayout.Content>
    </AdminLayout>
  );
};

export default EditArticleDPage;
