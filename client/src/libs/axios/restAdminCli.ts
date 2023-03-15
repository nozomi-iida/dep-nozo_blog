import axios, { AxiosError } from "axios";
import { localStorageKeys } from "utils/localstorageKeys";

export type RestErrorResponse = {
  code: number;
  message: string;
  type: string;
};

export const restAdminCli = axios.create({
  baseURL: `${process.env.NEXT_PUBLIC_REST_API_URI}/api/v1/admin`,
});

restAdminCli.interceptors.request.use((config) => {
  if (config.headers && !config.headers?.Authorization) {
    const token = localStorage.getItem(localStorageKeys.JWT_TOKEN);
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
  }
  return config;
});

restAdminCli.interceptors.response.use(
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
