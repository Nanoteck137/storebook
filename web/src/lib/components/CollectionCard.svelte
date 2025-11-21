<script lang="ts">
  type Props = {
    href: string;
    coverUrl?: string | null;
    name: string;
  };

  const { href, coverUrl, name }: Props = $props();
</script>

<a
  {href}
  class="group relative block aspect-[75/106] max-w-[280px] transform cursor-pointer overflow-hidden rounded bg-gray-900 shadow-md transition-transform duration-300 hover:-translate-y-1 hover:shadow-lg"
>
  <div
    class="flex min-w-[280px] items-center justify-center overflow-hidden bg-gray-800 text-gray-400"
  >
    {#if coverUrl}
      <!-- svelte-ignore a11y_img_redundant_alt -->
      <img
        src={coverUrl}
        alt="Cover image"
        class="h-full w-full object-cover transition-transform duration-300 group-hover:scale-105"
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
      title="Fullmetal Alchemist: Brotherhood"
    >
      {name}
    </h2>
  </div>
</a>
