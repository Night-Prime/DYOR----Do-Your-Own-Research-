import axios from "axios";
import { ApiResponse, AssetPayload} from "../data/models";

// here, all api calls are made:

axios.defaults.withCredentials = true;

export const saveAssets = async (payload: AssetPayload): Promise<ApiResponse<AssetPayload>> => {
  try {
    const response = await axios.post(`${process.env.NEXT_PUBLIC_API_URL}/asset/create-asset`, payload, {
      withCredentials: true
    });

    return {
      success: true,
      data: response.data,
    };
  } catch (error) {
    return handleApiError(error);
  }
};




// API Utils function:
export const handleApiError = (error: unknown): ApiResponse<never> => {
  if (axios.isAxiosError(error)) {
    return {
      success: false,
      errors: error.response?.data?.errors || {
        general: [error.response?.data?.message || 'Request failed'],
      },
    };
  }
  return { success: false, errors: { general: ['Unknown error'] } };
};
