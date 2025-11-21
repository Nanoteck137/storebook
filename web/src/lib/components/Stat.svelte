<script lang="ts">
  import {
    mediaTypes,
    mediaUserLists,
    statNames,
    statTypes,
  } from "$lib/api-types";
  import type { MainStat } from "$lib/api/types";
  import { Popover } from "@nanoteck137/nano-ui";

  type Props = {
    stat: MainStat;
  };

  const { stat }: Props = $props();
</script>

{#if stat.sub.length > 0}
  <Popover.Root>
    <Popover.Trigger class="flex flex-col items-center">
      <p class="text-xl">{stat.main.value}</p>
      <p class="text-center text-xs">
        {statNames.find((i) => i.value === stat.main.name)?.label}
      </p>
    </Popover.Trigger>
    <Popover.Content>
      <div class="flex flex-col gap-2">
        {#each stat.sub as sub}
          <p>
            {statTypes.find((i) => i.value === sub.name)?.label}: {sub.value}
          </p>
        {/each}
      </div>
    </Popover.Content>
  </Popover.Root>
{:else}
  <div class="flex flex-col items-center">
    <p class="text-xl">{stat.main.value}</p>
    <p class="text-center text-xs">
      {statNames.find((i) => i.value === stat.main.name)?.label}
    </p>
  </div>
{/if}
