<script>
  import { page } from '$app/stores';
</script>

<nav>
  <a href="/" aria-current={$page.url.pathname === '/'}>首页</a>
  <a href="/person-jgg" aria-current={$page.url.pathname === '/person-jgg'}>个人九宫格</a>
</nav>

<slot></slot>

<style>
  :global(*) {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
  }

  nav {
    background: var(--stone-1);
  }

  nav > a {
    appearance: none;
    text-decoration: none;
    display: inline-block;
    padding: var(--size-2) var(--size-3);
    background: transparent;
    color: var(--blue-10);
  }

  nav > a[aria-current="true"] {
    box-shadow: var(--shadow-4);
    background: var(--blue-4);
    color: white;
  }
</style>
