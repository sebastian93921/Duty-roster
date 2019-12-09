<template>


<div v-cloak>
  <div>
    <tool-bar />

    <div class="main-container">
      <router-view />
    </div>

    <Footer/>
  </div>
</div>

</template>

<script>
import ToolBar from './components/ToolBar'
import SectionHome from './sections/Home'
import Footer from './Footer'


//Service
import SwitchContentService_ from './services/SwitchContentService'

export default {
  name: 'App',
  components: {
    ToolBar,

    SectionHome,

    Footer
  },
  data() {
    return{
      contentConfig : SwitchContentService_.contentConfig
    }
  }
}

</script>


<style>
@import './assets/css/external-plugin.css';
@import './assets/css/main.css';


[v-cloak] {
  display: none;
}

.main-container{
  padding-top: 60px;
}
</style>
