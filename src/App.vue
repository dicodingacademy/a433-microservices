<template>
  <div>
    <div class="min-h-screen">
      <div class="flex flex-col justify-between max-w-screen-sm px-4 py-2 mx-auto">
        <app-header></app-header>
        <router-view></router-view>
      </div>
    </div>
  </div>
</template>

<script>

import header from './components/header.vue';

export default {
  components: {
    'app-header': header,
  },
  data () {
    return {

    }
  },
  methods: {

  }
}
</script>

<style>

</style>
