import axios from "./axios";
import type { CheckCreate, CheckCreateResponse } from "../types/check";

export const createCheck = (check: CheckCreate) =>
  axios.post<CheckCreateResponse>("/api/checks", check); 