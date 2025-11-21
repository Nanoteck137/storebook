<script lang="ts">
  import Spacer from "$lib/components/Spacer.svelte";
  import StandardPagination from "$lib/components/StandardPagination.svelte";
  import { Button } from "@nanoteck137/nano-ui";
  import NewCollectionModal from "./NewCollectionModal.svelte";
  import { Plus } from "lucide-svelte";

  const { data } = $props();

  let openNewCollectionModal = $state(false);
</script>

<Spacer size="md" />

<div class="flex items-center justify-between">
  <h2 class="text-bold text-xl">
    Collections
    <Button
      variant="ghost"
      size="icon"
      onclick={() => {
        openNewCollectionModal = true;
      }}
    >
      <Plus />
    </Button>
  </h2>
  <p class="text-sm">{data.page.totalItems} collections(s)</p>
</div>

<Spacer size="md" />

<div class="flex flex-col gap-2">
  {#each data.collections as collection}
    <div class="group border-b p-2">
      <a
        class="group-hover:cursor-pointer group-hover:underline"
        href="/collections/{collection.id}">{collection.title}</a
      >
    </div>
  {/each}
</div>

<Spacer size="sm" />

<StandardPagination pageData={data.page} />

<NewCollectionModal bind:open={openNewCollectionModal} />
