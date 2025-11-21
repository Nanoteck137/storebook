<script lang="ts">
  import { getApiClient, handleApiError } from "$lib";
  import { ProviderSearchResult } from "$lib/api/types";
  import FormItem from "$lib/components/FormItem.svelte";
  import Image from "$lib/components/Image.svelte";
  import type { Modal } from "$lib/components/modals";
  import {
    Button,
    Dialog,
    Input,
    Label,
    ScrollArea,
  } from "@nanoteck137/nano-ui";
  import { Check, Search } from "lucide-svelte";

  export type Props = {
    open: boolean;
    providerName: string;
    type: "media" | "collection" | "show";
    providerDisplayName: string;
  };

  let {
    open = $bindable(),
    type,
    providerName,
    providerDisplayName,

    onResult,
  }: Props & Modal<ProviderSearchResult[]> = $props();

  const apiClient = getApiClient();

  type SearchResult = {
    data: ProviderSearchResult;
    checked: boolean;
  };

  let results = $state<SearchResult[]>([]);
  let checkedItems = $derived(results.filter((i) => i.checked));

  $effect(() => {
    if (open) {
      results = [];
    }
  });

  function runSearch(query: string) {
    switch (type) {
      case "media":
        return apiClient.providerSearchMedia(providerName, {
          query: { query },
        });
      case "collection":
        return apiClient.providerSearchCollections(providerName, {
          query: { query },
        });
      case "show":
        return apiClient.providerSearchShows(providerName, {
          query: { query },
        });
    }
  }

  async function search(query: string) {
    const res = await runSearch(query);
    if (!res.success) {
      return handleApiError(res.error);
    }

    console.log(res.data);

    results = res.data.searchResults.map((d) => ({ data: d, checked: false }));
  }
</script>

<Dialog.Root bind:open>
  <Dialog.Content>
    <Dialog.Header>
      <Dialog.Title>{providerDisplayName}: Search</Dialog.Title>
    </Dialog.Header>

    <form
      onsubmit={(e) => {
        e.preventDefault();
        const data = new FormData(e.target as HTMLFormElement);
        const searchQuery = data.get("search");
        if (searchQuery) {
          search(searchQuery.toString());
        }
      }}
    >
      <FormItem>
        <Label for="search">Search</Label>
        <div class="flex items-center gap-2">
          <Input id="search" name="search" type="text" />

          <Button variant="ghost" size="icon" type="submit">
            <Search />
          </Button>
        </div>
        <!-- <Errors errors={$errors.name} /> -->
      </FormItem>
    </form>

    <ScrollArea class="h-60">
      <Label>Results</Label>

      <div class="flex flex-col gap-2">
        {#each results as result}
          <button
            class="group flex justify-between border-b py-2"
            onclick={() => {
              result.checked = !result.checked;
            }}
          >
            <div class="flex">
              <div class="relative h-20 w-14">
                <Image
                  class="h-20 min-w-14"
                  src={result.data.imageUrl}
                  alt="cover"
                />
                <div
                  class="absolute bottom-0 left-0 right-0 top-0 flex items-center justify-center group-hover:bg-black/40"
                >
                  {#if result.checked}
                    <div
                      class="flex h-8 w-8 items-center justify-center rounded-full bg-black/80"
                    >
                      <Check />
                    </div>
                  {/if}
                </div>
              </div>
              <div class="flex flex-col gap-2 px-4 py-1">
                <p
                  class="line-clamp-2 text-ellipsis text-start text-sm font-semibold"
                  title={result.data.title}
                >
                  {result.data.title}
                </p>
                <p class="text-start text-xs">ID: {result.data.providerId}</p>
              </div>
            </div>
            <!-- <div
              class="flex min-w-24 max-w-24 items-center justify-center border-l"
            >
              <Star size={18} class="fill-foreground" />
              <Spacer horizontal size="xs" />
              <p class="font-mono text-xs">
                {media.user?.score ?? "0"}
              </p>
            </div> -->
          </button>
        {/each}
      </div>
    </ScrollArea>

    <Dialog.Footer class="gap-2 sm:gap-0">
      <Button
        variant="outline"
        onclick={() => {
          open = false;
        }}
      >
        Close
      </Button>

      <Button
        disabled={checkedItems.length <= 0}
        onclick={() => {
          // editModalOpen = true;
          onResult(checkedItems.map((i) => i.data));
          open = false;
        }}
      >
        Add ({checkedItems.length} items)
      </Button>
    </Dialog.Footer>
  </Dialog.Content>
</Dialog.Root>
