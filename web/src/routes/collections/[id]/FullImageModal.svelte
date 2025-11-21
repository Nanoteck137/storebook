<script lang="ts">
  import type { Modal } from "$lib/components/modals";
  import { Button, Dialog, ScrollArea } from "@nanoteck137/nano-ui";

  export type Props = {
    open: boolean;
    imageUrl: string;
  };

  let { open = $bindable(), imageUrl }: Props = $props();
</script>

<Dialog.Root bind:open>
  <Dialog.Content class="max-w-fit overflow-clip p-0">
    <img class="mx-auto max-h-[95vh] object-contain" src={imageUrl} alt="" />
  </Dialog.Content>
</Dialog.Root>
