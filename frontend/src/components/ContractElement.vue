<script>
import ProgressBar from "./DesignElements/ProgressBar.vue";


export default {
  name: "Contract Element",
  components: { ProgressBar }
};
</script>

<template>
  <main>
    <div class="box" @click="$router.push('/contract/100')">
      <h1 class="title">Data collection for improvement of product.</h1>

      <div class="box-content">
        <p class="description">
          Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent
          auctor tempor libero, eget consectetur nisl viverra eu. Ut fringilla
          tempor metus, non imperdiet eros rhoncus eu. Cras euismod orci quam,
          eget blandit augue porttitor a. Aliquam quam erat, tincidunt sed
          ornare quis, convallis a mi. Etiam sit amet accumsan purus. Sed id
          justo a nisi blandit condimentum. Nam eu libero turpis. In varius quam
          ante, sed euismod neque condimentum at. Maecenas efficitur fermentum
          interdum. Suspendisse rhoncus pharetra mi, non vulputate urna suscipit
          at. Vestibulum in maximus ante, sed consectetur lectus.
        </p>
      </div>
      <div class="box-extra">
        <table class="additional">
          <tr>
            <td>Owner</td>
            <td>Facebook</td>
          </tr>
          <tr>
            <td>Reward</td>
            <td>10 dCFX</td>
          </tr>
          <tr>
            <td>Min participants</td>
            <td>100.000</td>
          </tr>
          <tr>
            <td>Data Privacy</td>
            <td>High</td>
          </tr>
          <tr>
            <td>Resell rights</td>
            <td>No</td>
          </tr>
        </table>
        <table class="additional">
          <tr>
            <td>Start date</td>
            <td>10-10-2022 10:23</td>
          </tr>
          <tr>
            <td>End Date</td>
            <td>04-11-2022 21:34</td>
          </tr>
          <tr>
            <td>Current Participants</td>
            <td>1.000</td>
          </tr>
          <tr>
            <td>Max Participants</td>
            <td>unlimited</td>
          </tr>
          <tr>
            <td>Country</td>
            <td>United States of America</td>
          </tr>
        </table>
      </div>

      <div style="width: 100%">
        <ProgressBar />
      </div>
    </div>
  </main>
</template>

<style>
.box {
  border-radius: 4px;
  background: white;
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 300px;
  padding: 24px;
  box-shadow: rgba(50, 50, 93, 0.25) 0px 2px 5px -1px,
    rgba(0, 0, 0, 0.3) 0px 1px 3px -1px;
}
.box:hover{
    box-shadow: rgba(0, 0, 0, 0.3) 0px 19px 38px, rgba(0, 0, 0, 0.22) 0px 15px 12px;
    cursor: pointer;
}
.title {
  font-size: 18px;
  font-weight: 600;
  margin-top: 0;
  margin-bottom: 0;
}
.top-box {
  display: flex;
  align-content: center;
  align-items: center;
  margin: 8px;
}
.description {
  font-size: 14px;
  font-weight: 200;

  /**Truncate text**/
  overflow: hidden;
  line-height: 1.3rem;
  max-height: 8rem;
  -webkit-box-orient: vertical;
  display: block;
  display: -webkit-box;
  overflow: hidden !important;
  text-overflow: ellipsis;
  -webkit-line-clamp: 4;
}

.additional {
  font-size: 12px;
  text-align: left;
  align-content: left;
  min-width: 100px;
}
.box-extra {
  margin: 10px;
  display: flex;
  gap: 5px;
  flex-wrap: wrap;
}
</style>