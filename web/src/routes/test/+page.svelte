<script lang="ts">
  let apiRoot = '';

  (async () => {
    apiRoot = await window.getApiRoot();
  })();
</script>

<div>
  <strong>API Root: </strong>{ apiRoot }
</div>
