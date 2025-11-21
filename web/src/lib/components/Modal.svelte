<script lang="ts">
  import type { Snippet } from "svelte";

  type Props = {
    open: boolean;
    onClose: () => void;

    children?: Snippet;
  };

  const { open, onClose, children }: Props = $props();

  $effect(() => {
    if (open) {
      document
        .getElementsByTagName("body")[0]
        .classList.add("overflow-y-hidden");
    } else {
      document
        .getElementsByTagName("body")[0]
        .classList.remove("overflow-y-hidden");
    }
  });
</script>

{#if open}
  <!-- svelte-ignore a11y_consider_explicit_label -->
  <button class="fixed inset-0 bg-black/70" onclick={() => onClose()}></button>
  {#if children}
    {@render children()}
  {/if}
{/if}
