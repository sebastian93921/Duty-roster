<template>

</template>

<script>
var contentConfig = { resumeView : false }

function setToResumeView(){
  contentConfig.resumeView = true
}

function setToNormalView(){
  contentConfig.resumeView = false
}

export default {
  //Var
  contentConfig,

  //Function
  setToResumeView,
  setToNormalView,


}
</script>

<style scoped>

</style>
