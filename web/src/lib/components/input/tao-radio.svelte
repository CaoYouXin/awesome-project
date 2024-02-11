<script lang="ts">
  export let group: any;
  export let value: any;
  export let name: string;
  export let label: string;
</script>

<label>
  <input type="radio" bind:group {name} {value}/>
  <span>{label}</span>
</label>
