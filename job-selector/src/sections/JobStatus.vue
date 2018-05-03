<template>

<div id="job" class="s-job target-section" >
  <div>
    <h1>{{ job.activityName }} {{ job.owed }}</h1>
    <h2 class="duty-key">{{ job.dutyKey }}</h2>
  </div>
  <hr/>
  <div>
    日期/時間： {{ job.date }} {{ job.time }}
  </div>
  <div>地點： {{ job.location }}</div>
  <div>主辦機構： {{ job.organiser }}</div>
  <div>出席/所需： {{ job.numberOfAttendance }} / {{ job.numberOfMember }}</div>
  <div>出席隊員： {{ job.listOfMembers }}</div>
  <div>另一支隊： {{ job.anotherDept }}</div>
  <div>收更日期： {{ job.receviedDate }}</div>
  <hr/>
  <div>
    <v-flex xs8>
      <v-text-field
        label="UI Number"
        id="uiNumber"
        placeholder="請在此填上隊員號碼"
        :rules="[rules.required]"
        v-model="uiNumber"
        />
    </v-flex>
  </div>
  <div>
    <v-btn block large href="javascript:;" @click="sendToWhatsapp(job, uiNumber)">
      <span v-if="senderror == ''">按此以 Whatsapp 報更</span>
      <span v-if="senderror != ''" class="senderror">{{ senderror }}</span>
    </v-btn>
    <v-btn block large :to="'location/' + job.location" href="javascript:;">當更地圖位置</v-btn>
  </div>
</div> <!-- end s-job -->
</template>


<script>
import moment from 'moment'

var sendPhoneNo = "85292516517"

export default {
  data (){
    return {
      rules: {
        required: (value) => !!value || 'Required.'
      },
      uiNumber: '',
      senderror: ''
    }
  },
  props: {
    job: {
    }
  },
  methods: {
    sendToWhatsapp(job, uiNumber){
      if(!!uiNumber){
        if(moment().isBefore(moment(job.date, 'DD-MM-YYYY'))){
          var messageToWhatsapp = "[ Auto Message ] 隊員 *"+uiNumber+"* 報更 %0D%0A"+
                                  "標題： *"+job.activityName+"* %0D%0A"+
                                  "更號： *"+job.dutyKey+"* %0D%0A"+
                                  "日期： *"+job.date+" / "+job.time+"* %0D%0A"
          var res = encodeURI(messageToWhatsapp)
          window.location.href = "https://api.whatsapp.com/send?phone="+sendPhoneNo+"&text="+messageToWhatsapp
        }else{
          this.senderror = "所報的更份已經過期"
        }
      }
    }
  },
  watch: {
    job: function(){
      this.senderror = ''
      this.uiNumber = ''
    }
  }
}

</script>

<style scoped>
.s-job{
  box-shadow: 0px -5px 5px #888;
  padding: 0.8em 0.6em 0.6em 0.6em;
  font-size: 1.2em;
}

.duty-key{
  font-size: 1.6em
}

.senderror{
  color: red
}
</style>
