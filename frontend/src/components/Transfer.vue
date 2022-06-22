<script setup>
import {reactive} from 'vue'
import {TransferFunds} from '../../wailsjs/go/main/App'

const data = reactive({
  result: "Transfer funds, there will be nothing displatyed here..  👇",
})

function transfer() {
  TransferFunds().then(result => {
    data.resultText = result
  })
}

</script>

<template>
  <main>
    
    <div id="input" class="input-box">
      <button class="btn" @click="transfer">Transfer to</button>
    </div>
    <div id="result" class="result">{{ data.resultText }}</div>
  </main>
</template>

<style scoped>
.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

</style>