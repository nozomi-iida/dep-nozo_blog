import axios from "axios";

export const strapiClient = axios.create({
  baseURL: `${process.env.STRAPI_URI}/api/`,
  headers: {
    Authorization: `Bearer ${process.env.STRAPI_TOKEN}`,
  },
});
