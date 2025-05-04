<template>
    <div>
      <h1>Actualizar Tarea</h1>
      <form @submit.prevent="updateTask">
        <div>
          <label for="title">Título:</label>
          <input type="text" v-model="task.title" id="title" required />
        </div>
        <div>
          <label for="description">Descripción:</label>
          <input type="text" v-model="task.description" id="description" required />
        </div>
        <div>
          <label for="completed">Completada:</label>
          <input type="checkbox" v-model="task.completed" id="completed" />
        </div>
        <div>
          <button type="submit">Actualizar Tarea</button>
        </div>
      </form>
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted } from 'vue'
  import axios from 'axios'
  import { useRoute, useRouter } from 'vue-router'
  
  const task = ref({
    title: '',
    description: '',
    completed: false,
  })
  
  const route = useRoute()
  const router = useRouter()
  
  // Obtener la tarea por su ID al montar el componente
  const fetchTask = async () => {
    const { id } = route.params
    try {
      const response = await axios.get(`http://localhost:8080/tasks/${id}`)
      task.value = response.data
    } catch (error) {
      console.error('Error al obtener la tarea:', error)
    }
  }
  
  // Llamar a la función fetchTask cuando el componente se monte
  onMounted(fetchTask)
  
  // Actualizar la tarea
  const updateTask = async () => {
    const { id } = route.params
    try {
      await axios.put(`http://localhost:8080/tasks/${id}/complete`, task.value)
      router.push('/tasks') // Redirigir a la lista de tareas
    } catch (error) {
      console.error('Error al actualizar la tarea:', error)
    }
  }
  </script>
  