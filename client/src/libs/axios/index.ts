import axios from "axios";

export const restCli = axios.create({
  baseURL: process.env.NEXT_PUBLIC_REST_API_URI,
});
