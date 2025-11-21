<script lang="ts">
  import Badge from "$lib/components/Badge.svelte";
  import Image from "$lib/components/Image.svelte";
  import { parseUserList } from "$lib/types";
  import { Calendar, Clapperboard, Star } from "lucide-svelte";

  type Props = {
    href: string;
    coverUrl?: string | null;
    title: string;
    startDate?: string | null;
    partCount: number;
    score?: number | null;
    userList?: string | null;
  };

  const {
    href,
    coverUrl,
    title,
    startDate,
    partCount,
    score,
    userList,
  }: Props = $props();

  function getYear() {
    if (!startDate) return null;

    const d = new Date(startDate);
    return d.getFullYear();
  }

  const year = $derived(getYear());
</script>

<a
  {href}
  class="group relative block aspect-[75/106] w-[160px] min-w-[160px] transform cursor-pointer overflow-hidden rounded bg-gray-900 shadow-md transition-transform duration-300 hover:-translate-y-1 hover:shadow-lg"
>
  <!-- Badge -->
  {#if userList}
    {@const list = parseUserList(userList)}
    {#if list}
      <Badge class="absolute left-2 top-2 z-10" {list} />
    {/if}
  {/if}

  <div
    class="flex w-[160px] min-w-[160px] items-center justify-center overflow-hidden bg-gray-800 text-gray-400"
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
    <!-- <div
      class="mt-1 flex hidden justify-center space-x-4 text-sm text-gray-300"
    >
      {#if score}
        <div class="flex items-center space-x-1">
          <Star class="h-4 w-4 fill-yellow-400 text-yellow-400" />
          <span>{score.toFixed(2)}</span>
        </div>
      {/if}
      <div class="flex items-center space-x-1">
        <Clapperboard class="h-4 w-4" />
        <span>{partCount} parts</span>
      </div>
      {#if year}
        <div class="flex items-center space-x-1">
          <Calendar class="h-4 w-4" />
          <span>{year}</span>
        </div>
      {/if}
    </div> -->
  </div>
</a>
