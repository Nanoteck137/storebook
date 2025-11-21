import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ parent }) => {
  const data = await parent();

  if (data.userToken) {
    throw redirect(301, "/collections");
  }

  return {
    ...data,
  };
};
