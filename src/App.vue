<template>
  <nav>
    <div class="nav-left">
        <h3>Logo</h3>
      <router-link to="/">Home</router-link> |
      <router-link to="/tasks">Tasks</router-link>
    </div>
    
    <div class="nav-right">
      <router-link to="/about">About</router-link>
      <router-link to="/contact">Contact</router-link>
      <router-link to=""></router-link>
    </div>
  </nav>
  <router-view/>
</template>

<style scoped>
  nav{
    border-bottom: 1px solid ;
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 10px;
    width: 100%;
    margin-left: -20px;
    margin-top: -10px;
  
  }

  .nav-left{
    display: flex;
    justify-content: space-around;
    align-items: center;
    padding: 10px;
  }

  .nav-left a{
    margin-right: 5px;
    margin-left: 5px;
}

.nav-right{
  display: flex;
  padding: 10px;
}

.nav-right a{
  margin: 5px;
}
</style>
