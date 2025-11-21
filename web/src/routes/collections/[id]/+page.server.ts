import { error } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ locals, params }) => {
  const collection = await locals.apiClient.getCollectionById(params.id);
  if (!collection.success) {
    throw error(collection.error.code, { message: collection.error.message });
  }

  const images = await locals.apiClient.getCollectionImages(params.id);
  if (!images.success) {
    throw error(images.error.code, { message: images.error.message });
  }

  return {
    collection: collection.data,
    images: images.data.images,
  };
};
