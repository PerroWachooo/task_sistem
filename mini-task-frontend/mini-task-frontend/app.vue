<template>
  <div>
    <!-- Puedes agregar aquí un encabezado global o navbar -->
    <header>
      <nav>
        <ul>
          <li><router-link to="/tasks">Lista de Tareas</router-link></li>
          <li><router-link to="/new-task">Crear Tarea</router-link></li>
        </ul>
      </nav>
    </header>

    <!-- Renderiza el contenido de cada página -->
    <NuxtPage />
  </div>
</template>

<script setup>
// Si necesitas usar algo global, como el estado de la aplicación, puedes hacerlo aquí
</script>

<style scoped>
/* Aquí puedes agregar estilos globales o del layout */
header {
  background-color: #333;
  color: white;
  padding: 10px;
}

nav ul {
  list-style-type: none;
  display: flex;
}

nav ul li {
  margin-right: 15px;
}

a {
  color: white;
  text-decoration: none;
}
</style>
