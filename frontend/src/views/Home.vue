<template>
  <div>
    <Title title="Overview" />

    <div class="homepage">
      <div class="top_homepage">
        <Balance class="balance"/>
        <TopOptions />
      </div>
      <div>
        <Chart />  
      </div>
      
    </div>
    
  
  </div>
</template>

<style>
.homepage{
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin: 12px;
  
  
}
.top_homepage{
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}
  .balance{
    align-self: flex-start;
  }
</style>

<script>
import Chart from '../components/Dashboard/Chart.vue';
import TopOptions from '../components/Dashboard/TopOptions.vue';
import Balance from '../components/DesignElements/Balance.vue';
import Title from '../components/DesignElements/Title.vue';
import GenerateWallet from '../components/GenerateWallet.vue';
import HelloWorld from '../components/HelloWorld.vue';
import Transfer from '../components/Transfer.vue';
export default {
    name: "Home",
    components: { HelloWorld, GenerateWallet, Transfer, Title, Balance, Chart, TopOptions }
};
</script>