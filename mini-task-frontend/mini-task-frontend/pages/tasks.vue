<template>
    <div>
      <h1>Lista de Tareas</h1>
      <ul>
        <li v-for="task in tasks" :key="task._id">
          <p><strong>{{ task.title }}</strong></p>
          <p>{{ task.description }}</p>
          <p>Completada: {{ task.completed ? 'Sí' : 'No' }}</p>
          <p><strong>Fecha de Creación:</strong> {{ task.createdAt }}</p>
          <router-link :to="'/update-task/' + task._id">Actualizar</router-link>
        </li>
      </ul>
      <router-link to="/new-task">Crear nueva tarea</router-link>
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted } from 'vue'
  import axios from 'axios'
  
  const tasks = ref([])
  
  const fetchTasks = async () => {
    try {
      const response = await axios.get('http://localhost:8080/tasks')
      tasks.value = response.data
    } catch (error) {
      console.error('Error al obtener las tareas:', error)
    }
  }
  
  // Cargar las tareas cuando el componente se monta
  onMounted(fetchTasks)
  </script>
  