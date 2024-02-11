<script lang="ts">
  import { ctx } from "$lib/components/ctx/tao-form";
  import {onMount} from "svelte";
  import {writable} from "svelte/store";

  export let inline: boolean = false;

  let inlineCtx = writable<boolean>(false);
  ctx.set(inlineCtx);

  onMount(() => {
    inlineCtx.set(inline);
  });
</script>

<slot></slot>
