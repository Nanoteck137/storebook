<script lang="ts">
  import { PUBLIC_COMMIT, PUBLIC_VERSION } from "$env/static/public";
  import { Button, Card } from "@nanoteck137/nano-ui";

  const { data } = $props();
</script>

<Card.Root>
  <div class="flex flex-col gap-4 border-b p-6">
    <h2 class="text-lg font-semibold">Server Infomation</h2>

    <div class="flex flex-col gap-2">
      <div class="flex justify-between">
        <p>API Version</p>
        <p>{data.systemInfo.version}</p>
      </div>

      <div class="flex justify-between">
        <p>Web Version</p>
        <p>{PUBLIC_VERSION}</p>
      </div>

      <div class="flex justify-between">
        <p>Web Commit</p>
        <p>{PUBLIC_COMMIT}</p>
      </div>
    </div>
  </div>

  <div class="flex flex-col gap-4 border-b p-6">
    <h2 class="text-lg font-semibold">Maintenance</h2>

    <div class="flex gap-2">
      <Button>Clear Cache</Button>
    </div>
  </div>

  <div class="flex flex-col gap-4 p-6">
    <h2 class="text-lg font-semibold">Logs</h2>

    <ul
      class="max-h-32 space-y-1 overflow-y-auto rounded-lg bg-muted p-3 font-mono text-sm text-muted-foreground"
    >
      <li>[12:01] Server started successfully</li>
      <li>[12:05] Cleared cache (35 MB)</li>
      <li>[12:10] User "John" logged in</li>
      <li>[12:15] Server running health check</li>
    </ul>
  </div>
</Card.Root>
