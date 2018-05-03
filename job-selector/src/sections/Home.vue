<template>
<!-- ==================================================
home
================================================== -->
<div id="home" class="s-home target-section" >

  <div id="divLoading" v-show="events.loading"/>

    <div class="home-content">

      <v-layout
        column
        align-center
        justify-center
        v-touch="{
          left: () => swipe('L'),
          right: () => swipe('R')
        }"
      >
        <full-calendar
          :events="events.calendarEvents"
          :config="config" />
      </v-layout>
    </div> <!-- end home-content -->
    <div id="job-status">
      <template v-if="events.subJob !== null">
        <JobStatus :job="events.subJob" />
      </template>
    </div>
</div> <!-- end s-home -->
</template>

<script>
import SwitchContentService_ from '../services/SwitchContentService'
import ToolBar from '../components/ToolBar'
import JobStatus from './JobStatus'

import {FullCalendar} from 'vue-full-calendar'
import 'fullcalendar/dist/fullcalendar.min.css'
import moment from 'moment'

import axios from 'axios'
import $ from 'jquery'

var calendarEvents = []
var subJob = []

var exportData = {calendarEvents,
                  subJob: null,
                  loading: true
                }

export default {
  data () {
  	return {
  	  events : exportData,
      config: {

        header: {
          left: '',
          center: 'title',
          right: ''
        },
        height: "auto",
        allDaySlot: false,
        slotEventOverlap: true,
        eventLimit: 2,
        defaultView: 'month',
        editable: false,
        aspectRatio: 0.8,
        handleWindowResize: true,
        showNonCurrentDates: false,
        displayEventTime: false,
        eventClick: function(calEvent, jsEvent, view) {
          doEventClick (calEvent, jsEvent, view)
        },
        dayClick: function(date){
          $('#calendar').fullCalendar('changeView', 'agendaDay', date)
        },
        viewRender: function (view, element) {
          exportData.loading = true
          genCal(view.start.format('MM'))
        }
      }
  	}
  },
  components: {
    FullCalendar,
    JobStatus
  },
  methods: {
    swipe (direction) {
      if( direction == 'L' ){
        $('#calendar').fullCalendar('next');
      }else if( direction == 'R' ){
        $('#calendar').fullCalendar('prev');
      }
    }
  },
  mounted() {
    ToolBar.setHomePage(true,'SSPBD報更網')
  }
}

function doEventClick (event, jsEvent, views) {
  exportData.subJob = event.job
  $('html, body').stop().animate({
     'scrollTop': $('#job-status').offset().top - 20
  }, 500, 'swing').promise().done(function () {

     window.location.hash = '#job-status'
  })
}

function genCal(month){
  axios.get("https://dutyserv.strsx.com/status/"+month).then((response) => {
    // console.log(response)
    exportData.calendarEvents = []
    var formatDate = "YYYY-MM-DDTHH:mm:ssZ"
    // console.log(response.data)
    for(var i = 0 ; i < response.data.length ; i++){
      // console.log(response.data[i].date,response.data[i].time)
      var dutyTime = response.data[i].time.split(" - ");
      var startTime = moment(response.data[i].date.substring(6,10) + '-' +
       response.data[i].date.substring(3,5) + '-' +
       response.data[i].date.substring(0,2) + 'T' + dutyTime[0] + ':00+0800'
       , formatDate);
      var endTime = moment(response.data[i].date.substring(6,10) + '-' +
       response.data[i].date.substring(3,5) + '-' +
       response.data[i].date.substring(0,2) + 'T' + dutyTime[1] + ':00+0800'
       , formatDate);

      if ( response.data[i].time == '待定' ){
        startTime = moment(response.data[i].date.substring(6,10) + '-' +
         response.data[i].date.substring(3,5) + '-' +
         response.data[i].date.substring(0,2) + '+0800', 'YYYY-MM-DDZ');
        endTime = startTime;
      }

      if(endTime.isBefore(startTime)){
        var endTime = endTime.add(1, 'days');
      }
      // console.log(response.data[i].time, startTime, endTime);
      var jobData = {
        title: response.data[i].activityName,
        start : startTime.format(formatDate),
        end : endTime.format(formatDate),
        color: response.data[i].isFull ? 'lightblue' : 'red',
        job : response.data[i]
      }
      exportData.calendarEvents.push(jobData)
    }
    exportData.loading = false
  })
}


</script>

<style scoped>
.home-content{
  margin-bottom: 5px;
  font-size: 1.6em;
}

.home-load{
  height: 100%;
  width: 100%;
  background-color: gray;
}

#divLoading
{
  display : block;
  position : fixed;
  z-index: 100;
  background-image : url('../images/sunloading.gif');
  opacity : 0.4;
  background-repeat : no-repeat;
  background-position : center;
  left : 0;
  bottom : 0;
  right : 0;
  top : 0;
}
#loadinggif
{
  left : 50%;
  top : 50%;
  position : absolute;
  z-index : 101;
  width : 32px;
  height : 32px;
  margin-left : -16px;
  margin-top : -16px;
}
</style>
