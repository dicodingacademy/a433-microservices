<template>
  <div id="show-jobs-list">
                <header class="mb-10">
                    <div class="text-center font-bold text-4xl text-gray-500 py-4">Karsa Jobs</div>
                    <div class="text-center">Search jobs near you.</div>
                </header>
                <div class="border border-red-500 rounded flex mb-10">
                    <input type="text" v-model="search" class="px-4 py-2 w-full" placeholder="Search...">
                    <button class="flex px-4 border-l items-center bg-red-500">
                        <svg class="h-4 w-4 text-white" fill="currentColor" xmlns="http://www.w3.org/2000/svg"
                            viewBox="0 0 24 24">
                            <path
                                d="M16.32 14.9l5.39 5.4a1 1 0 0 1-1.42 1.4l-5.38-5.38a8 8 0 1 1 1.41-1.41zM10 16a6 6 0 1 0 0-12 6 6 0 0 0 0 12z" />
                        </svg>
                    </button>
                </div>
                <div class="jobs-list">
                    <p class="text-center text-gray-500" v-if="status.status === 0">can't connect to backend server.<br/>{{ status.url }} <b class="text-red-500">{{status.statusText}}</b></p>
                    <p class="text-center text-gray-500" v-if="!filteredJobs.length && status.status !== 0">no result found.</p>
                    <div class="flex border-b border-gray-300 items-center py-4" v-for="job in filteredJobs" :key="job.ID">
                        <div class="w-3/4">
                            <p class="text-gray-600 text-lg font-bold"><router-link v-bind:to="'/job/' + job.ID">{{ job.Role }}</router-link></p>
                            <span class="text-gray-500">{{ job.Company }}</span><small
                                class="ml-2 text-xs bg-red-400 rounded rounded-md text-white p-1">Full Time</small>
                        </div>
                        <div class="w-1/3 text-right">
                            <p class="text-gray-600">{{ job.Location }}</p>
                            <p class="text-gray-600 text-sm">{{ getHumanDate(job.PublishedAt) }}</p>
                        </div>
                    </div>
                </div>
  </div>
</template>

<script>
import moment from 'moment';

export default {
  data() {
    return {
      jobs: [],
      search: '',
      status: ''
    }
  },
  methods: {
    getHumanDate : function (date) {
      return moment(date, 'YYYY-MM-DD').fromNow();
    }
  },
  created() {
    this.$http.get(process.env.VUE_APP_BACKEND + '/jobs').then(function(data){
      this.jobs = data.body;
    }, error => {
      this.status = error;
    });
  },
  computed: {
    filteredJobs(){
      return this.jobs.filter((job) => {
        return job.Role.toLowerCase().match(this.search.toLowerCase());
      });
    }
    
  }
}
</script>

<style>

</style>
