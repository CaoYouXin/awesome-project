<script lang="ts">
  import TaoButton from "$lib/components/button/tao-button.svelte";
  import TaoForm from "$lib/components/form/tao-form.svelte";
  import TaoFormItem from "$lib/components/form/tao-form-item.svelte";
  import TaoInput from "$lib/components/input/tao-input.svelte";
  import TaoNumber from "$lib/components/input/tao-number.svelte";
  import TaoRadio from "$lib/components/input/tao-radio.svelte";
  import request from "$lib/request";

  interface Ge {
    solarGe: string,
    lunarGe: string,
    solarDate: string,
    lunarDate: string,
    element: string,
  }

  let date: string = '';
  let hour: number = -1;
  let solar: boolean = true;
  let leapMonth: boolean = false;

  let result: Ge | null = null;

  async function onSubmit() {
    const data = {date, hour, solar, leapMonth};
    console.log('data', JSON.stringify(data));
    try {
      result = await request.POST<Ge>('/birthday', data);
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
    <TaoFormItem>
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

{#if result}
  <div class="content card">
    <TaoForm inline={true}>
      <TaoFormItem label="阳格">{result.solarGe}</TaoFormItem>
      <TaoFormItem label="阴格">{result.lunarGe}</TaoFormItem>
      <TaoFormItem label="阳历生日">{result.solarDate}</TaoFormItem>
      <TaoFormItem label="阴历生日">{result.lunarDate}</TaoFormItem>
      <TaoFormItem label="出生时间五行">{result.element}</TaoFormItem>
    </TaoForm>
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

    @media (--motionOK)  {
      animation: var(--animation-fade-in);
    }
  }

  .content {
    margin-top: var(--size-5);
    background: var(--blue-10) !important;
    color: var(--blue-1);
  }
</style>
