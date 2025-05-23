<script setup lang="ts">
import DefaultNavbar from "../components/navbar/DefaultNavbar.vue";
</script>
<template>
  <div class="min-h-screen">
    <DefaultNavbar />
    <slot />
  </div>
</template>
