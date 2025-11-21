<script lang="ts">
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";
  import type { Page } from "$lib/api/types";
  import { Button, Pagination } from "@nanoteck137/nano-ui";
  import { ChevronLeft, ChevronRight } from "lucide-svelte";

  type Props = {
    pageData: Page;
  };

  const { pageData }: Props = $props();

  function gotoPage(p: number) {
    const query = $page.url.searchParams;
    query.set("page", (p - 1).toString());

    goto(`?${query.toString()}`, { invalidateAll: true, keepFocus: true });
  }
</script>

<div class="flex items-center justify-center gap-4 sm:hidden">
  <Button
    size="icon-lg"
    variant="ghost"
    onclick={() => gotoPage(pageData.page)}
    disabled={pageData.page <= 0}
  >
    <ChevronLeft />
  </Button>

  <p class="text-md px-3">{pageData.page + 1}</p>

  <Button
    size="icon-lg"
    variant="ghost"
    onclick={() => gotoPage(pageData.page + 2)}
    disabled={pageData.page >= pageData.totalPages - 1}
  >
    <ChevronRight />
  </Button>
</div>

<Pagination.Root
  class="hidden sm:flex"
  page={pageData.page + 1}
  count={pageData.totalItems}
  perPage={pageData.perPage}
  siblingCount={1}
  onPageChange={(p) => {
    const query = $page.url.searchParams;
    query.set("page", (p - 1).toString());

    goto(`?${query.toString()}`, { invalidateAll: true, keepFocus: true });
  }}
>
  {#snippet children({ pages, currentPage })}
    <Pagination.Content>
      <Pagination.Item>
        <Pagination.PrevButton class="sm:w-28" />
      </Pagination.Item>
      {#each pages as page (page.key)}
        {#if page.type === "ellipsis"}
          <Pagination.Item>
            <Pagination.Ellipsis />
          </Pagination.Item>
        {:else}
          <Pagination.Item>
            <Pagination.Link
              href="?page={page.value}"
              {page}
              isActive={currentPage === page.value}
            >
              {page.value}
            </Pagination.Link>
          </Pagination.Item>
        {/if}
      {/each}
      <Pagination.Item>
        <Pagination.NextButton class="sm:w-28" />
      </Pagination.Item>
    </Pagination.Content>
  {/snippet}
</Pagination.Root>
