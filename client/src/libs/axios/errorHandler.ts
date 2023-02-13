import { useToast } from "@chakra-ui/react";
import { RestErrorResponse } from ".";

export const getRestErrorMessage = (error: any) => {
  const restError: RestErrorResponse = error;
  if (restError.message) {
    return restError.message;
  } else {
    return "Something Wrong!";
  }
};
