<script setup lang="ts">
// No script logic needed for this static example
</script>

<template>
  <nav class="bg-white shadow-sm sticky top-0 z-10">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between h-16 items-center">

        <!-- Left side - Brand logo -->
        <div class="flex items-center">
          <NuxtLink to="/"
            class="text-xl font-bold text-gray-900 hover:text-primary-600 transition-colors flex items-center gap-2">
            <!-- Placeholder for actual logo SVG if needed -->
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-green-700" fill="none" viewBox="0 0 24 24"
              stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round"
                d="M9 17V7m0 10a2 2 0 01-2 2H5a2 2 0 01-2-2V7a2 2 0 012-2h2a2 2 0 012 2m0 10a2 2 0 002 2h2a2 2 0 002-2V7a2 2 0 00-2-2h-2a2 2 0 00-2 2m3 14l9-9m-9 9L5 5" />
            </svg>
            Dish
            Covery
          </NuxtLink>
        </div>

        <!-- Center - Navigation links -->
        <div class="hidden md:flex flex-1 justify-center items-center space-x-6">
          <NuxtLink to="/recipes"
            class="text-sm font-medium text-gray-700 hover:text-primary-600 transition-colors uppercase"
            active-class="text-primary-600">
            Recipes
          </NuxtLink>
          <NuxtLink to="/articles"
            class="text-sm font-medium text-gray-700 hover:text-primary-600 transition-colors uppercase"
            active-class="text-primary-600">
            Articles
          </NuxtLink>
          <NuxtLink to="/about"
            class="text-sm font-medium text-gray-700 hover:text-primary-600 transition-colors uppercase"
            active-class="text-primary-600">
            About
          </NuxtLink>
        </div>

        <!-- Right side - Action items -->
        <div class="flex items-center space-x-4">
          <NuxtLink to="/recipes/create"
            class="hidden md:flex items-center text-sm font-medium text-green-700 hover:text-green-800 transition-colors whitespace-nowrap">
            + ADD A RECIPE
          </NuxtLink>

          <NuxtLink to="/bookmarks" class="text-gray-500 hover:text-gray-700">
            <Icon name="tabler:bookmark" class="h-6 w-6" />
          </NuxtLink>

          <!-- User dropdown -->
          <div class="relative">
            <button class="flex items-center text-gray-500 hover:text-gray-700">
              <Icon name="tabler:user-circle" class="h-6 w-6" />
              <Icon name="tabler:chevron-down" class="h-4 w-4 ml-0.5" />
            </button>
            <!-- Dropdown content would go here -->
          </div>

        </div>
      </div>
    </div>

    <!-- Mobile menu button - not shown in image but good for responsiveness -->
    <!-- Add mobile menu content here -->

  </nav>
</template>

<style scoped>
/* Add component specific styles here if needed */
</style>
