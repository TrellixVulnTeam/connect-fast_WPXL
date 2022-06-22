<script setup>
import {reactive} from 'vue'
import {GetBalance} from '../../wailsjs/go/main/App'

const data = reactive({
  result: "Transfer funds, there will be nothing displatyed here..  👇",
})

function readBalance() {
  GetBalance("0x23e2408c6d509aF73B77f5d6E44b6Ba79BC657E4").then(result => {
    data.resultText = result
  })
}

</script>

<template>
  <main>
    
    <div id="input" class="input-box">
      <button class="btn" @click="readBalance">Balance to</button>
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