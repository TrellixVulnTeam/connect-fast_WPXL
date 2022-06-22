<template>
  <main>
    <div class="content">
      <h1>IPFS Settings</h1>
      <div>
        <button class="success" v-if="!connected" @click="connect">
          Connect IPFS.
        </button>
        <button class="danger" v-if="connected" @click="disconnect">
          Disconnect.
        </button>
      </div>
      <div>
        <div
          class="message-output"
          v-for="output in outputs.slice().reverse()"
          :key="output.date"
        >
          {{ output.detail }} <br />
          <span class="time">{{ Date(output.date) }}</span>
        </div>
      </div>
    </div>
  </main>
</template>
<script>
import { StartIPFS, StopIPFS } from "../../../wailsjs/go/main/App";

export default {
  name: "IPFS Settings",
  data() {
    return {
      connected: false,
      outputs: [],
    };
  },
  methods: {
    connect() {
      this.connected = true;
      this.outputs.push({ date: Date.now(), detail: "connected" });
      StartIPFS().then((result) => {
        console.log(result);
        for (let i = 0; i < result.length; i++) {
          this.outputs.push({ date: Date.now(), detail: result[i] });
        }
      });
    },
    disconnect() {
      this.connected = false;
      this.outputs.push({ date: Date.now(), detail: "Disconnected to Daemon" });
      StopIPFS().then((result) => console.log(result));
    },
  },
};
</script>

<style scoped>
button {
  font-size: 14px;
  text-transform: uppercase;
  padding: 8px 16px;
  border: none;
  border-radius: 6px;
}
.success {
  background-color: rgb(142, 255, 142);
}
.danger {
  background-color: rgb(254, 139, 139);
}
.content {
  display: flex;
  flex-direction: column;
}
.time {
  font-size: 6;
  font-weight: 300;
  font-style: italic;
  color: #d0d0d0;
}
.message-output {
  margin-bottom: 0.5em;
  transition: cubic-bezier(0.075, 0.82, 0.165, 1);
}
</style>