<template>
  <v-fade-transition mode="out-in">
    <v-toolbar v-if="nav.show" fixed color="white">
      <template v-if="nav.isHomePage">
        <v-btn icon to="/">
          <v-icon>home</v-icon>
        </v-btn>
      </template>
      <template v-if="!nav.isHomePage">
        <v-btn icon to="/">
          <v-icon>arrow_back</v-icon>
        </v-btn>
      </template>
      <v-toolbar-title>{{ nav.navTitle }}</v-toolbar-title>
      <v-spacer></v-spacer>
      <template v-if="nav.isHomePage">
        <v-btn icon to="/" @click="calMonthView">
          <v-icon>date_range</v-icon>
        </v-btn>
        <v-btn icon to="/" @click="calListMonthView">
          <v-icon>view_list</v-icon>
        </v-btn>
      </template>
    </v-toolbar>
  </v-fade-transition>
</template>

<script>
import $ from 'jquery'

var navConfig = {
  show: true,
  isHomePage: true,
  navTitle: 'SSPBD報更網'
}

export default {
  data (){
    return {
      nav: navConfig
    }
  },

  methods: {
    handleScroll: function (event) {
      if(window.scrollY < $(window).height() * 0.4){
        navConfig.show = true
      }else{
        navConfig.show = false
      }
    },
    calMonthView: function(){
      $('#calendar').fullCalendar('changeView', 'month')
    },
    calListMonthView: function(){
      $('#calendar').fullCalendar('changeView', 'listMonth')
    },
  },

  created () {
    window.addEventListener('scroll', this.handleScroll);
  },

  destroyed () {
    window.removeEventListener('scroll', this.handleScroll);
  },

  setHomePage
}

function setHomePage(isHomePage, title){
  navConfig.isHomePage = isHomePage
  navConfig.navTitle = title
}

</script>

<style scoped>

</style>
