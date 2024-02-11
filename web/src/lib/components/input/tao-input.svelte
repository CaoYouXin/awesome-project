<script lang="ts">
  export let value: any;
  export let placeholder: string;
</script>

<input type="text" {placeholder} bind:value/>
