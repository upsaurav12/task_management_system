<template>

<div>
  <form @submit.prevent="addtask" class="form-container">
        <input type="text" class="title" placeholder="Enter the title" v-model="title" required>
        <input type="text" class="description" placeholder="Enter the desciption" v-model="description" required>
        <button type="submit">Submit</button>
  </form>
</div>
  <div>
    <h1>Data from API:</h1>
    <ol>
      <li v-for="item in tasks" :key="item.id">
        <div class="task">
          <h4>{{item.title}}</h4>
          <p>{{item.description}}</p>
          <button class="delete" @click="deleteTask(item.id)">Delete</button>
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
      title: '',
      description: '',
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
    },

    addtask(){
      const api_url1 = 'http://localhost:8000/api/tasks';
      
      const newTask = {
        title: this.title,
        description: this.description
      };

      axios
      .post(api_url1 ,newTask)
      .then((response) => {
        this.tasks.push(response.data)

        this.title = "";
      this.description = "";

      this.fetchDatafromAPI();
      })
      .catch((error) => {
        console.error("Error adding task: ",error)
      });
    },
    deleteTask(id){
      console.log("Hello this is delete button")

      const api_url2 = `http://localhost:8000/api/tasks/${id}`
      
      axios
      .delete(api_url2)
      .then((response) => {
        console.log("Task deleted:", response.data)

        this.tasks = this.tasks.filter((task) => task.id !== id)
      })
      .catch((error) =>{
        console.error("Error deleting task:" , error)
      });
    }
  }
}
</script>
<style scoped>
.task{
  border-bottom: 1px solid;
  border-top: 1px solid ;
  margin: 5px;
  padding: 4px;
  text-align: center;
  
}

.form-container{
  padding: 25px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.form-container input{
  padding: 10px;
  margin: 5px;
}

.form-container button{
  padding: 0 10px 0 10px;
  height: 40px;
}
</style>
