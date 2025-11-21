<script lang="ts">
  import { Button } from "@nanoteck137/nano-ui";
  import UploadModal from "./UploadModal.svelte";
  import Image from "./Image.svelte";
  import EditCollectionModal from "./EditCollectionModal.svelte";

  const { data } = $props();

  let openUploadModal = $state(false);
  let openEditModal = $state(false);
</script>

<p>{data.collection.title}</p>

<Button
  onclick={() => {
    openUploadModal = true;
  }}
>
  Upload
</Button>

<Button
  onclick={() => {
    openEditModal = true;
  }}
>
  Edit Modal
</Button>

<div class="flex flex-wrap justify-center gap-2">
  {#each data.images as image}
    <Image url={image.url} />
  {/each}
</div>

<UploadModal bind:open={openUploadModal} collectionId={data.collection.id} />
<EditCollectionModal bind:open={openEditModal} collection={data.collection} />
