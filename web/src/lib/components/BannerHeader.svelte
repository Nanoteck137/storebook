<script lang="ts">
  import Badge from "$lib/components/Badge.svelte";
  import ShowLogoModal from "$lib/components/ShowLogoModal.svelte";
  import { cn } from "$lib/utils";
  import { Button, buttonVariants, Card } from "@nanoteck137/nano-ui";
  import { FileQuestion, Image as ImageIcon } from "lucide-svelte";
  import type { Snippet } from "svelte";

  type Props = {
    title: string;
    description?: string | null;
    coverUrl: string | null;
    bannerUrl: string | null;
    logoUrl: string | null;

    imageContent?: Snippet<[]>;
    buttons?: Snippet<[]>;
    underText?: Snippet<[]>;
  };

  const {
    title,
    description,
    coverUrl,
    bannerUrl,
    logoUrl,
    imageContent,
    buttons,
    underText,
  }: Props = $props();
</script>

<Card.Root>
  <div
    class="relative h-48 w-full overflow-hidden rounded-xl shadow-lg sm:h-60 md:h-72"
  >
    {#if bannerUrl}
      <img
        src={bannerUrl}
        alt="Collection Banner"
        class="h-full w-full object-cover"
      />

      <div
        class="absolute inset-0 rounded-lg border-b bg-black bg-opacity-40"
      ></div>
    {:else}
      <div
        class="h-full w-full bg-gradient-to-br from-[#9333EA] to-[#F59E0B]"
      ></div>
    {/if}

    {#if logoUrl}
      <div
        class="pointer-events-none absolute inset-0 hidden items-center justify-center md:flex"
      >
        <div
          class="w-40 rounded-lg border bg-black bg-opacity-60 p-4 shadow-lg"
        >
          <img
            src={logoUrl}
            alt="Collection Logo"
            class="mx-auto max-h-24 w-auto object-contain"
          />
        </div>
      </div>
    {/if}

    {#if logoUrl}
      <ShowLogoModal
        class={cn(
          "absolute right-4 top-4 z-20 md:hidden",
          buttonVariants({ variant: "secondary", size: "icon" }),
        )}
        {logoUrl}
        onResult={() => {}}
      >
        <ImageIcon />
      </ShowLogoModal>
    {/if}
  </div>

  <div
    class="relative -mt-20 flex flex-col items-center p-2 sm:flex-row sm:items-stretch"
  >
    <div
      class="relative z-10 aspect-[75/106] w-40 min-w-40 flex-shrink-0 overflow-hidden rounded-lg border bg-black shadow-lg sm:w-48 sm:min-w-48"
    >
      {#if coverUrl}
        <img
          src={coverUrl}
          alt="Collection Cover"
          class="h-full w-full object-cover"
        />
      {:else}
        <div class="flex h-full w-full items-center justify-center">
          <FileQuestion size={52} />
        </div>
      {/if}

      {#if imageContent}
        {@render imageContent()}
      {/if}
    </div>

    <div
      class="mt-2 flex w-full flex-col items-center justify-center gap-4 text-center sm:mt-20 sm:items-start sm:justify-between sm:pl-2 sm:text-left"
    >
      <div class="flex flex-col items-center gap-2 pl-2 sm:items-start">
        <h1 class="text-2xl font-bold sm:line-clamp-2 sm:pt-2" {title}>
          {title}
        </h1>

        {#if underText}
          {@render underText()}
        {/if}
      </div>

      {#if buttons}
        {@render buttons()}
      {/if}

      <!-- {#if description}
        <div class="flex flex-col gap-1">
          <p
            class={`text-ellipsis whitespace-pre-line text-sm ${!descriptionShowMore ? "line-clamp-4" : ""}`}
          >
            {description}
          </p>

          <Button
            class="w-fit"
            size="sm"
            variant="outline"
            onclick={() => {
              descriptionShowMore = !descriptionShowMore;
            }}
          >
            Show More
          </Button>
        </div>
      {/if} -->
      <!-- <p class="text-sm italic text-gray-400">
      Updated regularly with new releases and extras.
    </p> -->
    </div>
  </div>
</Card.Root>
