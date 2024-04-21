<template>
  <div>
    <h1>Data from API:</h1>
    <ol>
      <li v-for="item in tasks" :key="item.id">
        <div class="task">
          <h4>{{item.title}}</h4>
        <p>{{item.description}}</p>
        </div>
        </li>
    </ol>
  </div>
</template>

<script>

import axios from 'axios';

export default {
  data(){
    return {
      tasks: [],
    };
  },

  mounted() {
    this.fetchDatafromAPI();
  },

  methods :{
    fetchDatafromAPI(){
      const api_url = 'http://localhost:8000/api/tasks';

      axios.get(api_url)
      .then((response) => {
        console.log(response.data)
        this.tasks = response.data;
        console.log(this.tasks)

      })
      .catch((error) => {
        console.error('Error fetching data:',error)
      });
    }
  }
}
</script>
<style scoped>
.task{
  border: 2px solid;
  margin: 5px;
  padding: 4px;
  text-align: center;
  
}
</style>
