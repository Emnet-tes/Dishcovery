<script setup>
import { ref } from 'vue';
import FilterPanel from '~/components/FilterPanel.vue';
import RecipeCard from '~/components/RecipeCard.vue';
import SearchBar from '~/components/SearchBar.vue';
import { dummyRecipes } from '~/constants/constants';

const searchIngredient = ref('');
const showMobileFilters = ref(false);

// State for the buy modal
const showBuyModal = ref(false);
const selectedRecipe = ref(null);

// Function to handle the event emitted by RecipeCard
const handleShowBuyModal = (recipe) => {
  selectedRecipe.value = recipe;
  showBuyModal.value = true;
};

// Function to simulate buying and navigate
const handleBuyConfirm = () => {
  // In a real app, you would handle payment/purchase logic here
  console.log(`Simulating purchase for recipe: ${selectedRecipe.value.title}`);

  // Navigate to the recipe detail page
  if (selectedRecipe.value && selectedRecipe.value.id) {
    navigateTo(`/recipes/${selectedRecipe.value.id}`);
  }

  // Close the modal
  showBuyModal.value = false;
  selectedRecipe.value = null; // Clear selected recipe
};

const handleCloseModal = () => {
  showBuyModal.value = false;
  selectedRecipe.value = null; // Clear selected recipe
};
</script>
<template>
  <div class="min-h-screen bg-white">
    <!-- Header with SearchBar -->

    <div class="bg-green-600 py-4 mx-auto">
      <nav class="text-sm text-white mb-4 ml-5">
        <ol class="list-none p-0 inline-flex">
          <li class="flex items-center">
            <NuxtLink to="/" class=" hover:text-gray-700">Home</NuxtLink>
            <svg class="fill-current w-3 h-3 mx-2" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 320 512">
              <path
                d="M285.476 272.971L91.132 467.314c-9.373 9.373-24.569 9.373-33.941 0l-22.667-22.667c-9.357-9.357-9.375-24.522-.04-33.901L188.505 256 34.484 67.255c-9.335-9.379-9.317-24.544.04-33.901l22.667-22.667c9.373-9.373 24.569-9.373 33.941 0L285.475 239.03c9.373 9.372 9.373 24.568.001 33.941z" />
            </svg>
          </li>
          <li class="flex items-center">
            <NuxtLink to="/recipes" class=" hover:text-gray-700">Search</NuxtLink>

          </li>

        </ol>
      </nav>
      <div class="max-w-7xl mx-auto px-4">
        <div class="flex flex-col items-center">
          <SearchBar class="w-full max-w-2xl" />
        </div>
      </div>
    </div>

    <!-- Main Content -->
    <div class="max-w-7xl mx-auto px-4 py-6">

      <div class="grid grid-cols-1 md:grid-cols-4 gap-6">
        <!-- Filters Sidebar -->
        <FilterPanel />
        <!-- Recipes & Sorting -->
        <section class="md:col-span-3 space-y-6">
          <!-- Sort by -->
          <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
            <p class="text-lg font-semibold">Showing results for "Chicken"</p>
            <div class="flex items-center gap-2 w-full sm:w-auto">
              <label class="text-sm text-gray-600 whitespace-nowrap">Sort by:</label>
              <select class="border rounded px-3 py-1 text-sm w-full sm:w-auto">
                <option>Relevance</option>
                <option>Newest</option>
                <option>Rating</option>
                <option>Cooking Time</option>
              </select>
            </div>
          </div>

          <!-- Recipe Cards Grid -->
          <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
            <RecipeCard v-for="recipe in dummyRecipes" :key="recipe.id" :recipe="recipe"
              @show-buy-modal="handleShowBuyModal" />
          </div>
        </section>
      </div>
    </div>
  </div>

  <!-- Buy Recipe Modal -->
  <div v-if="showBuyModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
    <div class="bg-white p-6 rounded-lg shadow-xl max-w-sm w-full mx-4">
      <h2 class="text-xl font-bold text-gray-800 mb-4">Unlock Recipe Details</h2>
      <p class="text-gray-700 mb-6">
        To view the full details for "<span class="font-semibold">{{ selectedRecipe?.title }}</span>", you need to
        purchase it.
      </p>
      <div class="flex justify-end gap-4">
        <button @click="handleCloseModal"
          class="px-4 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-100 transition-colors">
          Cancel
        </button>
        <button @click="handleBuyConfirm"
          class="px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700 transition-colors">
          Buy to View
        </button>
      </div>
    </div>
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