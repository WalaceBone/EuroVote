<template>
  <div>
    <button @click="fetchData">Load Data</button>
    <div v-if="personData">
      <PersonCard 
        v-for="person in personData" 
        :key="person.Identifier" 
        :personData="person" 
      />
    </div>
  </div>
</template>
  
<script>
import axios from 'axios';
import PersonCard from './PersonCard.vue';

export default {
  data() {
    return {
      personData: []
    };
  },
  components: {
    PersonCard
  },
  methods: {
    fetchData() {
      axios.get('http://localhost:8080/meps')
        .then(response => {
          this.personData = response.data;
          console.log(this.personData);
          this.$emit('data-loaded', this.personData);
          console.log(this.personData.length);

        })
        .catch(error => {
          console.error("There was an error fetching the data: ", error);
        });
    }
  },
}
</script>
  