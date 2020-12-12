/* eslint-disable no-unused-vars */

import Vue from 'vue'
import App from './App'

//Http
import axios from 'axios'
import VueAxios from 'vue-axios'

import router from './router'
import Vuex from 'vuex'
import moment from 'moment'
import Vuetify from 'vuetify'
import * as VueGoogleMaps from "vue2-google-maps"
import BootstrapVue from 'bootstrap-vue'

Vue.use(VueAxios, axios, Vuex, moment)
Vue.use(VueGoogleMaps, {
  load: {
    key: "AIzaSyCPBOYOTnFOXbUMIImdqCMO1Uxgahw3_VI",
    libraries: "places"
  },
  installComponents: true,
});
Vue.use(Vuetify)
Vue.use(BootstrapVue);

Vue.config.productionTip = false;

import $ from 'jquery'
Vue.$ = $

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import 'vuetify/dist/vuetify.min.css'
import 'material-design-icons-iconfont/dist/material-design-icons.css'
import 'normalize.css'

/* eslint-disable no-new */
new Vue({
    el: '#app',
    router,
    components: { App },
    render: h => h(App),
    vuetify: new Vuetify()
})

var scrollDuration = 800

async function init() {
  console.log("Start Web")
  smoothScroll()
  backToTop()
}

function smoothScroll() {

    $('.smoothscroll').on('click', function (e) {
        var target = this.hash,
        $target    = $(target);

            e.preventDefault();
            e.stopPropagation();

        $('html, body').stop().animate({
            'scrollTop': $target.offset().top
        }, scrollDuration, 'swing').promise().done(function () {

            window.location.hash = target;
        });
    });

}

function backToTop() {

    var pxShow  = 200,         // height on which the button will show
    fadeInTime  = 400,         // how slow/fast you want the button to show
    fadeOutTime = 400,         // how slow/fast you want the button to hide
    scrollSpeed = 300,         // how slow/fast you want the button to scroll to top. can be a value, 'slow', 'normal' or 'fast'
    goTopButton = $(".go-top")

    // Show or hide the sticky footer button
    $(window).on('scroll', function() {
        if ($(window).scrollTop() >= pxShow) {
            goTopButton.fadeIn(fadeInTime);
        } else {
            goTopButton.fadeOut(fadeOutTime);
        }
    });
}

init()
