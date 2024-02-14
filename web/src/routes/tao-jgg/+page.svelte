<script lang="ts">
  import TaoButton from "$lib/components/button/tao-button.svelte";
  import TaoForm from "$lib/components/form/tao-form.svelte";
  import TaoFormItem from "$lib/components/form/tao-form-item.svelte";
  import TaoInput from "$lib/components/input/tao-input.svelte";
  import TaoNumber from "$lib/components/input/tao-number.svelte";
  import TaoRadio from "$lib/components/input/tao-radio.svelte";
  import request from "$lib/request";
  import {onMount} from "svelte";

  onMount(getList)

  interface Ge {
    id: number,
    solarGe: string,
    lunarGe: string,
    solarDate: string,
    lunarDate: string,
    element: string,
    hour: number,
  }

  let date: string = '';
  let hour: number = -1;
  let solar: boolean = true;
  let leapMonth: boolean = false;

  let list: Array<Ge> = [];

  async function getList() {
    try {
      list = await request.GET<Array<Ge>>('/birthday')
    } catch (err) {
      alert(err);
    }
  }

  async function onSubmit() {
    const data = {date, hour, solar, leapMonth};
    console.log('data', JSON.stringify(data));
    try {
      await request.POST('/birthday', data);
      await getList()
    } catch (err) {
      alert(err);
    }
  }

  async function delItem(id: number) {
    try {
      await request.DELETE('/birthday/' + id)
      await getList()
    } catch (err) {
      alert(err);
    }
  }
</script>

<div class="header card">
  <TaoForm inline={true}>
    <TaoFormItem label="请输入八位数字生日">
      <TaoInput placeholder="" bind:value={date}/>
    </TaoFormItem>
    <TaoFormItem label="历制">
      <TaoRadio name="solar" label="阳历" value={true} bind:group={solar}/>
      <TaoRadio name="solar" label="阴历" value={false} bind:group={solar}/>
    </TaoFormItem>
    <TaoFormItem label="是否闰月">
      <TaoRadio name="leapMonth" label="是" value={true} bind:group={leapMonth}/>
      <TaoRadio name="leapMonth" label="否" value={false} bind:group={leapMonth}/>
    </TaoFormItem>
    <TaoFormItem label="请输入几点出生，-1代表不知道">
      <TaoNumber placeholder="" bind:value={hour}/>
    </TaoFormItem>
    <TaoFormItem>
      <TaoButton on:click={onSubmit}>提交</TaoButton>
    </TaoFormItem>
  </TaoForm>
</div>

{#if list.length > 0}
  <div class="content card">
    <table class="zigzag">
      <thead>
      <tr>
        <th>阳格</th>
        <th>阴格</th>
        <th>阳历生日</th>
        <th>阴历生日</th>
        <th>出生时间</th>
        <th>操作</th>
      </tr>
      </thead>
      <tbody>
      {#each list as item (item.id)}
        <tr>
          <td>{item.solarGe}</td>
          <td>{item.lunarGe}</td>
          <td>{item.solarDate}</td>
          <td>{item.lunarDate}</td>
          <td>
            {#if item.element}
              {item.hour}点属{item.element}
            {:else}
              未知
            {/if}
          </td>
          <td>
            <TaoButton on:click={() => delItem(item.id)}>删除</TaoButton>
          </td>
        </tr>
      {/each}
      </tbody>
    </table>
  </div>
{/if}

<style>
  .card {
    background: var(--blue-1);
    border-radius: var(--radius-2);
    padding: var(--size-fluid-3);
    box-shadow: var(--shadow-2);

    &:hover {
      box-shadow: var(--shadow-3);
    }

    @media (--motionOK) {
      animation: var(--animation-fade-in);
    }
  }

  .content {
    margin-top: var(--size-5);
    background: var(--blue-10) !important;
    color: var(--blue-1);
  }

  table {
    width: 100%;
  }

  table th, table td {
    padding: var(--size-relative-2) var(--size-relative-5);
    text-align: left;
  }

  table td {
    background-color: var(--gray-3);
    color: var(--gray-9)
  }

  table th {
    background-color: var(--blue-9);
    color: white;
  }

  .zigzag {
    border-collapse: separate;
    border-spacing: var(--size-relative-2) var(--size-relative-3);
  }
</style>
