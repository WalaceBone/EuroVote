<template>
    <div>
      <D3Chart :data="chartData" />
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  import D3Chart from './D3Chart.vue';
  
  export default {
    data() {
      return {
        chartData: null
      };
    },
    components: {
      D3Chart
    },
    created() {
      axios.get('https://your-api-endpoint/meps')
        .then(response => {
          this.chartData = response.data;
        })
        .catch(error => {
          console.error("There was an error fetching the data: ", error);
        });
    }
  }
  </script>
  