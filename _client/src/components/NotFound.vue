<template>
  <div class="NotFound">
    <h1>{{ msg }}</h1>
    <h2>Pages</h2>
    <ul>
      <li><router-link to="/">Todo</router-link></li>
      <li><router-link to="/hello">Hello</router-link></li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'hello',
  data () {
    return {
      msg: 'Not Found (´・ω・｀)'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}
</style>
