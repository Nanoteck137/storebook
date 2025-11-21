<script lang="ts">
  import { getApiClient, handleApiError } from "$lib";
  import type { Provider } from "$lib/api/types";
  import toast from "svelte-5-french-toast";
  import { invalidateAll } from "$app/navigation";
  import { Button } from "@nanoteck137/nano-ui";
  import ProviderSearchModal from "./ProviderSearchModal.svelte";

  type Props = {
    provider: Provider;
  };

  const { provider }: Props = $props();
  const apiClient = getApiClient();

  let providerSearchMediaModalOpen = $state(false);
  let providerSearchCollectionModalOpen = $state(false);
  let providerSearchShowModalOpen = $state(false);
</script>

<div class="flex flex-col items-start gap-1">
  <p>{provider.displayName}</p>

  <div class="flex gap-2">
    <Button
      disabled={!provider.supports.searchMedia}
      onclick={() => {
        providerSearchMediaModalOpen = true;
      }}
    >
      Search
    </Button>

    <Button disabled={!provider.supports.getMedia} onclick={() => {}}>
      Get Media
    </Button>

    <Button
      disabled={!provider.supports.searchCollection}
      onclick={() => {
        providerSearchCollectionModalOpen = true;
      }}
    >
      Search Collection
    </Button>

    <Button disabled={!provider.supports.getCollection} onclick={() => {}}>
      Get Collection
    </Button>

    <Button
      disabled={!provider.supports.searchCollection}
      onclick={() => {
        providerSearchShowModalOpen = true;
      }}
    >
      Search Show
    </Button>

    <Button disabled={!provider.supports.getCollection} onclick={() => {}}>
      Get Show
    </Button>
  </div>
</div>

<ProviderSearchModal
  bind:open={providerSearchMediaModalOpen}
  type="media"
  providerName={provider.name}
  providerDisplayName={provider.displayName}
  onResult={async (items) => {
    const res = await apiClient.providerImportMedia(provider.name, {
      ids: items.map((i) => i.providerId),
    });
    if (!res.success) {
      return handleApiError(res.error);
    }

    toast.success("Successfully imported media");
    invalidateAll();
  }}
/>

<ProviderSearchModal
  bind:open={providerSearchCollectionModalOpen}
  type="collection"
  providerName={provider.name}
  providerDisplayName={provider.displayName}
  onResult={async (items) => {
    const res = await apiClient.providerImportCollections(provider.name, {
      ids: items.map((i) => i.providerId),
    });
    if (!res.success) {
      return handleApiError(res.error);
    }
    toast.success("Successfully imported collections");
    invalidateAll();
  }}
/>

<ProviderSearchModal
  bind:open={providerSearchShowModalOpen}
  type="collection"
  providerName={provider.name}
  providerDisplayName={provider.displayName}
  onResult={async (items) => {
    const res = await apiClient.providerImportShows(provider.name, {
      ids: items.map((i) => i.providerId),
    });
    if (!res.success) {
      return handleApiError(res.error);
    }
    toast.success("Successfully imported shows");
    invalidateAll();
  }}
/>
