<template>
  <div class="container mx-auto p-6">
    <!-- Título Principal -->
    <h1 class="text-4xl font-extrabold text-center text-gray-800 mb-6 bg-gradient-to-r from-blue-500 to-teal-400 text-transparent bg-clip-text">
      Bienvenido a tu Gestionador de tareas
    </h1>
    
    <!-- Descripción -->
    <p class="text-xl text-center text-gray-700 mb-10 max-w-3xl mx-auto">
      Esta es la página principal de la aplicación de tareas. Administra tus tareas de manera fácil y eficiente.
    </p>
    <!-- Subtítulo -->
     <div class="padding-4 mb-6"> 
      <h2 class="text-2xl font-semibold text-gray-800 mb-6">Tus Tareas</h2>
      <!-- Botón para abrir el modal -->
      <div class="mt-8 text-center">
        <button @click="showModal = true" class="bg-blue-500 text-white px-6 py-3 rounded-lg hover:bg-blue-600 transition duration-300 cursor-pointer">
          Crear nueva tarea
      </button>
     </div>
    
    </div>

    <!-- Lista de Tareas -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div v-for="task in tasks" :key="task._id" class="task-card">
        
        <!-- Card component with button -->
        <Card 
          :title="task.title" 
          :description="task.description" 
          :createdAt="task.createdAt"
          :class="{'bg-green-100': task.completed, 'bg-white': !task.completed}">
          
          <button 
            @click="updateTask(task._id)" 
            class="mt-4 px-6 py-2 rounded-md transition duration-300 cursor-pointer"
            :class="task.completed ? 'bg-gray-500' : 'bg-blue-500'">
            {{ task.completed ? 'Completada' : 'Marcar como completada' }}
          </button>
        </Card>
      </div>
    </div>


    <!-- Modal de Nueva Tarea -->
    <div v-if="showModal" class="fixed inset-0 bg-gray-500 bg-opacity-75 flex justify-center items-center z-50">
      <div class="bg-white p-6 rounded-lg shadow-lg w-1/2">
        <h2 class="text-2xl font-semibold mb-4">Crear Nueva Tarea</h2>
        
        <!-- Aquí puedes colocar el formulario para crear una nueva tarea -->
        <form @submit.prevent="createTask">
          <div class="mb-4">
            <label for="title" class="block text-gray-700">Título</label>
            <input type="text" v-model="newTask.title" id="title" class="w-full p-2 border rounded-md" required />
          </div>
          
          <div class="mb-4">
            <label for="description" class="block text-gray-700">Descripción</label>
            <input type="text" v-model="newTask.description" id="description" class="w-full p-2 border rounded-md" required />
          </div>
          
          <div class="flex justify-end">
            <button type="button" @click="closeModal" class="bg-gray-500 text-white px-4 py-2 rounded-md cursor-pointer hover:bg-gray-600">Cancelar</button>
            <button type="submit" class="bg-blue-500 text-white px-6 py-2 ml-2 rounded-md cursor-pointer hover:bg-blue-600" >Guardar</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const tasks = ref([])
const showModal = ref(false)  // Para controlar la visibilidad del modal
const newTask = ref({
  title: '',
  description: ''
})

// Función para abrir el modal
const openModal = () => {
  showModal.value = true
}

// Función para cerrar el modal
const closeModal = () => {
  showModal.value = false
  // Reiniciar el formulario de nueva tarea
  newTask.value = {
    title: '',
    description: ''
  }
  fetchTasks()  // Volver a cargar las tareas después de cerrar el modal
}

// Función para crear una nueva tarea
const createTask = async () => {
  try {
    const response = await axios.post('http://localhost:8080/tasks', newTask.value)
    tasks.value.push(response.data)  // Agregar la nueva tarea a la lista de tareas
    closeModal()  // Cerrar el modal después de agregar la tarea
  } catch (error) {
    console.error('Error al crear la tarea:', error)
  }
}

//Función para actualizar la tarea
const updateTask = async (taskId) => {
  try {
    const response = await axios.put(`http://localhost:8080/tasks/${taskId}`)
    const updatedTask = tasks.value.find(task => task._id === taskId)
    if (updatedTask) {
      updatedTask.completed = response.data.completed
    }
  } catch (error) {
    console.error('Error al actualizar la tarea:', error)
  }
  fetchTasks()  // Volver a cargar las tareas después de actualizar
}

// Obtener las tareas
const fetchTasks = async () => {
  try {
    const response = await axios.get('http://localhost:8080/tasks')
    tasks.value = response.data
  } catch (error) {
    console.error('Error al obtener las tareas:', error)
  }
}

onMounted(fetchTasks)
</script>

<style scoped>
/* Estilo opcional para el modal */
</style>
