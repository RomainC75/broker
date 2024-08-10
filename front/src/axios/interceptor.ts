import { AuthenticationResult } from "@azure/msal-browser";
import axios from "axios";

const SERVER_PORT = process.env.BROKER_PORT
const axiosInstance = axios.create({
  baseURL: `http://localhost:${SERVER_PORT}`,
  // headers: {
  //   "Access-Control-Allow-Origin": "*",
  //   "Access-Control-Allow-Headers": "Origin, X-Requested-With, Content-Type, Accept"
  // }
});


axiosInstance.interceptors.request.use(
  (config) => {
    const storage: string | null = localStorage.getItem("token");
    console.log("-> access string ", storage)
    const accessToken: AuthenticationResult | null = JSON.parse(storage ?? "");
    console.log("é-< access token", accessToken)
    if (accessToken) {
      console.log("-> access token present !!!!!")
      if (config.headers) config.headers.Authorization = `Bearer ${accessToken.idToken ?? ""}`;
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
