<template>
<h2>Task Management System</h2>

<div>
    <h1>All Tasks</h1>
    <ul>
      <li v-for="task in tasks" :key="task.id">
        {{ task.title }}
      </li>
    </ul>
  </div>


</template>

<script>
// @ is an alias to /src
import axios from 'axios';
export default {
  data(){
    return {
      tasks: []
    }
  },
  mounted(){
    axios.get('http://www.localhost:8000/api/tasks')
    .then(response => {
      this.tasks = response.data;
    })
    .catch(error => {
      console.error('Error fetching:', error)
    });
  }
}
</script>

<style scoped>
h2{
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
