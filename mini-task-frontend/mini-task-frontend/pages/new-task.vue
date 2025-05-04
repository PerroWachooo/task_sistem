<template>
    <div>
      <h1>Crear Nueva Tarea</h1>
      <form @submit.prevent="createTask">
        <div>
          <label for="title">Título:</label>
          <input type="text" v-model="title" id="title" required />
        </div>
        <div>
          <label for="description">Descripción:</label>
          <input type="text" v-model="description" id="description" required />
        </div>
        <div>
          <button type="submit">Crear Tarea</button>
        </div>
      </form>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  import axios from 'axios'
  import { useRouter } from 'vue-router'
  
  const title = ref('')
  const description = ref('')
  const router = useRouter()
  
  const createTask = async () => {
    try {
      await axios.post('http://localhost:8080/tasks', {
        title: title.value,
        description: description.value,
        completed: false,
      })
      router.push('/tasks') // Redirige a la página de tareas
    } catch (error) {
      console.error('Error al crear la tarea:', error)
    }
  }
  </script>
  