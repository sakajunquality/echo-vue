<template>
    <section class="todoapp">
      <header class="header">
        <h1>todos</h1>
        <input class="new-todo"
          autofocus autocomplete="off"
          placeholder="What needs to be done?">
      </header>
      <section class="main">
        <input class="toggle-all" type="checkbox">
        <ul class="todo-list">
          <li class="todo">
            <div class="view"><input type="checkbox" class="toggle"> <label>testtest</label> <button class="destroy"></button></div> <input type="text" class="edit">
          </li>
        </ul>
      </section>
    </section>
</template>

<script>
export default {
  name: 'todo'
}
</script>
