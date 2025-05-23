<script setup>
import BookmarkButton from "./BookmarkButton.vue";
import LikeButton from "./LikeButton.vue";
import RateStars from "./RateStars.vue";

const props = defineProps({
  image: {
    type: String,
    default: "/landing.svg",
  },
  recipe: {
    type: Object,
    required: true,
    default: () => ({}),
  },
});

const emit = defineEmits(['show-buy-modal'])

const handleCardClick = () => {
  emit('show-buy-modal', props.recipe);
}

</script>

<template>
  <div class="w-full rounded-lg shadow-md overflow-hidden bg-white relative cursor-pointer" @click="handleCardClick">
    <img :src="recipe.image || image" :alt="recipe.title || 'Recipe Image'" class="w-full h-36 object-cover" />
    <BookmarkButton :recipeId="recipe.id" />

    <div class="p-3 relative">
      <h3 class="font-semibold text-sm text-gray-900">{{ recipe.title || 'Recipe Title' }}</h3>
      <p class="text-xs text-gray-600">
        By <span class="text-rose-500 font-semibold">{{ recipe.author || 'Author' }}</span>
      </p>
      <p class="text-xs text-gray-600 mt-1">Total time: {{ recipe.time || 'N/A' }}</p>
      <div class="flex items-center mt-1">
        <RateStars :rating="recipe.rating" />
        <span class="text-xs text-gray-600 ml-1">({{ recipe.reviews || 0 }})</span>
      </div>

      <LikeButton :recipeId="recipe.id" />
    </div>
  </div>
</template>
