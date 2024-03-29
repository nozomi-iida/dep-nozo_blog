import axios, { AxiosError } from "axios";
import { localStorageKeys } from "utils/localstorageKeys";

export type RestErrorResponse = {
  code: number;
  message: string;
  type: string;
};

export const restCli = axios.create({
  baseURL: `${process.env.NEXT_PUBLIC_REST_API_URI}/api/v1/public`,
});

restCli.interceptors.response.use(
  (res) => {
    return res;
  },
  (error: AxiosError<{ error: RestErrorResponse }>) => {
    const errorData = error.response?.data;
    if (errorData) {
      throw errorData.error;
    }
  }
);
