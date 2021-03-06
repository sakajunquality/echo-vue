<template>
    <section class="todoapp">
      <header class="header">
        <h1>todos</h1>
        <input class="new-todo"
          autofocus autocomplete="off"
          placeholder="What needs to be done?"
          v-model="newTodo"
          @keyup.enter="addTodo">
      </header>
      <section class="main" v-show="todos.length" v-cloak>
        <input class="toggle-all" type="checkbox" v-model="allDone">
        <ul class="todo-list">
          <li v-for="todo in filteredTodos"
            class="todo"
            :key="todo.id"
            :class="{ completed: todo.completed, editing: todo == editedTodo }">
            <div class="view">
              <input class="toggle" type="checkbox" v-model="todo.completed">
              <label @dblclick="editTodo(todo)">{{ todo.title }}</label>
              <button class="destroy" @click="removeTodo(todo)"></button>
            </div>
            <input class="edit" type="text"
              v-model="todo.title"
              v-todo-focus="todo == editedTodo"
              @blur="doneEdit(todo)"
              @keyup.enter="doneEdit(todo)"
              @keyup.esc="cancelEdit(todo)">
          </li>
        </ul>
      </section>
    </section>
</template>

<script>

// visibility filters
var filters = {
  all: function (todos) {
    return todos
  },
  active: function (todos) {
    return todos.filter(function (todo) {
      return !todo.completed
    })
  },
  completed: function (todos) {
    return todos.filter(function (todo) {
      return todo.completed
    })
  }
}

// app Vue instance
export default {
  data: function () {
    return {
      todos: this.fetchData(),
      newTodo: '',
      editedTodo: null,
      visibility: 'all',
      uid: null
    }
  },
  watch: {
    todos: 'fetchData'
  },

  // computed properties
  // https://vuejs.org/guide/computed.html
  computed: {
    filteredTodos: function () {
      return filters[this.visibility](this.todos)
    },
    remaining: function () {
      return filters.active(this.todos).length
    },
    allDone: {
      get: function () {
        return this.remaining === 0
      },
      set: function (value) {
        this.todos.forEach(function (todo) {
          todo.completed = value
        })
      }
    }
  },

  filters: {
    pluralize: function (n) {
      return n === 1 ? 'item' : 'items'
    }
  },

  // methods that implement data logic.
  // note there's no DOM manipulation here at all.
  methods: {
    fetchData: function () {
      var xhr = new XMLHttpRequest()
      var self = this
      xhr.open('GET', 'http://localhost:9090/todos')
      xhr.onload = function () {
        var list = []
        var todos = JSON.parse(xhr.responseText)
        for (var key in todos) {
          list.push(
            {
              id: key,
              title: todos[key]['title'],
              completed: todos[key]['completed']
            }
          )
        }
        self.uid = list.length
        self.todos = list
      }
      xhr.send()
    },
    addTodo: function () {
      var value = this.newTodo && this.newTodo.trim()
      if (!value) {
        return
      }
      var xhr = new XMLHttpRequest()
      var params = 'title=' + value
      xhr.open('POST', 'http://localhost:9090/todos/add')
      xhr.setRequestHeader('Content-type', 'application/x-www-form-urlencoded')

      xhr.onload = function () {
        // @todo idをちゃんと取ってくる
        // var id = JSON.parse(xhr.responseText)
      }
      xhr.send(params)

      this.todos.push({
        id: this.uid++,
        title: value,
        completed: false
      })
      this.newTodo = ''
    },

    removeTodo: function (todo) {
      var delId = todo.id
      var xhr = new XMLHttpRequest()
      var self = this
      xhr.open('DELETE', 'http://localhost:9090/todos/' + delId)
      xhr.onload = function () {
        // レスポンスチェック
        console.info('delete:' + delId)
        self.todos.splice(self.todos.indexOf(todo), 1)
        console.log(JSON.parse(xhr.responseText))
      }
      xhr.send()
    },

    editTodo: function (todo) {
      this.beforeEditCache = todo.title
      this.editedTodo = todo
    },

    doneEdit: function (todo) {
      if (!this.editedTodo) {
        return
      }
      this.editedTodo = null
      todo.title = todo.title.trim()
      if (!todo.title) {
        this.removeTodo(todo)
      }
    },

    cancelEdit: function (todo) {
      this.editedTodo = null
      todo.title = this.beforeEditCache
    },

    removeCompleted: function () {
      this.todos = filters.active(this.todos)
    }
  },

  // a custom directive to wait for the DOM to be updated
  // before focusing on the input field.
  // https://vuejs.org/guide/custom-directive.html
  directives: {
    'todo-focus': function (el, binding) {
      if (binding.value) {
        el.focus()
      }
    }
  }
}
</script>
