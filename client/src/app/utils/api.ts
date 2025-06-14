import axios from "axios";
import { ApiResponse, AssetPayload} from "../data/models";
import { UpdatePortfolioFormState, UpdatePortfolioSchema } from "./validation";

// here, all api calls are made to the server directly:

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

export async function updatePortfolio(
  state: UpdatePortfolioFormState,
  formData: FormData
): Promise<ApiResponse<UpdatePortfolioFormState> | undefined> {
  const portfolioRaw = formData.get('portfolio');
  const assetsRaw = formData.get('assets');

  const portfolio = portfolioRaw ? JSON.parse(portfolioRaw as string) : {};
  const assets = assetsRaw ? JSON.parse(assetsRaw as string) : {};

  const validatedData = UpdatePortfolioSchema.safeParse({
    portfolio: {
      id: portfolio.id,
      user_id: portfolio.user_id,
      name: portfolio.name,
      total_value: portfolio.total_value,
      investment_goals: portfolio.investment_goals,
      asset_preference: portfolio.asset_preference,
    },
    assets: {
      add: assets.add ?? [],
      update: assets.update ?? [],
      delete: assets.delete ?? [],
    },
  });

  if (!validatedData.success) {
    return {
      success: false,
      errors: validatedData.error.flatten().fieldErrors,
    };
  }

  console.log("Validated: ", validatedData);


  try {
    const response = await axios.put(
      `${process.env.NEXT_PUBLIC_API_URL}/user/portfolio`,
      validatedData.data,
      {
        headers: {
          'Content-Type': 'application/json',
        },
        withCredentials: true,
      }
    );

    return {
      success: true,
      data: response.data,
    };
  } catch (error) {
    return handleApiError(error);
  }
}

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
