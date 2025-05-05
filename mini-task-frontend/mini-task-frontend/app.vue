<template>
  <div>
    <!-- Header con navbar -->
    <header class="bg-gray-800 text-white p-4 shadow-md">
      <nav class="container mx-auto flex justify-between items-center">
        <h1 class="text-3xl font-semibold">Mini Task Manager</h1>
        
      </nav>
    </header>

    <!-- Renderiza el contenido de cada página -->
    <main class="container mx-auto p-6">
      <NuxtPage />
    </main>
  </div>
</template>

<script setup>
// Aquí podrías agregar lógica global o de la aplicación si lo necesitas.
</script>

<style scoped>
/* El diseño ya está cubierto por Tailwind CSS, no necesitas agregar estilos adicionales aquí */
</style>
