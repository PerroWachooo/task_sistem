<template>
    <div class="bg-white shadow-lg rounded-lg p-6 w-80">
      <h2 class="text-xl font-semibold mb-4">{{ title }}</h2>
      <p class="text-gray-600">{{ description }}</p>
      <slot></slot> <!-- Se puede agregar contenido adicional con slot -->
    </div>
  </template>
  
  <script setup>
  import { computed } from 'vue'
  
  // Recibimos las propiedades
  defineProps({
    title: String,
    description: String,
    createdAt: String, // La fecha que llega en formato ISO
  })
  
  // Computed property para formatear la fecha
  const formattedDate = computed(() => {
    const date = new Date(createdAt) // Convertimos el string ISO a un objeto Date
    const formatted = new Intl.DateTimeFormat('es-ES', {
      year: 'numeric',
      month: 'long',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit',
      second: '2-digit',
    }).format(date) // Formateamos la fecha en el formato deseado
    console.log('Fecha formateada:', formatted) // Imprime la fecha formateada
    return formatted
})
  </script>
  
  <style scoped>
  </style>
