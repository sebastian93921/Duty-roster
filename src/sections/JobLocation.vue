<template>

<div id="location" class="s-location target-section" >
  <h5>請刪減搜尋地址直到得到其中一個結果</h5>
  <google-map :location="jobLocation" />
</div> <!-- end s-job -->
</template>


<script>
// https://api.whatsapp.com/send?phone=85266710749&text=
import GoogleMap from "../components/GoogleMap"
import ToolBar from '../components/ToolBar'

export default {
  props: {
    jobLocation: String
  },
  components: {
    GoogleMap
  },
  mounted() {
    ToolBar.setHomePage(false, '當更位置')
  },
}

</script>

<style scoped>

</style>
