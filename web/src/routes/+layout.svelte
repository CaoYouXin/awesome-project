<script lang="ts">
  import {page} from '$app/stores';
  import IndexIcon from '$lib/assets/images/crystal.png';
  import '$lib';
</script>

<svelte:head>
  <title>道医九宫格测算</title>
</svelte:head>

<nav>
  <a href="/" aria-current={$page.url.pathname === '/'}><img class="icon" src={IndexIcon} alt="克里斯托水晶" />首页</a>
  <a href="/tao-jgg" aria-current={$page.url.pathname === '/tao-jgg'}>道医九宫格</a>
  <a href="/calendar" aria-current={$page.url.pathname === '/calendar'}>运势日历</a>
  <a href="/test" aria-current={$page.url.pathname === '/test'}>测试信息</a>
</nav>

<div class="main">
  <slot></slot>
</div>

<style>
  :global(*) {
    margin: 0;
    padding: 0;
    box-sizing: border-box;

    font-family: var(--font-serif);
  }

  nav {
    display: flex;
    justify-content: flex-start;
    background: var(--stone-1);
  }

  nav > a {
    appearance: none;
    text-decoration: none;
    padding: 10px 20px;
    line-height: 20px;
    /*font-size: 20px;*/
    background: transparent;
    color: var(--blue-10);
    box-shadow: var(--shadow-1);
    transition: all 400ms var(--ease-out-3);
    display: flex;
    justify-content: center;
    align-items: center;
  }

  nav > a[aria-current="true"] {
    box-shadow: var(--shadow-4);
    background: var(--blue-4);
    color: white;
  }

  nav > a > .icon {
    width: 20px;
    height: 20px;
  }

  .main {
    width: 100vw;
    height: calc(100vh - 40px);
    padding: 20px;
  }
</style>
