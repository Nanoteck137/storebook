<script lang="ts">
  import { goto, invalidateAll } from "$app/navigation";
  import { getApiClient, handleApiError } from "$lib";
  import Errors from "$lib/components/Errors.svelte";
  import FormItem from "$lib/components/FormItem.svelte";
  import { Button, Dialog, Input, Label } from "@nanoteck137/nano-ui";
  import toast from "svelte-5-french-toast";
  import { zod } from "sveltekit-superforms/adapters";
  import { defaults, fileProxy, superForm } from "sveltekit-superforms/client";
  import { z } from "zod";
  import Spinner from "$lib/components/Spinner.svelte";

  const Schema = z.object({
    file: z.instanceof(File, { message: "Please upload a file." }),
  });

  export type Props = {
    open: boolean;
    collectionId: string;
  };

  let { open = $bindable(), collectionId }: Props = $props();
  const apiClient = getApiClient();

  $effect(() => {
    if (open) {
      reset({});
    }
  });

  const { form, errors, enhance, reset, submitting } = superForm(
    defaults(zod(Schema)),
    {
      SPA: true,
      validators: zod(Schema),
      dataType: "json",
      resetForm: true,
      async onUpdate({ form }) {
        if (form.valid) {
          const formData = form.data;

          const data = new FormData();
          data.append("file", formData.file);

          const res = await apiClient.uploadToCollection(collectionId, data);
          if (!res.success) {
            return handleApiError(res.error);
          }

          open = false;
          toast.success("Successfully uploaded file");
          invalidateAll();
        }
      },
    },
  );

  const file = fileProxy(form, "file");
</script>

<Dialog.Root bind:open>
  <Dialog.Content>
    <Dialog.Header>
      <Dialog.Title>Create new collection</Dialog.Title>
    </Dialog.Header>

    <form
      class="flex flex-col gap-4"
      enctype="multipart/form-data"
      use:enhance
    >
      <FormItem>
        <Label for="file">File</Label>
        <input
          id="file"
          name="file"
          type="file"
          accept="image/png, image/jpeg, application/zip"
          bind:files={$file}
        />
        <Errors errors={$errors.file} />
      </FormItem>

      <Dialog.Footer class="gap-2 sm:gap-0">
        <Button
          variant="outline"
          onclick={() => {
            open = false;
          }}
        >
          Close
        </Button>

        <Button type="submit" disabled={$submitting}>
          Create
          {#if $submitting}
            <Spinner />
          {/if}
        </Button>
      </Dialog.Footer>
    </form>
  </Dialog.Content>
</Dialog.Root>
