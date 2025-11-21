import { error } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ locals }) => {
  const systemInfo = await locals.apiClient.getSystemInfo();
  if (!systemInfo.success) {
    throw error(systemInfo.error.code, { message: systemInfo.error.message });
  }

  return {
    systemInfo: systemInfo.data,
  };
};
