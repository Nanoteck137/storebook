<script lang="ts">
  import Badge from "$lib/components/Badge.svelte";
  import Image from "$lib/components/Image.svelte";
  import { parseUserList, type UserList } from "$lib/types";
  import { Calendar, Clapperboard, Star } from "lucide-svelte";

  type Props = {
    href: string;
    coverUrl?: string | null;
    title: string;
    score?: number | null;
    currentPart?: number | null;
    partCount?: number | null;
    list?: string | null;
  };

  const { href, coverUrl, title, score, currentPart, partCount, list }: Props =
    $props();
</script>

<a
  {href}
  class="group relative block aspect-[75/106] max-w-[210px] transform cursor-pointer overflow-hidden rounded bg-gray-900 shadow-md transition-transform duration-300 hover:-translate-y-1 hover:shadow-lg"
>
  <!-- Badge -->
  {#if list}
    {@const parsed = parseUserList(list)}
    {#if parsed}
      <Badge class="absolute left-2 top-2 z-10" list={parsed} />
    {/if}
  {/if}

  <!-- <button
    class="absolute right-2 top-2 z-10 rounded-full bg-black/60 p-2 text-white transition hover:bg-black/80"
  >
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="h-5 w-5"
      fill="none"
      viewBox="0 0 24 24"
      stroke="currentColor"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M6 12h.01M12 12h.01M18 12h.01"
      />
    </svg>
  </button> -->

  <div
    class="flex min-w-[210px] items-center justify-center overflow-hidden bg-gray-800 text-gray-400"
  >
    {#if coverUrl}
      <!-- svelte-ignore a11y_img_redundant_alt -->
      <Image
        src={coverUrl}
        alt="Cover image"
        class="aspect-[75/106] h-full w-full object-cover transition-transform duration-300 group-hover:scale-105"
      />
    {:else}
      <!-- TODO(patrik): Fix this -->
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-12 w-12"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V7M3 7l9 6 9-6"
        />
      </svg>
    {/if}
  </div>

  <div
    class="absolute bottom-0 left-0 right-0 rounded-t-lg border border-black/40 bg-black/50 p-3 text-center backdrop-blur-lg"
  >
    <h2
      class="line-clamp-2 text-ellipsis text-sm font-semibold text-gray-100"
      {title}
    >
      {title}
    </h2>
    <div class="mt-1 flex justify-center space-x-4 text-sm text-gray-300">
      <div class="flex items-center space-x-1">
        <Star class="h-4 w-4 fill-yellow-400 text-yellow-400" />
        <span>{score ?? 0}</span>
      </div>
      <div class="flex items-center space-x-1">
        <Clapperboard class="h-4 w-4" />
        <span>{currentPart ?? 0} / {partCount ?? 0}</span>
      </div>
      <!-- {#if year}
        <div class="flex items-center space-x-1">
          <Calendar class="h-4 w-4" />
          <span>{year}</span>
        </div>
      {/if} -->
    </div>
  </div>
</a>
