<template>
  <div class="max-w-3xl mx-auto p-6 bg-white shadow-md rounded-lg my-8">
    <h1 class="text-3xl font-bold text-center mb-4 text-gray-800">Add a Recipe</h1>
    <p class="text-center text-gray-600 mb-8">
      Feeling like a kitchen Picasso? We want to see your masterpiece! Add your recipe and show off your
      culinary creativity.
    </p>

    <hr class="mb-8" />

    <!-- Add Photo Section -->
    <div class="border border-dashed border-gray-400 p-6 rounded-lg text-center bg-gray-100 bg-opacity-50 mb-8">
      <button class="px-6 py-3 bg-green-500 text-white rounded-full flex items-center justify-center mx-auto hover:bg-green-800 transition-colors">
       <Icon name="uil:plus" class="h-5 w-5 mr-2" />
        Add a photo
      </button>
    </div>

    <!-- Recipe Details -->
    <div class="space-y-6">
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1">Recipe Title <span class="text-red-500">*</span></label>
        <input type="text" placeholder="Enter your recipe title" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-green-600 focus:border-transparent">
      </div>

      <div>
        <label class="block text-sm font-medium text-gray-700 mb-1">Description <span class="text-red-500">*</span></label>
        <textarea placeholder="Describe your recipe in a way that makes mouths water." rows="4" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-green-600 focus:border-transparent"></textarea>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Servings <span class="text-red-500">*</span></label>
          <input type="text" placeholder="e.g., 4" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-green-600 focus:border-transparent">
        </div>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Cook Time <span class="text-red-500">*</span></label>
          <div class="flex gap-4">
            <input type="text" placeholder="hrs" class="w-1/2 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-green-600 focus:border-transparent">
            <input type="text" placeholder="mins" class="w-1/2 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-green-600 focus:border-transparent">
          </div>
        </div>
      </div>

      <hr class="my-8" />

      <!-- Ingredients -->
     <IngredientForm/>

      <hr class="my-8" />

      <!-- Instructions -->
      <div>
        <label class="block text-sm font-medium text-gray-700 mb-2">Instructions <span class="text-red-500">*</span></label>

        <div class="space-y-4">
          <!-- Example Instruction Step -->
          <div class="flex items-start gap-4">
            <div class="flex-grow">
              <label class="block text-sm font-medium text-gray-700 mb-1">Step 1</label>
              <textarea placeholder="Input Text" rows="2" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-green-600 focus:border-transparent"></textarea>
            </div>
            <button class="text-gray-500 hover:text-red-600 transition-colors mt-2">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
           <!-- Repeat similar structure for more steps or use v-for -->
        </div>

         <button class="mt-4 text-green-700 flex items-center hover:text-green-800 transition-colors">
          <Icon name="uil:plus" class="h-5 w-5 mr-2" />
           Add step
        </button>
      </div>

      <hr class="my-8" />

      <!-- Cook's Tips -->
       <div>
        <label class="block text-sm font-medium text-gray-700 mb-1">Cook's Tips</label>
        <textarea placeholder="Share your kitchen secrets! Oven hacks, swaps, or any tips for ultimate recipe success." rows="4" class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-green-600 focus:border-transparent"></textarea>
      </div>

      <hr class="my-8" />

      <!-- Tags -->
      <div class="space-y-4">
        <label class="block text-sm font-medium text-gray-700">Tags</label>
        <div>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-green-600 focus:border-transparent">
                <option value="">Cuisine</option>
                <!-- Add cuisine options here -->
            </select>
        </div>
         <div>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-green-600 focus:border-transparent">
                <option value="">Meal Type</option>
                <!-- Add meal type options here -->
            </select>
        </div>
         <div>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-green-600 focus:border-transparent">
                <option value="">Dietary Restrictions</option>
                <!-- Add dietary restrictions options here -->
            </select>
        </div>
         <div>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-green-600 focus:border-transparent">
                <option value="">Cooking Method</option>
                 <!-- Add cooking method options here -->
            </select>
        </div>
         <div>
            <select class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-green-600 focus:border-transparent">
                <option value="">Main Ingredient</option>
                 <!-- Add main ingredient options here -->
            </select>
        </div>
      </div>

      <!-- Buttons -->
      <div class="flex justify-end gap-4 mt-8">
          <button class="px-6 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-100 transition-colors">Cancel</button>
          <button class="px-6 py-2 bg-green-700 text-white rounded-md hover:bg-green-800 transition-colors">Submit Recipe</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// Add script logic here for form handling, adding/removing ingredients/instructions, etc.
</script>

<style scoped>
/* Add component specific styles here if needed */
</style>
