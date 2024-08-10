import axios from "axios";

const SERVER_URL = import.meta.env.BROKER_HOST_FROM_FRONT
const SERVER_PORT = import.meta.env.BROKER_HOST
const axiosInstance = axios.create({
  baseURL: `${SERVER_URL}:$${SERVER_PORT}`,
});

axiosInstance.interceptors.request.use(
  (config) => {
    const storage: string | null = localStorage.getItem("token");
    const accessToken: string | null = JSON.parse(storage ?? "");

    if (accessToken) {
      if (config.headers) config.headers.token = accessToken;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

axiosInstance.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    return Promise.reject(error);
  }
);

export default axiosInstance;
