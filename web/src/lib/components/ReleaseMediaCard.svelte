<script lang="ts">
  import { mediaReleaseStatus, mediaReleaseTypes } from "$lib/api-types";
  import type { MediaRelease } from "$lib/api/types";
  import Badge from "$lib/components/Badge.svelte";
  import Image from "$lib/components/Image.svelte";
  import { parseUserList } from "$lib/types";
  import { clamp, formatTimeDiff, getTimeDifference } from "$lib/utils";

  type Props = {
    href: string;
    coverUrl?: string | null;
    title: string;
    // startDate?: string | null;
    // partCount: number;
    // score?: number | null;
    userList?: string | null;

    release: MediaRelease;
  };

  const {
    href,
    coverUrl,
    title,
    // startDate,
    // partCount,
    // score,
    userList,
    release,
  }: Props = $props();

  const now = new Date();
  const percent = $derived(release!.currentPart / release!.numExpectedParts);
  const time = $derived(
    getTimeDifference(now, new Date(release!.nextAiring ?? "")),
  );
</script>

<a
  {href}
  class="group relative block aspect-[75/106] max-w-[240px] transform cursor-pointer overflow-hidden rounded bg-gray-900 shadow-md transition-transform duration-300 hover:-translate-y-1 hover:shadow-lg"
>
  <!-- Badge -->
  {#if userList}
    {@const list = parseUserList(userList)}
    {#if list}
      <Badge class="absolute left-2 top-2 z-10" {list} />
    {/if}
  {/if}

  <span
    class="absolute right-2 top-2 z-10 inline-block select-none rounded-full bg-gray-900 px-3 py-1 text-xs font-semibold"
  >
    {mediaReleaseTypes.find((i) => i.value === release.releaseType)?.label}
  </span>

  <span
    class="absolute right-2 top-10 z-10 inline-block select-none rounded-full bg-gray-900 px-3 py-1 text-xs font-semibold"
  >
    {mediaReleaseStatus.find((i) => i.value === release.status)?.label}
  </span>

  <div
    class="flex min-w-[240px] items-center justify-center overflow-hidden bg-gray-800 text-gray-400"
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
    class="absolute bottom-0 left-0 right-0 rounded-t-lg
                  border border-black/40 bg-black/50 p-3 text-center
                  backdrop-blur-lg"
  >
    <h2
      class="line-clamp-2 text-ellipsis text-lg font-semibold text-gray-100"
      {title}
    >
      {title}
    </h2>
    <!-- <div
      class="mt-1 flex justify-center space-x-4 text-sm text-gray-300"
    ></div> -->

    <div class="flex flex-col gap-2">
      <div class="flex flex-col gap-1">
        <p class="text-xs text-gray-400">
          {release!.currentPart} / {release!.numExpectedParts} eps
        </p>

        <div class="h-1 w-full rounded-full bg-gray-700">
          <div
            class="h-1 rounded-full bg-blue-500"
            style={`width: ${clamp(percent * 100, 0, 100).toFixed(0)}%;`}
          ></div>
        </div>
      </div>

      {#if release!.status !== "completed"}
        <div class="flex items-center gap-1">
          <p class="text-xs text-gray-300">Next in:</p>
          <p class="text-sm font-bold text-blue-400">
            {formatTimeDiff(time)}
          </p>
        </div>
      {/if}

      <!-- <p>{release!.status}</p>
      <p>{release!.releaseType}</p> -->
    </div>
  </div>
</a>
