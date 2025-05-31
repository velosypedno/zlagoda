import axios from "axios";

const instance = axios.create({
  baseURL: import.meta.env.VITE_FRONTEND_BASE_API_URL || "http://localhost:8080",
});

export default instance;
