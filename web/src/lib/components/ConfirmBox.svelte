<script lang="ts">
  import type { Modal } from "$lib/components/modals";
  import { Button, Dialog } from "@nanoteck137/nano-ui";

  export type Props = {
    open: boolean;

    title: string;
    description?: string;
    confirmText?: string;
  };

  let {
    open = $bindable(),
    title,
    description,
    confirmText = "Confirm",
    onResult,
  }: Props & Modal<void> = $props();
</script>

<Dialog.Root bind:open>
  <Dialog.Content>
    <Dialog.Header>
      <Dialog.Title>{title}</Dialog.Title>
      {#if description}
        <Dialog.Description>
          {description}
        </Dialog.Description>
      {/if}
    </Dialog.Header>

    <Dialog.Footer class="gap-2 sm:gap-0">
      <Button
        variant="outline"
        onclick={() => {
          open = false;
        }}
      >
        Cancel
      </Button>

      <Button
        variant="destructive"
        onclick={() => {
          open = false;
          onResult();
        }}
      >
        {confirmText}
      </Button>
    </Dialog.Footer>
  </Dialog.Content>
</Dialog.Root>
