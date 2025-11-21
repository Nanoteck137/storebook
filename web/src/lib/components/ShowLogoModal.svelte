<script lang="ts">
  import type { Modal } from "$lib/components/modals";
  import { Button, Dialog, ScrollArea } from "@nanoteck137/nano-ui";

  export type Props = {
    logoUrl: string;
  };

  const {
    logoUrl,

    class: className,
    children,
    onResult,
  }: Props & Modal<void> = $props();

  let open = $state(false);
</script>

<Dialog.Root bind:open>
  <Dialog.Trigger class={className}>
    {@render children?.()}
  </Dialog.Trigger>

  <Dialog.Content>
    <ScrollArea class="max-h-[280px]">
      <img class="mx-auto max-h-48 object-contain" src={logoUrl} alt="Logo" />
    </ScrollArea>

    <Dialog.Footer>
      <Button
        variant="outline"
        onclick={() => {
          open = false;
        }}
      >
        Close
      </Button>
    </Dialog.Footer>
  </Dialog.Content>
</Dialog.Root>
