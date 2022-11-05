<template>
  <div id="show-job">
      <div class="jobs-detail">
        <div class="flex items-center py-4">
          <div class="w-3/4">
            <p class="text-gray-600 text-3xl font-bold">{{ jobdetail.Role }}</p>
            <span class="text-gray-500">{{ jobdetail.Company }}</span><small class="ml-2 text-xs bg-red-400 rounded rounded-md text-white p-1">Full Time</small>
          </div>
          <div class="w-1/3 text-right">
            <button class="bg-red-500 rounded-md text-gray-100 px-4 py-2">Apply Now</button>
          </div>
        </div>
        <div class="jobs-dec text-gray-700">
            <p>{{ jobdetail.Description }}</p>
        </div>
      </div>
  </div>
</template>

<script>

export default {
  data() {
    return {
      id: this.$route.params.id,
      jobdetail: {}
    }
  },
  created() {
    this.$http.get(process.env.VUE_APP_BACKEND + '/job/' + this.id).then(function(data){
      this.jobdetail = data.body;
    });
  },
}
</script>

<style>

</style>
