<script lang="ts">
  import FullImageModal from "./FullImageModal.svelte";

  type Props = {
    url: string;
  };

  const { url }: Props = $props();

  let openFullImage = $state(false);
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
<img
  class="max-w-[320px]"
  src={url}
  alt=""
  onclick={() => {
    openFullImage = true;
  }}
/>

<FullImageModal bind:open={openFullImage} imageUrl={url} />
