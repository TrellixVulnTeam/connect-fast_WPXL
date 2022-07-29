<template>
 <main>
    <div class="spinner"></div>

 </main>   
</template>
<style scoped>
svg {
  max-width: 25em;
  border-radius: 3px;
  box-shadow: 2px 2px 5px #000;
  background: #111;
  fill: none;
  stroke: #222;
  stroke-linecap: round;
  stroke-width: 8%
}

use {
  stroke: #fff;
  animation: a 2s linear infinite
}

@keyframes a { to { stroke-dashoffset: 0px } }
</style>