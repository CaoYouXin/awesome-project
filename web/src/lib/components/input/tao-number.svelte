<script lang="ts">
  export let value: any;
  export let placeholder: string;
</script>

<input type="number" {placeholder} bind:value/>
