<script setup lang="ts">
import { ingredients, timeRanges } from '~/constants/constants'

const showMobileFilters = ref(false)

const filteredIngredients = computed(() => {
  return ingredients.filter(ingredient => 
    ingredient.toLowerCase().includes(searchIngredient.value.toLowerCase())
  )
})
const searchIngredient = ref('')

const toggleFilters = () => {
  showMobileFilters.value = !showMobileFilters.value
}
</script>

<template>
  <div>
    <!-- Mobile Filter Toggle -->
    <div class="md:hidden mb-4">
      <button 
        @click="toggleFilters"
        class="w-full bg-green-600 text-white py-2 px-4 rounded-md flex items-center justify-center gap-2 hover:bg-green-700 transition-colors"
      >
        <span>{{ showMobileFilters ? 'Hide' : 'Show' }} Filters</span>
        <svg 
          class="w-5 h-5 transition-transform duration-200" 
          :class="{ 'transform rotate-180': showMobileFilters }"
          fill="none" 
          stroke="currentColor" 
          viewBox="0 0 24 24"
        >
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
        </svg>
      </button>
    </div>

    <aside 
      :class="[
        'md:block space-y-6',
        showMobileFilters ? 'block' : 'hidden'
      ]"
      class="md:col-span-1"
    >
      <!-- Ingredients Filter -->
      <div>
        <h3 class="font-semibold mb-2">Ingredients</h3>
        <div class="mb-3">
          <input
            v-model="searchIngredient"
            type="text"
            placeholder="Search ingredients..."
            class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-green-600 focus:border-transparent"
          />
        </div>
        <div class="space-y-1 max-h-64 overflow-y-auto pr-2 custom-scrollbar">
          <label
            v-for="ingredient in filteredIngredients"
            :key="ingredient"
            class="flex items-center gap-2 text-sm hover:bg-gray-50 p-1 rounded"
          >
            <input type="checkbox" class="accent-green-600" />
            {{ ingredient }}
          </label>
        </div>
      </div>

      <!-- Divider -->
      <div class="border-t border-gray-300 my-4"></div>

      <!-- Preparation Time Filter -->
      <div>
        <h3 class="font-semibold mb-2">Preparation Time</h3>
        <div class="space-y-1">
          <label
            v-for="range in timeRanges"
            :key="range.value"
            class="flex items-center gap-2 text-sm"
          >
            <input type="checkbox" class="accent-green-600" />
            {{ range.label }}
          </label>
        </div>
      </div>
    </aside>
  </div>
</template>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  width: 6px;
}

.custom-scrollbar::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 3px;
}

.custom-scrollbar::-webkit-scrollbar-thumb {
  background: #888;
  border-radius: 3px;
}

.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: #555;
}
</style>