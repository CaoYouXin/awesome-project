<script lang="ts">
  import {ctx} from "$lib/components/ctx/tao-form";
  import type {Writable} from "svelte/store";

  export let label: string = '';

  let inline: Writable<boolean> = ctx.get();
</script>

<label class="form-item" class:inline={$inline}>
  <span class="label">{label ? `${label}: ` : ''}</span>
  <br/>
  <slot></slot>
</label>

<style>
  :global(.form-item) + .form-item {
    margin-left: 0;
    margin-top: var(--size-3);
  }

  :global(.form-item.inline) + .form-item.inline {
    margin-left: var(--size-7);
    margin-top: 0;
  }

  .form-item {
    display: block;
  }

  .form-item.inline {
    display: inline-block;
  }
</style>
