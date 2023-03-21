import {
  Box,
  Button,
  Heading,
  HStack,
  Table,
  Tbody,
  Td,
  Th,
  Thead,
  Tr,
} from "@chakra-ui/react";
import { AdminRouter } from "components/AdminRouter";
import { AdminLayout } from "components/Layout/AdminLayout";
import { ListTopicResponse } from "libs/api/models/topic";
import { restAdminCli } from "libs/axios/restAdminCli";
import { pagesPath } from "libs/pathpida/$path";
import Link from "next/link";
import { useRouter } from "next/router";
import { NextPageWithLayout } from "pages/_app.page";
import { ReactElement } from "react";
import { AiOutlinePlus } from "react-icons/ai";
import useSWR from "swr";

const TopicIndexPage: NextPageWithLayout = () => {
  const router = useRouter();
  const fetchtopics = (url: string) =>
    restAdminCli.get<ListTopicResponse>(url).then((res) => res.data);
  const { data: topicData } = useSWR("/topics", fetchtopics);

  return (
    <AdminRouter>
      <Box>
        <HStack justify="space-between">
          <Heading>Topics</Heading>
          <Button
            onClick={() =>
              router.push(pagesPath.admin.managements.topics.create.$url())
            }
            leftIcon={<AiOutlinePlus />}
            colorScheme="primary"
          >
            Create
          </Button>
        </HStack>
        <Table>
          <Thead>
            <Tr>
              <Th>Title</Th>
              <Th>Action</Th>
            </Tr>
          </Thead>
          <Tbody>
            {topicData?.topics.map((topic) => (
              <Tr key={topic.topicId}>
                <Td>{topic.name}</Td>
                <Td>
                  <Link
                    href={pagesPath.admin.managements.topics
                      ._id(topic.topicId)
                      .edit.$url()}
                  >
                    <Button colorScheme="primary">編集</Button>
                  </Link>
                </Td>
              </Tr>
            ))}
          </Tbody>
        </Table>
      </Box>
    </AdminRouter>
  );
};

TopicIndexPage.getLayout = (page: ReactElement) => {
  return (
    <AdminLayout>
      <AdminLayout.Sidebar />
      <AdminLayout.Content>{page}</AdminLayout.Content>
    </AdminLayout>
  );
};

export default TopicIndexPage;
