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
  .form-item {
    margin-bottom: var(--size-3);
    margin-right: 0;
    display: block;

    /*font-size: var(--font-size-6);*/
  }

  .form-item.inline {
    margin-bottom: 0;
    margin-right: var(--size-7);
    display: inline-block;
  }

  :global(.form-item:last-of-type) {
    margin-bottom: 0;
    margin-right: 0;
  }

  .form-item .label {
    font-size: var(--font-size-2);
    font-weight: var(--font-weight-9);
  }
</style>
