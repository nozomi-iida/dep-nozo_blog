import { Box, Center, Spinner } from "@chakra-ui/react";
import { pagesPath } from "libs/pathpida/$path";
import { useRouter } from "next/router";
import { FC, ReactElement, useEffect, useState } from "react";
import { localStorageKeys } from "utils/localstorageKeys";

type AdminRouterProps = {
  children: ReactElement;
};
export const AdminRouter: FC<AdminRouterProps> = ({ children }) => {
  const [jwtToken, setJwtToken] = useState<string>();
  const router = useRouter();
  useEffect(() => {
    const jwtToken = localStorage.getItem(localStorageKeys.JWT_TOKEN);

    if (jwtToken) {
      setJwtToken(jwtToken);
    } else {
      router.push(pagesPath.admin.sign_in.$url());
    }
  }, []);

  return jwtToken ? (
    children
  ) : (
    <Center h="100vh">
      <Spinner size="xl" color="activeColor" label="Loading..." />
    </Center>
  );
};
