<script setup lang="ts">
import { ref } from 'vue';
import RecipeCard from '~/components/RecipeCard.vue';
import { dummyRecipes } from '~/constants/constants'; // Using dummy data as placeholder
import type { Recipe } from '~/types'; // Assuming you have a type definition for Recipe

// In a real application, you would fetch bookmarked recipes here
const bookmarkedRecipes = ref(dummyRecipes.slice(0, 6)) // Displaying a subset of dummy data

// State for the buy modal
const showBuyModal = ref(false)
const selectedRecipe = ref<Recipe | null>(null)

// Function to handle the event emitted by RecipeCard
const handleShowBuyModal = (recipe: Recipe) => {
    selectedRecipe.value = recipe
    showBuyModal.value = true
}

// Function to simulate buying and navigate
const handleBuyConfirm = () => {
    // In a real app, you would handle payment/purchase logic here
    console.log(`Simulating purchase for recipe: ${selectedRecipe.value?.title}`)

    // Navigate to the recipe detail page
    if (selectedRecipe.value && selectedRecipe.value.id) {
        navigateTo(`/recipes/${selectedRecipe.value.id}`)
    }

    // Close the modal
    showBuyModal.value = false
    selectedRecipe.value = null // Clear selected recipe
}

const handleCloseModal = () => {
    showBuyModal.value = false
    selectedRecipe.value = null // Clear selected recipe
}

</script>

<template>
    <div class="min-h-screen bg-white">
        <div class="max-w-7xl mx-auto px-4 py-8">
            <h1 class="text-3xl font-bold text-gray-800 mb-6">Saved Recipes</h1>

            <div v-if="bookmarkedRecipes.length > 0" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
                <RecipeCard v-for="recipe in bookmarkedRecipes" :key="recipe.id" :recipe="recipe"
                    @show-buy-modal="handleShowBuyModal" />
            </div>
            <div v-else class="text-center text-gray-500">
                <p>You haven't bookmarked any recipes yet.</p>
            </div>
        </div>
    </div>

    <!-- Buy Recipe Modal -->
    <div v-if="showBuyModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
        <div class="bg-white p-6 rounded-lg shadow-xl max-w-sm w-full mx-4">
            <h2 class="text-xl font-bold text-gray-800 mb-4">Unlock Recipe Details</h2>
            <p class="text-gray-700 mb-6">
                To view the full details for "<span class="font-semibold">{{ selectedRecipe?.title }}</span>", you need
                to purchase it.\n </p>
            <div class="flex justify-end gap-4">
                <button @click="handleCloseModal"
                    class="px-4 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-100 transition-colors">
                    Cancel\n </button>
                <button @click="handleBuyConfirm"
                    class="px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700 transition-colors">
                    Buy to View\n </button>
            </div>
        </div>
    </div>
</template>

<style scoped>
/* Add component specific styles here if needed */
</style>
